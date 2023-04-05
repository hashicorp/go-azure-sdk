package pipelineruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunQueryFilterOperand string

const (
	RunQueryFilterOperandActivityName        RunQueryFilterOperand = "ActivityName"
	RunQueryFilterOperandActivityRunEnd      RunQueryFilterOperand = "ActivityRunEnd"
	RunQueryFilterOperandActivityRunStart    RunQueryFilterOperand = "ActivityRunStart"
	RunQueryFilterOperandActivityType        RunQueryFilterOperand = "ActivityType"
	RunQueryFilterOperandLatestOnly          RunQueryFilterOperand = "LatestOnly"
	RunQueryFilterOperandPipelineName        RunQueryFilterOperand = "PipelineName"
	RunQueryFilterOperandRunEnd              RunQueryFilterOperand = "RunEnd"
	RunQueryFilterOperandRunGroupId          RunQueryFilterOperand = "RunGroupId"
	RunQueryFilterOperandRunStart            RunQueryFilterOperand = "RunStart"
	RunQueryFilterOperandStatus              RunQueryFilterOperand = "Status"
	RunQueryFilterOperandTriggerName         RunQueryFilterOperand = "TriggerName"
	RunQueryFilterOperandTriggerRunTimestamp RunQueryFilterOperand = "TriggerRunTimestamp"
)

func PossibleValuesForRunQueryFilterOperand() []string {
	return []string{
		string(RunQueryFilterOperandActivityName),
		string(RunQueryFilterOperandActivityRunEnd),
		string(RunQueryFilterOperandActivityRunStart),
		string(RunQueryFilterOperandActivityType),
		string(RunQueryFilterOperandLatestOnly),
		string(RunQueryFilterOperandPipelineName),
		string(RunQueryFilterOperandRunEnd),
		string(RunQueryFilterOperandRunGroupId),
		string(RunQueryFilterOperandRunStart),
		string(RunQueryFilterOperandStatus),
		string(RunQueryFilterOperandTriggerName),
		string(RunQueryFilterOperandTriggerRunTimestamp),
	}
}

type RunQueryFilterOperator string

const (
	RunQueryFilterOperatorEquals    RunQueryFilterOperator = "Equals"
	RunQueryFilterOperatorIn        RunQueryFilterOperator = "In"
	RunQueryFilterOperatorNotEquals RunQueryFilterOperator = "NotEquals"
	RunQueryFilterOperatorNotIn     RunQueryFilterOperator = "NotIn"
)

func PossibleValuesForRunQueryFilterOperator() []string {
	return []string{
		string(RunQueryFilterOperatorEquals),
		string(RunQueryFilterOperatorIn),
		string(RunQueryFilterOperatorNotEquals),
		string(RunQueryFilterOperatorNotIn),
	}
}

type RunQueryOrder string

const (
	RunQueryOrderASC  RunQueryOrder = "ASC"
	RunQueryOrderDESC RunQueryOrder = "DESC"
)

func PossibleValuesForRunQueryOrder() []string {
	return []string{
		string(RunQueryOrderASC),
		string(RunQueryOrderDESC),
	}
}

type RunQueryOrderByField string

const (
	RunQueryOrderByFieldActivityName        RunQueryOrderByField = "ActivityName"
	RunQueryOrderByFieldActivityRunEnd      RunQueryOrderByField = "ActivityRunEnd"
	RunQueryOrderByFieldActivityRunStart    RunQueryOrderByField = "ActivityRunStart"
	RunQueryOrderByFieldPipelineName        RunQueryOrderByField = "PipelineName"
	RunQueryOrderByFieldRunEnd              RunQueryOrderByField = "RunEnd"
	RunQueryOrderByFieldRunStart            RunQueryOrderByField = "RunStart"
	RunQueryOrderByFieldStatus              RunQueryOrderByField = "Status"
	RunQueryOrderByFieldTriggerName         RunQueryOrderByField = "TriggerName"
	RunQueryOrderByFieldTriggerRunTimestamp RunQueryOrderByField = "TriggerRunTimestamp"
)

func PossibleValuesForRunQueryOrderByField() []string {
	return []string{
		string(RunQueryOrderByFieldActivityName),
		string(RunQueryOrderByFieldActivityRunEnd),
		string(RunQueryOrderByFieldActivityRunStart),
		string(RunQueryOrderByFieldPipelineName),
		string(RunQueryOrderByFieldRunEnd),
		string(RunQueryOrderByFieldRunStart),
		string(RunQueryOrderByFieldStatus),
		string(RunQueryOrderByFieldTriggerName),
		string(RunQueryOrderByFieldTriggerRunTimestamp),
	}
}
