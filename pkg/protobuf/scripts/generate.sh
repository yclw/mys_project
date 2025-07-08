PROTO_ROOT="proto"
GEN_DIR="gen"

find $PROTO_ROOT -name '*.proto' | while read -r proto_file; do
  protoc \
    --proto_path=$PROTO_ROOT \
    --go_out=$GEN_DIR \
    --go-grpc_out=$GEN_DIR \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    "$proto_file"
done