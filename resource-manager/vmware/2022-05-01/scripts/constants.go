package scripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OptionalParamEnum string

const (
	OptionalParamEnumOptional OptionalParamEnum = "Optional"
	OptionalParamEnumRequired OptionalParamEnum = "Required"
)

func PossibleValuesForOptionalParamEnum() []string {
	return []string{
		string(OptionalParamEnumOptional),
		string(OptionalParamEnumRequired),
	}
}

type ScriptExecutionParameterType string

const (
	ScriptExecutionParameterTypeCredential  ScriptExecutionParameterType = "Credential"
	ScriptExecutionParameterTypeSecureValue ScriptExecutionParameterType = "SecureValue"
	ScriptExecutionParameterTypeValue       ScriptExecutionParameterType = "Value"
)

func PossibleValuesForScriptExecutionParameterType() []string {
	return []string{
		string(ScriptExecutionParameterTypeCredential),
		string(ScriptExecutionParameterTypeSecureValue),
		string(ScriptExecutionParameterTypeValue),
	}
}

type ScriptExecutionProvisioningState string

const (
	ScriptExecutionProvisioningStateCanceled   ScriptExecutionProvisioningState = "Canceled"
	ScriptExecutionProvisioningStateCancelled  ScriptExecutionProvisioningState = "Cancelled"
	ScriptExecutionProvisioningStateCancelling ScriptExecutionProvisioningState = "Cancelling"
	ScriptExecutionProvisioningStateDeleting   ScriptExecutionProvisioningState = "Deleting"
	ScriptExecutionProvisioningStateFailed     ScriptExecutionProvisioningState = "Failed"
	ScriptExecutionProvisioningStatePending    ScriptExecutionProvisioningState = "Pending"
	ScriptExecutionProvisioningStateRunning    ScriptExecutionProvisioningState = "Running"
	ScriptExecutionProvisioningStateSucceeded  ScriptExecutionProvisioningState = "Succeeded"
)

func PossibleValuesForScriptExecutionProvisioningState() []string {
	return []string{
		string(ScriptExecutionProvisioningStateCanceled),
		string(ScriptExecutionProvisioningStateCancelled),
		string(ScriptExecutionProvisioningStateCancelling),
		string(ScriptExecutionProvisioningStateDeleting),
		string(ScriptExecutionProvisioningStateFailed),
		string(ScriptExecutionProvisioningStatePending),
		string(ScriptExecutionProvisioningStateRunning),
		string(ScriptExecutionProvisioningStateSucceeded),
	}
}

type ScriptOutputStreamType string

const (
	ScriptOutputStreamTypeError       ScriptOutputStreamType = "Error"
	ScriptOutputStreamTypeInformation ScriptOutputStreamType = "Information"
	ScriptOutputStreamTypeOutput      ScriptOutputStreamType = "Output"
	ScriptOutputStreamTypeWarning     ScriptOutputStreamType = "Warning"
)

func PossibleValuesForScriptOutputStreamType() []string {
	return []string{
		string(ScriptOutputStreamTypeError),
		string(ScriptOutputStreamTypeInformation),
		string(ScriptOutputStreamTypeOutput),
		string(ScriptOutputStreamTypeWarning),
	}
}

type ScriptParameterTypes string

const (
	ScriptParameterTypesBool         ScriptParameterTypes = "Bool"
	ScriptParameterTypesCredential   ScriptParameterTypes = "Credential"
	ScriptParameterTypesFloat        ScriptParameterTypes = "Float"
	ScriptParameterTypesInt          ScriptParameterTypes = "Int"
	ScriptParameterTypesSecureString ScriptParameterTypes = "SecureString"
	ScriptParameterTypesString       ScriptParameterTypes = "String"
)

func PossibleValuesForScriptParameterTypes() []string {
	return []string{
		string(ScriptParameterTypesBool),
		string(ScriptParameterTypesCredential),
		string(ScriptParameterTypesFloat),
		string(ScriptParameterTypesInt),
		string(ScriptParameterTypesSecureString),
		string(ScriptParameterTypesString),
	}
}

type VisibilityParameterEnum string

const (
	VisibilityParameterEnumHidden  VisibilityParameterEnum = "Hidden"
	VisibilityParameterEnumVisible VisibilityParameterEnum = "Visible"
)

func PossibleValuesForVisibilityParameterEnum() []string {
	return []string{
		string(VisibilityParameterEnumHidden),
		string(VisibilityParameterEnumVisible),
	}
}
