package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AliasPathAttributes string

const (
	AliasPathAttributesModifiable AliasPathAttributes = "Modifiable"
	AliasPathAttributesNone       AliasPathAttributes = "None"
)

func PossibleValuesForAliasPathAttributes() []string {
	return []string{
		string(AliasPathAttributesModifiable),
		string(AliasPathAttributesNone),
	}
}

type AliasPathTokenType string

const (
	AliasPathTokenTypeAny          AliasPathTokenType = "Any"
	AliasPathTokenTypeArray        AliasPathTokenType = "Array"
	AliasPathTokenTypeBoolean      AliasPathTokenType = "Boolean"
	AliasPathTokenTypeInteger      AliasPathTokenType = "Integer"
	AliasPathTokenTypeNotSpecified AliasPathTokenType = "NotSpecified"
	AliasPathTokenTypeNumber       AliasPathTokenType = "Number"
	AliasPathTokenTypeObject       AliasPathTokenType = "Object"
	AliasPathTokenTypeString       AliasPathTokenType = "String"
)

func PossibleValuesForAliasPathTokenType() []string {
	return []string{
		string(AliasPathTokenTypeAny),
		string(AliasPathTokenTypeArray),
		string(AliasPathTokenTypeBoolean),
		string(AliasPathTokenTypeInteger),
		string(AliasPathTokenTypeNotSpecified),
		string(AliasPathTokenTypeNumber),
		string(AliasPathTokenTypeObject),
		string(AliasPathTokenTypeString),
	}
}

type AliasPatternType string

const (
	AliasPatternTypeExtract      AliasPatternType = "Extract"
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
)

func PossibleValuesForAliasPatternType() []string {
	return []string{
		string(AliasPatternTypeExtract),
		string(AliasPatternTypeNotSpecified),
	}
}

type AliasType string

const (
	AliasTypeMask         AliasType = "Mask"
	AliasTypeNotSpecified AliasType = "NotSpecified"
	AliasTypePlainText    AliasType = "PlainText"
)

func PossibleValuesForAliasType() []string {
	return []string{
		string(AliasTypeMask),
		string(AliasTypeNotSpecified),
		string(AliasTypePlainText),
	}
}

type ChangeType string

const (
	ChangeTypeCreate   ChangeType = "Create"
	ChangeTypeDelete   ChangeType = "Delete"
	ChangeTypeDeploy   ChangeType = "Deploy"
	ChangeTypeIgnore   ChangeType = "Ignore"
	ChangeTypeModify   ChangeType = "Modify"
	ChangeTypeNoChange ChangeType = "NoChange"
)

func PossibleValuesForChangeType() []string {
	return []string{
		string(ChangeTypeCreate),
		string(ChangeTypeDelete),
		string(ChangeTypeDeploy),
		string(ChangeTypeIgnore),
		string(ChangeTypeModify),
		string(ChangeTypeNoChange),
	}
}

type DeploymentMode string

const (
	DeploymentModeComplete    DeploymentMode = "Complete"
	DeploymentModeIncremental DeploymentMode = "Incremental"
)

func PossibleValuesForDeploymentMode() []string {
	return []string{
		string(DeploymentModeComplete),
		string(DeploymentModeIncremental),
	}
}

type ExpressionEvaluationOptionsScopeType string

const (
	ExpressionEvaluationOptionsScopeTypeInner        ExpressionEvaluationOptionsScopeType = "Inner"
	ExpressionEvaluationOptionsScopeTypeNotSpecified ExpressionEvaluationOptionsScopeType = "NotSpecified"
	ExpressionEvaluationOptionsScopeTypeOuter        ExpressionEvaluationOptionsScopeType = "Outer"
)

func PossibleValuesForExpressionEvaluationOptionsScopeType() []string {
	return []string{
		string(ExpressionEvaluationOptionsScopeTypeInner),
		string(ExpressionEvaluationOptionsScopeTypeNotSpecified),
		string(ExpressionEvaluationOptionsScopeTypeOuter),
	}
}

type OnErrorDeploymentType string

const (
	OnErrorDeploymentTypeLastSuccessful     OnErrorDeploymentType = "LastSuccessful"
	OnErrorDeploymentTypeSpecificDeployment OnErrorDeploymentType = "SpecificDeployment"
)

func PossibleValuesForOnErrorDeploymentType() []string {
	return []string{
		string(OnErrorDeploymentTypeLastSuccessful),
		string(OnErrorDeploymentTypeSpecificDeployment),
	}
}

type PropertyChangeType string

const (
	PropertyChangeTypeArray  PropertyChangeType = "Array"
	PropertyChangeTypeCreate PropertyChangeType = "Create"
	PropertyChangeTypeDelete PropertyChangeType = "Delete"
	PropertyChangeTypeModify PropertyChangeType = "Modify"
)

func PossibleValuesForPropertyChangeType() []string {
	return []string{
		string(PropertyChangeTypeArray),
		string(PropertyChangeTypeCreate),
		string(PropertyChangeTypeDelete),
		string(PropertyChangeTypeModify),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreated      ProvisioningState = "Created"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateReady        ProvisioningState = "Ready"
	ProvisioningStateRunning      ProvisioningState = "Running"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreated),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateReady),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type WhatIfResultFormat string

const (
	WhatIfResultFormatFullResourcePayloads WhatIfResultFormat = "FullResourcePayloads"
	WhatIfResultFormatResourceIdOnly       WhatIfResultFormat = "ResourceIdOnly"
)

func PossibleValuesForWhatIfResultFormat() []string {
	return []string{
		string(WhatIfResultFormatFullResourcePayloads),
		string(WhatIfResultFormatResourceIdOnly),
	}
}
