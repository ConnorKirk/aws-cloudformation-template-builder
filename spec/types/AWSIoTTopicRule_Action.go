
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSIoTTopicRule_ActionProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html", Properties:map[string]cf.Property{"Firehose":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-firehose", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"FirehoseAction", UpdateType:"Mutable"}, "Kinesis":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-kinesis", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"KinesisAction", UpdateType:"Mutable"}, "Republish":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-republish", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"RepublishAction", UpdateType:"Mutable"}, "S3":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-s3", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"S3Action", UpdateType:"Mutable"}, "CloudwatchMetric":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchmetric", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"CloudwatchMetricAction", UpdateType:"Mutable"}, "Elasticsearch":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-elasticsearch", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"ElasticsearchAction", UpdateType:"Mutable"}, "CloudwatchAlarm":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchalarm", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"CloudwatchAlarmAction", UpdateType:"Mutable"}, "IotAnalytics":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotanalytics", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"IotAnalyticsAction", UpdateType:"Mutable"}, "Sns":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sns", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"SnsAction", UpdateType:"Mutable"}, "Sqs":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sqs", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"SqsAction", UpdateType:"Mutable"}, "DynamoDBv2":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodbv2", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"DynamoDBv2Action", UpdateType:"Mutable"}, "Lambda":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-lambda", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"LambdaAction", UpdateType:"Mutable"}, "DynamoDB":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodb", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"DynamoDBAction", UpdateType:"Mutable"}, "StepFunctions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-stepfunctions", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"StepFunctionsAction", UpdateType:"Mutable"}}}
}

