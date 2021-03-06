
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSECSService_PlacementStrategyProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html", Properties:map[string]cf.Property{"Field":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-field", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Immutable"}, "Type":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-placementstrategy.html#cfn-ecs-service-placementstrategy-type", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}}}
}


