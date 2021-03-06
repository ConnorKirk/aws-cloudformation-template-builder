
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSEMRInstanceFleetConfig_EbsBlockDeviceConfigProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html", Properties:map[string]cf.Property{"VolumeSpecification":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumespecification", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:true, Type:"VolumeSpecification", UpdateType:"Immutable"}, "VolumesPerInstance":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig.html#cfn-elasticmapreduce-instancefleetconfig-ebsblockdeviceconfig-volumesperinstance", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Immutable"}}}
}


