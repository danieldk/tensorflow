#include <stddef.h>

#include <tensor_c_api.h>

TF_Tensor *tfgo_tensor(TF_DataType type, long long *dims, int num_dims,
    void *data, size_t len);
void tfgo_dealloc(void *data, size_t len, void *arg);
