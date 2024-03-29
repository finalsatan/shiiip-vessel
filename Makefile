build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/vessel/vessel.proto

	docker build -t shiiip-vessel .

run:
	docker run -d -p 50052:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	shiiip-vessel