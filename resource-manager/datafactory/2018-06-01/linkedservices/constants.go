package linkedservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type Type string

const (
	TypeIntegrationRuntimeReference Type = "IntegrationRuntimeReference"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeIntegrationRuntimeReference),
	}
}
