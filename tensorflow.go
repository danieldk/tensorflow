package tensorflow

// #cgo CFLAGS: -I/home/ddekok/tensorflow/cpu/0.12
// #cgo LDFLAGS: -L/home/ddekok/tensorflow/cpu/0.12 -lc_api -ltensorflow
// #include <stdlib.h>
// #include <string.h>
// #include <tensor_c_api.h>
// #include "tensorflow.h"
import "C"
import (
	"errors"
	"unsafe"

	tfconfig "github.com/danieldk/tensorflow/config"
	"github.com/golang/protobuf/proto"
)

type SessionOptions struct {
	options *C.TF_SessionOptions
}

func NewSessionOptions() *SessionOptions {
	opts := &SessionOptions{
		options: C.TF_NewSessionOptions(),
	}

	return opts
}

func (opts *SessionOptions) Close() {
	C.TF_DeleteSessionOptions(opts.options)
}

func (opts *SessionOptions) SetConfig(config tfconfig.ConfigProto) error {
	status := C.TF_NewStatus()
	defer C.TF_DeleteStatus(status)

	data, err := proto.Marshal(&config)
	if err != nil {
		return err
	}

	C.TF_SetConfig(opts.options, unsafe.Pointer(&data[0]), C.size_t(len(data)), status)

	if C.TF_GetCode(status) == 0 {
		return nil
	}

	return errors.New(C.GoString(C.TF_Message(status)))
}

type Session struct {
	session *C.TF_Session
}

func NewSession(graph *Graph, opts *SessionOptions) (*Session, error) {
	status := C.TF_NewStatus()
	defer C.TF_DeleteStatus(status)

	session := &Session{
		session: C.TF_NewSession(graph.graph, opts.options, status),
	}

	if C.TF_GetCode(status) != 0 {
		return nil, errors.New(C.GoString(C.TF_Message(status)))
	}

	return session, nil
}

func (s *Session) Close() error {
	status := C.TF_NewStatus()
	defer C.TF_DeleteStatus(status)

	C.TF_CloseSession(s.session, status)
	if C.TF_GetCode(status) != 0 {
		return errors.New(C.GoString(C.TF_Message(status)))
	}

	C.TF_DeleteSession(s.session, status)
	if C.TF_GetCode(status) != 0 {
		return errors.New(C.GoString(C.TF_Message(status)))
	}

	s.session = nil

	return nil
}

func (s *Session) ExtendGraph(data []byte) error {
	status := C.TF_NewStatus()
	defer C.TF_DeleteStatus(status)

	C.TF_ExtendGraph(s.session, unsafe.Pointer(&data[0]), C.size_t(len(data)), status)

	if C.TF_GetCode(status) != 0 {
		return errors.New(C.GoString(C.TF_Message(status)))
	}

	return nil
}

func (s *Session) Run(inputs map[string]Tensor, outputs []string) (map[string]Tensor, error) {
	inputNames := make([]*C.char, len(inputs))
	inputTensors := make([]*C.TF_Tensor, len(inputs))
	idx := 0
	for input, tensor := range inputs {
		cStr := C.CString(input)
		defer C.free(unsafe.Pointer(cStr))
		inputNames[idx] = cStr
		inputTensors[idx] = tensor.toCTensor()
		idx++
	}

	outputNames := make([]*C.char, len(outputs))
	outputTensors := make([]*C.TF_Tensor, len(outputs))
	for idx, output := range outputs {
		cStr := C.CString(output)
		defer C.free(unsafe.Pointer(cStr))
		outputNames[idx] = cStr
	}

	status := C.TF_NewStatus()
	defer C.TF_DeleteStatus(status)

	C.TF_Run(s.session, nil, &inputNames[0], &inputTensors[0], C.int(len(inputNames)),
		&outputNames[0], &outputTensors[0], C.int(len(outputNames)),
		nil, 0, nil, status)

	if C.TF_GetCode(status) != 0 {
		return nil, errors.New(C.GoString(C.TF_Message(status)))
	}

	outputMap := make(map[string]Tensor)
	for idx, name := range outputs {
		outputMap[name] = adoptTensor(outputTensors[idx])
	}

	return outputMap, nil
}
