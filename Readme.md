goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc

goctl model pg datasource -url="host=localhost port=5433 user=root dbname=eventostec password=172983456 sslmode=disable" -table="event" -dir="./zrpc/internal/models"

goctl model pg datasource -url="host=localhost port=5433 user=root dbname=eventostec password=172983456 sslmode=disable" -table="address" -dir="./zrpc/internal/models"

goctl model pg datasource -url="host=localhost port=5433 user=root dbname=eventostec password=172983456 sslmode=disable" -table="coupon" -dir="./zrpc/internal/models"
