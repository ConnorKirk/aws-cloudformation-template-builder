
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSEC2EC2Fleet_TargetCapacitySpecificationRequestProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html", Properties:map[string]cf.Property{"SpotTargetCapacity":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-spottargetcapacity", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "DefaultTargetCapacityType":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-defaulttargetcapacitytype", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "TotalTargetCapacity":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-totaltargetcapacity", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:true, Type:"", UpdateType:"Mutable"}, "OnDemandTargetCapacity":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-ondemandtargetcapacity", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}}}
}


