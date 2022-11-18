package deployments

import "strings"

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

func parseAliasPathAttributes(input string) (*AliasPathAttributes, error) {
	vals := map[string]AliasPathAttributes{
		"modifiable": AliasPathAttributesModifiable,
		"none":       AliasPathAttributesNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AliasPathAttributes(input)
	return &out, nil
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

func parseAliasPathTokenType(input string) (*AliasPathTokenType, error) {
	vals := map[string]AliasPathTokenType{
		"any":          AliasPathTokenTypeAny,
		"array":        AliasPathTokenTypeArray,
		"boolean":      AliasPathTokenTypeBoolean,
		"integer":      AliasPathTokenTypeInteger,
		"notspecified": AliasPathTokenTypeNotSpecified,
		"number":       AliasPathTokenTypeNumber,
		"object":       AliasPathTokenTypeObject,
		"string":       AliasPathTokenTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AliasPathTokenType(input)
	return &out, nil
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

func parseAliasPatternType(input string) (*AliasPatternType, error) {
	vals := map[string]AliasPatternType{
		"extract":      AliasPatternTypeExtract,
		"notspecified": AliasPatternTypeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AliasPatternType(input)
	return &out, nil
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

func parseAliasType(input string) (*AliasType, error) {
	vals := map[string]AliasType{
		"mask":         AliasTypeMask,
		"notspecified": AliasTypeNotSpecified,
		"plaintext":    AliasTypePlainText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AliasType(input)
	return &out, nil
}

type ChangeType string

const (
	ChangeTypeCreate      ChangeType = "Create"
	ChangeTypeDelete      ChangeType = "Delete"
	ChangeTypeDeploy      ChangeType = "Deploy"
	ChangeTypeIgnore      ChangeType = "Ignore"
	ChangeTypeModify      ChangeType = "Modify"
	ChangeTypeNoChange    ChangeType = "NoChange"
	ChangeTypeUnsupported ChangeType = "Unsupported"
)

func PossibleValuesForChangeType() []string {
	return []string{
		string(ChangeTypeCreate),
		string(ChangeTypeDelete),
		string(ChangeTypeDeploy),
		string(ChangeTypeIgnore),
		string(ChangeTypeModify),
		string(ChangeTypeNoChange),
		string(ChangeTypeUnsupported),
	}
}

func parseChangeType(input string) (*ChangeType, error) {
	vals := map[string]ChangeType{
		"create":      ChangeTypeCreate,
		"delete":      ChangeTypeDelete,
		"deploy":      ChangeTypeDeploy,
		"ignore":      ChangeTypeIgnore,
		"modify":      ChangeTypeModify,
		"nochange":    ChangeTypeNoChange,
		"unsupported": ChangeTypeUnsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeType(input)
	return &out, nil
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

func parseDeploymentMode(input string) (*DeploymentMode, error) {
	vals := map[string]DeploymentMode{
		"complete":    DeploymentModeComplete,
		"incremental": DeploymentModeIncremental,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentMode(input)
	return &out, nil
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

func parseExpressionEvaluationOptionsScopeType(input string) (*ExpressionEvaluationOptionsScopeType, error) {
	vals := map[string]ExpressionEvaluationOptionsScopeType{
		"inner":        ExpressionEvaluationOptionsScopeTypeInner,
		"notspecified": ExpressionEvaluationOptionsScopeTypeNotSpecified,
		"outer":        ExpressionEvaluationOptionsScopeTypeOuter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpressionEvaluationOptionsScopeType(input)
	return &out, nil
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

func parseOnErrorDeploymentType(input string) (*OnErrorDeploymentType, error) {
	vals := map[string]OnErrorDeploymentType{
		"lastsuccessful":     OnErrorDeploymentTypeLastSuccessful,
		"specificdeployment": OnErrorDeploymentTypeSpecificDeployment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnErrorDeploymentType(input)
	return &out, nil
}

type PropertyChangeType string

const (
	PropertyChangeTypeArray    PropertyChangeType = "Array"
	PropertyChangeTypeCreate   PropertyChangeType = "Create"
	PropertyChangeTypeDelete   PropertyChangeType = "Delete"
	PropertyChangeTypeModify   PropertyChangeType = "Modify"
	PropertyChangeTypeNoEffect PropertyChangeType = "NoEffect"
)

func PossibleValuesForPropertyChangeType() []string {
	return []string{
		string(PropertyChangeTypeArray),
		string(PropertyChangeTypeCreate),
		string(PropertyChangeTypeDelete),
		string(PropertyChangeTypeModify),
		string(PropertyChangeTypeNoEffect),
	}
}

func parsePropertyChangeType(input string) (*PropertyChangeType, error) {
	vals := map[string]PropertyChangeType{
		"array":    PropertyChangeTypeArray,
		"create":   PropertyChangeTypeCreate,
		"delete":   PropertyChangeTypeDelete,
		"modify":   PropertyChangeTypeModify,
		"noeffect": PropertyChangeTypeNoEffect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PropertyChangeType(input)
	return &out, nil
}

type ProviderAuthorizationConsentState string

const (
	ProviderAuthorizationConsentStateConsented    ProviderAuthorizationConsentState = "Consented"
	ProviderAuthorizationConsentStateNotRequired  ProviderAuthorizationConsentState = "NotRequired"
	ProviderAuthorizationConsentStateNotSpecified ProviderAuthorizationConsentState = "NotSpecified"
	ProviderAuthorizationConsentStateRequired     ProviderAuthorizationConsentState = "Required"
)

func PossibleValuesForProviderAuthorizationConsentState() []string {
	return []string{
		string(ProviderAuthorizationConsentStateConsented),
		string(ProviderAuthorizationConsentStateNotRequired),
		string(ProviderAuthorizationConsentStateNotSpecified),
		string(ProviderAuthorizationConsentStateRequired),
	}
}

func parseProviderAuthorizationConsentState(input string) (*ProviderAuthorizationConsentState, error) {
	vals := map[string]ProviderAuthorizationConsentState{
		"consented":    ProviderAuthorizationConsentStateConsented,
		"notrequired":  ProviderAuthorizationConsentStateNotRequired,
		"notspecified": ProviderAuthorizationConsentStateNotSpecified,
		"required":     ProviderAuthorizationConsentStateRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProviderAuthorizationConsentState(input)
	return &out, nil
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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"created":      ProvisioningStateCreated,
		"creating":     ProvisioningStateCreating,
		"deleted":      ProvisioningStateDeleted,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"notspecified": ProvisioningStateNotSpecified,
		"ready":        ProvisioningStateReady,
		"running":      ProvisioningStateRunning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseWhatIfResultFormat(input string) (*WhatIfResultFormat, error) {
	vals := map[string]WhatIfResultFormat{
		"fullresourcepayloads": WhatIfResultFormatFullResourcePayloads,
		"resourceidonly":       WhatIfResultFormatResourceIdOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WhatIfResultFormat(input)
	return &out, nil
}
