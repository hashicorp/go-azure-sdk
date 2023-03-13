package dataflowdebugsession

import "strings"

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

func parseDataFlowDebugCommandType(input string) (*DataFlowDebugCommandType, error) {
	vals := map[string]DataFlowDebugCommandType{
		"executeexpressionquery": DataFlowDebugCommandTypeExecuteExpressionQuery,
		"executepreviewquery":    DataFlowDebugCommandTypeExecutePreviewQuery,
		"executestatisticsquery": DataFlowDebugCommandTypeExecuteStatisticsQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataFlowDebugCommandType(input)
	return &out, nil
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

func parseIntegrationRuntimeType(input string) (*IntegrationRuntimeType, error) {
	vals := map[string]IntegrationRuntimeType{
		"managed":    IntegrationRuntimeTypeManaged,
		"selfhosted": IntegrationRuntimeTypeSelfHosted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeType(input)
	return &out, nil
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

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":        ParameterTypeArray,
		"bool":         ParameterTypeBool,
		"float":        ParameterTypeFloat,
		"int":          ParameterTypeInt,
		"object":       ParameterTypeObject,
		"securestring": ParameterTypeSecureString,
		"string":       ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}
