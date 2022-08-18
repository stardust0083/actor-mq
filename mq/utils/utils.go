package utils

import (
	"actor-mq/pb"
)

func NewCreateChannelMsg(channelName string, channelType pb.ChannelType) *pb.CreateChannelMsg {
	return &pb.CreateChannelMsg{Name: channelName, Type: channelType}
}

func NewPublishMsg(channelName string, content []byte) *pb.PublishMsg {
	return &pb.PublishMsg{Name: channelName, Content: content}
}

func NewSubscribeMsg(channelName string, subscriber int64) *pb.SubscribeMsg {
	return &pb.SubscribeMsg{channelName, subscriber}
}
