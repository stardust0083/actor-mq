protoc ../actor-mq/actor/actor.proto --gogoslick_out=plugins=grpc:..  --proto_path=..
protoc ../actor-mq/remote/remote_protocol.proto --gogoslick_out=plugins=grpc:.. --proto_path=..
protoc ../actor-mq/actor_test/remoting/messages/messages.proto --gogoslick_out=plugins=grpc:.. --proto_path=..
protoc ../actor-mq/mq/pb/messages.proto --gogoslick_out=plugins=grpc:.. --proto_path=..
pause