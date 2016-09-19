# Purpose

This binding was created to use Tensorflow in some existing Go projects that I
have/had. For this reason it only provides the functionality that I need(ed).
There are currently Go bindings brewing in upstream:

  * Source: https://github.com/tensorflow/tensorflow/tree/master/tensorflow/go
  * Issue: https://github.com/tensorflow/tensorflow/issues/10#issuecomment-245687757

Once this is done, it's better the use the upstream binding, since it is always
in sync with the latest Tensorflow.

# Building

Get v0.8.0 of Tensorflow and build it:

~~~{.bash}
$ git clone https://github.com/tensorflow/tensorflow.git
$ cd tensorflow
$ git checkout v0.8.0
$ git submodule init
$ git submodule update
$ ./configure
$ bazel build -c opt tensorflow:libtensorflow.so
~~~

Now you need to put `tensor_c_api.h` and `libtensorflow.so` visible somewhere for cgo
and your C compiler. On OS X using Homebrew:

~~~{.bash}
$ mkdir -p /usr/local/Cellar/tensorflow/0.8.0/{lib,include}
$ cp bazel-bin/tensorflow/libtensorflow.so /usr/local/Cellar/tensorflow/0.8.0/lib
$ cp tensorflow/core/public/tensor_c_api.h /usr/local/Cellar/tensorflow/0.8.0/include
$ brew link tensorflow
~~~

Now you can build the Go package as normal.
