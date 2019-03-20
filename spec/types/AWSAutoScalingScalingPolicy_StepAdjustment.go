
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSAutoScalingScalingPolicy_StepAdjustmentProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html", Properties:map[string]cf.Property{"MetricIntervalLowerBound":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervallowerbound", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Double", Required:false, Type:"", UpdateType:"Mutable"}, "MetricIntervalUpperBound":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-metricintervalupperbound", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Double", Required:false, Type:"", UpdateType:"Mutable"}, "ScalingAdjustment":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-stepadjustments.html#cfn-autoscaling-scalingpolicy-stepadjustment-scalingadjustment", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:true, Type:"", UpdateType:"Mutable"}}}
}

