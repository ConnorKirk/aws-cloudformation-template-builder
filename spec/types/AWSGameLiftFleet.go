
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSGameLiftFleetResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute(nil), Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html", Properties:map[string]cf.Property{"Description":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-description", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "DesiredEC2Instances":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-desiredec2instances", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:true, Type:"", UpdateType:"Mutable"}, "EC2InboundPermissions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2inboundpermissions", DuplicatesAllowed:false, ItemType:"IpPermission", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "EC2InstanceType":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2instancetype", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}, "LogPaths":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-logpaths", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Immutable"}, "MinSize":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-minsize", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "ServerLaunchParameters":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchparameters", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Immutable"}, "BuildId":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-buildid", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}, "MaxSize":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-maxsize", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "Name":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-name", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "ServerLaunchPath":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchpath", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}}, AdditionalProperties:false}
}

