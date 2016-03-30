#!/bin/sh

# int32 tensors
genny -in=generic_tensor.go -out=tensor_int32.go gen "ValueType=int32 ValueCType=C.TF_INT32"
sed -i -e 's/TF_FLOAT/TF_INT32/g' -e '/tensorflow.h/a import "C"' tensor_int32.go

# float32 tensors
genny -in=generic_tensor.go -out=tensor_float32.go gen "ValueType=float32"
sed -i '/tensorflow.h/a import "C"' tensor_float32.go
