# $1 子系统

goctl rpc protoc ./rpc/$1/$1.proto --zrpc_out=./rpc/$1/ --go_out=./rpc/$1/ --go-grpc_out=./rpc/$1/  --style go_zero