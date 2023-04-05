package sourcecontrolconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceStateType string

const (
	ComplianceStateTypeCompliant    ComplianceStateType = "Compliant"
	ComplianceStateTypeFailed       ComplianceStateType = "Failed"
	ComplianceStateTypeInstalled    ComplianceStateType = "Installed"
	ComplianceStateTypeNoncompliant ComplianceStateType = "Noncompliant"
	ComplianceStateTypePending      ComplianceStateType = "Pending"
)

func PossibleValuesForComplianceStateType() []string {
	return []string{
		string(ComplianceStateTypeCompliant),
		string(ComplianceStateTypeFailed),
		string(ComplianceStateTypeInstalled),
		string(ComplianceStateTypeNoncompliant),
		string(ComplianceStateTypePending),
	}
}

type MessageLevelType string

const (
	MessageLevelTypeError       MessageLevelType = "Error"
	MessageLevelTypeInformation MessageLevelType = "Information"
	MessageLevelTypeWarning     MessageLevelType = "Warning"
)

func PossibleValuesForMessageLevelType() []string {
	return []string{
		string(MessageLevelTypeError),
		string(MessageLevelTypeInformation),
		string(MessageLevelTypeWarning),
	}
}

type OperatorScopeType string

const (
	OperatorScopeTypeCluster   OperatorScopeType = "cluster"
	OperatorScopeTypeNamespace OperatorScopeType = "namespace"
)

func PossibleValuesForOperatorScopeType() []string {
	return []string{
		string(OperatorScopeTypeCluster),
		string(OperatorScopeTypeNamespace),
	}
}

type OperatorType string

const (
	OperatorTypeFlux OperatorType = "Flux"
)

func PossibleValuesForOperatorType() []string {
	return []string{
		string(OperatorTypeFlux),
	}
}

type ProvisioningStateType string

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = "Accepted"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeRunning   ProvisioningStateType = "Running"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
)

func PossibleValuesForProvisioningStateType() []string {
	return []string{
		string(ProvisioningStateTypeAccepted),
		string(ProvisioningStateTypeDeleting),
		string(ProvisioningStateTypeFailed),
		string(ProvisioningStateTypeRunning),
		string(ProvisioningStateTypeSucceeded),
	}
}
