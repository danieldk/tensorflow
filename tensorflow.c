#include <stdlib.h>

#include "tensorflow.h"

TF_Tensor *tfgo_tensor(TF_DataType type, long long *dims, int num_dims,
    void *data, size_t len)
{
  return TF_NewTensor(type, dims, num_dims, data, len, tfgo_dealloc, 0);
}

void tfgo_dealloc(void *data, size_t len, void *arg)
{
  free(data);
}
