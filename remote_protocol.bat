IF NOT EXIST "./protocol" (MD "./protocol")


protoc remote_protocol.proto --gogoslick_out=plugins=grpc:./protocol --proto_path=proto 
protoc actor.proto --gogoslick_out=plugins=grpc:./protocol  --proto_path=proto 
pause