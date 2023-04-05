package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugCommandType string

const (
	DataFlowDebugCommandTypeExecuteExpressionQuery DataFlowDebugCommandType = "executeExpressionQuery"
	DataFlowDebugCommandTypeExecutePreviewQuery    DataFlowDebugCommandType = "executePreviewQuery"
	DataFlowDebugCommandTypeExecuteStatisticsQuery DataFlowDebugCommandType = "executeStatisticsQuery"
)

func PossibleValuesForDataFlowDebugCommandType() []string {
	return []string{
		string(DataFlowDebugCommandTypeExecuteExpressionQuery),
		string(DataFlowDebugCommandTypeExecutePreviewQuery),
		string(DataFlowDebugCommandTypeExecuteStatisticsQuery),
	}
}

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    IntegrationRuntimeType = "Managed"
	IntegrationRuntimeTypeSelfHosted IntegrationRuntimeType = "SelfHosted"
)

func PossibleValuesForIntegrationRuntimeType() []string {
	return []string{
		string(IntegrationRuntimeTypeManaged),
		string(IntegrationRuntimeTypeSelfHosted),
	}
}

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeObject),
		string(ParameterTypeSecureString),
		string(ParameterTypeString),
	}
}
