# grpc-vs-rest

GRPC vs REST client server exchange example. 

Simple client server exchange with performance test. Study 

## Install

Install packages

```
pip install -r requirements.txt
go mod tidy
```

Install protobuf tools (MacOS)

```
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

## Compile code (optional)

Should recompile code only if `user.proto` is changed

Compile go code:

```
protoc --go_out=. --go-grpc_out=. grpc-example/user.proto
```

Compile python code:

```
python -m grpc_tools.protoc -Igrpc-example --python_out=grpc-example/pb --pyi_out=grpc-example/pb --grpc_python_out=grpc-example/pb user.proto
```

## Run code

Run go rest servers:

```
go run ./rest-example/server
go run ./grpc-example/server
```

Now we have REST server on port :8081 and GRPC server on port :50005


Run python rest client:

```
python3 rest-example/client.py
```

Run python rest client:

```
python3 grpc-example/client.py
```

