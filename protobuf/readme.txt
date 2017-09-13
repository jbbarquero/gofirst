From https://developers.google.com/protocol-buffers/docs/gotutorial

See Compiling your protocol buffers

1. Download the package

Go to https://github.com/google/protobuf then go to https://github.com/google/protobuf/releases and go to Downloads

I've installed the application in ~\Applications\protoc-3.4.0-win32

After fixing the environment variables, just try:

$ protoc --version
libprotoc 3.4.0

2. Then choose the appropriate proto compiler, protoc-3.4.0-win32.zip, for instance.

Run the following command to install the Go protocol buffers plugin:

go get -u github.com/golang/protobuf/protoc-gen-go

The compiler plugin protoc-gen-go will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler protoc to find it.

For instance: C:\Users\jbbar\Documents\go\bin\protoc-gen-go.exe

3. Run the compiler, specifying the source directory, the destination directory, and the path to your .proto.

In this case:

protoc --go_out=. addressbook.proto

NOTE: the generated file uses the import proto "github.com/golang/protobuf/proto", so you can get if it's not available yet:

go get github.com/golang/protobuf/proto




