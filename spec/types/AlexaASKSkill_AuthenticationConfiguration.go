
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AlexaASKSkill_AuthenticationConfigurationProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html", Properties:map[string]cf.Property{"RefreshToken":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-refreshtoken", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "ClientSecret":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-clientsecret", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "ClientId":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-authenticationconfiguration.html#cfn-ask-skill-authenticationconfiguration-clientid", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}}}
}

