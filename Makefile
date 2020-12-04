gen:
	protoc apis/employeepb/*.proto --go_out=plugins=grpc:.