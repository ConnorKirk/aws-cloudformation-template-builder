
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSKinesisStreamConsumerResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"ConsumerCreationTimestamp":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}, "ConsumerName":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}, "ConsumerARN":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}, "ConsumerStatus":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}, "StreamARN":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html", Properties:map[string]cf.Property{"ConsumerName":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-consumername", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}, "StreamARN":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-streamarn", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}}, AdditionalProperties:false}
}


