
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSCloudFrontDistributionResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"DomainName":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html", Properties:map[string]cf.Property{"DistributionConfig":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-distributionconfig", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:true, Type:"DistributionConfig", UpdateType:"Mutable"}, "Tags":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html#cfn-cloudfront-distribution-tags", DuplicatesAllowed:false, ItemType:"Tag", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}}, AdditionalProperties:false}
}


