package tensorflow

// #include <string.h>
// #include <tensor_c_api.h>
// #include "tensorflow.h"
import "C"

import "fmt"

type Tensor interface {
	toCTensor() *C.TF_Tensor
}

var _ Tensor = &Float32Tensor{}

func adoptTensor(ct *C.TF_Tensor) Tensor {
	defer C.TF_DeleteTensor(ct)

	ttype := C.TF_TensorType(ct)
	switch ttype {
	case C.TF_INT32:
		return adoptint32Tensor(ct)
	case C.TF_FLOAT:
		return adoptfloat32Tensor(ct)
	default:
		panic(fmt.Sprintf("Support for adopting tensor type %d is not implemented", ttype))
	}
}
