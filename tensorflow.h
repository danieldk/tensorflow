#include <stddef.h>

#include <tensorflow/c/c_api.h>

TF_Tensor *tfgo_tensor(TF_DataType type, int64_t const *dims, int num_dims,
    void *data, size_t len);
void tfgo_dealloc(void *data, size_t len, void *arg);
