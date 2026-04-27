GRPC_DIR=internal/adapters/framwork/left/grpc
PROTO_DIR=${GRPC_DIR}/proto

proto_struct:
	@protoc --go_out=${GRPC_DIR} --proto_path=${PROTO_DIR} ${PROTO_DIR}/number_msg.proto

proto_service:
	@protoc --go-grpc_out=${GRPC_DIR} --proto_path=${PROTO_DIR} ${PROTO_DIR}/arithmatic_service.proto

protoc: proto_struct proto_service
