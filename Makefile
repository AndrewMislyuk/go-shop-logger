## generate_audit: generate protobuf audit
generate_audit:
	@echo "Generation protobuf..."
	protoc -I proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/audit.proto
	@echo "Complete"