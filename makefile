
default: clean proto

proto:
	protoc -I=. --go_out=pb --go_opt=module=github.com/Hellizer/binancedataserviceproto/pb  --go-grpc_out=pb --go-grpc_opt=paths=source_relative *.proto 

clean:
	-@rm -r ./pb/* || echo 
	-@mkdir pb || echo 

.PHONY: proto