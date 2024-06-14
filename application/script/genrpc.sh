cd application/user/rpc

goctl rpc protoc ./user.proto --go_out=. --go-grpc_out=. --zrpc_out=./
goctl rpc protoc ./article.proto --go_out=. --go-grpc_out=. --zrpc_out=./