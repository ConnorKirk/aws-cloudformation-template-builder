
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSApiGatewayRestApiResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"RootResourceId":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html", Properties:map[string]cf.Property{"ApiKeySourceType":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-apikeysourcetype", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "EndpointConfiguration":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-endpointconfiguration", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"EndpointConfiguration", UpdateType:"Mutable"}, "MinimumCompressionSize":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-minimumcompressionsize", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "Name":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-name", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "Parameters":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-parameters", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"Map", UpdateType:"Mutable"}, "BinaryMediaTypes":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-binarymediatypes", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Body":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-body", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Json", Required:false, Type:"", UpdateType:"Mutable"}, "BodyS3Location":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-bodys3location", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"S3Location", UpdateType:"Mutable"}, "CloneFrom":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-clonefrom", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "Description":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-description", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "FailOnWarnings":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-failonwarnings", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "Policy":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-policy", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Json", Required:false, Type:"", UpdateType:"Mutable"}}, AdditionalProperties:false}
}

