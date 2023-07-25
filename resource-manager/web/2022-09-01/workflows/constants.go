package workflows

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyType string

const (
	KeyTypeNotSpecified KeyType = "NotSpecified"
	KeyTypePrimary      KeyType = "Primary"
	KeyTypeSecondary    KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeNotSpecified),
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

func parseKeyType(input string) (*KeyType, error) {
	vals := map[string]KeyType{
		"notspecified": KeyTypeNotSpecified,
		"primary":      KeyTypePrimary,
		"secondary":    KeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyType(input)
	return &out, nil
}

type Kind string

const (
	KindStateful  Kind = "Stateful"
	KindStateless Kind = "Stateless"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindStateful),
		string(KindStateless),
	}
}

func parseKind(input string) (*Kind, error) {
	vals := map[string]Kind{
		"stateful":  KindStateful,
		"stateless": KindStateless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Kind(input)
	return &out, nil
}

type OpenAuthenticationProviderType string

const (
	OpenAuthenticationProviderTypeAAD OpenAuthenticationProviderType = "AAD"
)

func PossibleValuesForOpenAuthenticationProviderType() []string {
	return []string{
		string(OpenAuthenticationProviderTypeAAD),
	}
}

func parseOpenAuthenticationProviderType(input string) (*OpenAuthenticationProviderType, error) {
	vals := map[string]OpenAuthenticationProviderType{
		"aad": OpenAuthenticationProviderTypeAAD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenAuthenticationProviderType(input)
	return &out, nil
}

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeNotSpecified ParameterType = "NotSpecified"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureObject ParameterType = "SecureObject"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeNotSpecified),
		string(ParameterTypeObject),
		string(ParameterTypeSecureObject),
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
		"notspecified": ParameterTypeNotSpecified,
		"object":       ParameterTypeObject,
		"secureobject": ParameterTypeSecureObject,
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

type WorkflowProvisioningState string

const (
	WorkflowProvisioningStateAccepted      WorkflowProvisioningState = "Accepted"
	WorkflowProvisioningStateCanceled      WorkflowProvisioningState = "Canceled"
	WorkflowProvisioningStateCompleted     WorkflowProvisioningState = "Completed"
	WorkflowProvisioningStateCreated       WorkflowProvisioningState = "Created"
	WorkflowProvisioningStateCreating      WorkflowProvisioningState = "Creating"
	WorkflowProvisioningStateDeleted       WorkflowProvisioningState = "Deleted"
	WorkflowProvisioningStateDeleting      WorkflowProvisioningState = "Deleting"
	WorkflowProvisioningStateFailed        WorkflowProvisioningState = "Failed"
	WorkflowProvisioningStateInProgress    WorkflowProvisioningState = "InProgress"
	WorkflowProvisioningStateMoving        WorkflowProvisioningState = "Moving"
	WorkflowProvisioningStateNotSpecified  WorkflowProvisioningState = "NotSpecified"
	WorkflowProvisioningStatePending       WorkflowProvisioningState = "Pending"
	WorkflowProvisioningStateReady         WorkflowProvisioningState = "Ready"
	WorkflowProvisioningStateRegistered    WorkflowProvisioningState = "Registered"
	WorkflowProvisioningStateRegistering   WorkflowProvisioningState = "Registering"
	WorkflowProvisioningStateRenewing      WorkflowProvisioningState = "Renewing"
	WorkflowProvisioningStateRunning       WorkflowProvisioningState = "Running"
	WorkflowProvisioningStateSucceeded     WorkflowProvisioningState = "Succeeded"
	WorkflowProvisioningStateUnregistered  WorkflowProvisioningState = "Unregistered"
	WorkflowProvisioningStateUnregistering WorkflowProvisioningState = "Unregistering"
	WorkflowProvisioningStateUpdating      WorkflowProvisioningState = "Updating"
	WorkflowProvisioningStateWaiting       WorkflowProvisioningState = "Waiting"
)

func PossibleValuesForWorkflowProvisioningState() []string {
	return []string{
		string(WorkflowProvisioningStateAccepted),
		string(WorkflowProvisioningStateCanceled),
		string(WorkflowProvisioningStateCompleted),
		string(WorkflowProvisioningStateCreated),
		string(WorkflowProvisioningStateCreating),
		string(WorkflowProvisioningStateDeleted),
		string(WorkflowProvisioningStateDeleting),
		string(WorkflowProvisioningStateFailed),
		string(WorkflowProvisioningStateInProgress),
		string(WorkflowProvisioningStateMoving),
		string(WorkflowProvisioningStateNotSpecified),
		string(WorkflowProvisioningStatePending),
		string(WorkflowProvisioningStateReady),
		string(WorkflowProvisioningStateRegistered),
		string(WorkflowProvisioningStateRegistering),
		string(WorkflowProvisioningStateRenewing),
		string(WorkflowProvisioningStateRunning),
		string(WorkflowProvisioningStateSucceeded),
		string(WorkflowProvisioningStateUnregistered),
		string(WorkflowProvisioningStateUnregistering),
		string(WorkflowProvisioningStateUpdating),
		string(WorkflowProvisioningStateWaiting),
	}
}

func parseWorkflowProvisioningState(input string) (*WorkflowProvisioningState, error) {
	vals := map[string]WorkflowProvisioningState{
		"accepted":      WorkflowProvisioningStateAccepted,
		"canceled":      WorkflowProvisioningStateCanceled,
		"completed":     WorkflowProvisioningStateCompleted,
		"created":       WorkflowProvisioningStateCreated,
		"creating":      WorkflowProvisioningStateCreating,
		"deleted":       WorkflowProvisioningStateDeleted,
		"deleting":      WorkflowProvisioningStateDeleting,
		"failed":        WorkflowProvisioningStateFailed,
		"inprogress":    WorkflowProvisioningStateInProgress,
		"moving":        WorkflowProvisioningStateMoving,
		"notspecified":  WorkflowProvisioningStateNotSpecified,
		"pending":       WorkflowProvisioningStatePending,
		"ready":         WorkflowProvisioningStateReady,
		"registered":    WorkflowProvisioningStateRegistered,
		"registering":   WorkflowProvisioningStateRegistering,
		"renewing":      WorkflowProvisioningStateRenewing,
		"running":       WorkflowProvisioningStateRunning,
		"succeeded":     WorkflowProvisioningStateSucceeded,
		"unregistered":  WorkflowProvisioningStateUnregistered,
		"unregistering": WorkflowProvisioningStateUnregistering,
		"updating":      WorkflowProvisioningStateUpdating,
		"waiting":       WorkflowProvisioningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowProvisioningState(input)
	return &out, nil
}

type WorkflowSkuName string

const (
	WorkflowSkuNameBasic        WorkflowSkuName = "Basic"
	WorkflowSkuNameFree         WorkflowSkuName = "Free"
	WorkflowSkuNameNotSpecified WorkflowSkuName = "NotSpecified"
	WorkflowSkuNamePremium      WorkflowSkuName = "Premium"
	WorkflowSkuNameShared       WorkflowSkuName = "Shared"
	WorkflowSkuNameStandard     WorkflowSkuName = "Standard"
)

func PossibleValuesForWorkflowSkuName() []string {
	return []string{
		string(WorkflowSkuNameBasic),
		string(WorkflowSkuNameFree),
		string(WorkflowSkuNameNotSpecified),
		string(WorkflowSkuNamePremium),
		string(WorkflowSkuNameShared),
		string(WorkflowSkuNameStandard),
	}
}

func parseWorkflowSkuName(input string) (*WorkflowSkuName, error) {
	vals := map[string]WorkflowSkuName{
		"basic":        WorkflowSkuNameBasic,
		"free":         WorkflowSkuNameFree,
		"notspecified": WorkflowSkuNameNotSpecified,
		"premium":      WorkflowSkuNamePremium,
		"shared":       WorkflowSkuNameShared,
		"standard":     WorkflowSkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowSkuName(input)
	return &out, nil
}

type WorkflowState string

const (
	WorkflowStateCompleted    WorkflowState = "Completed"
	WorkflowStateDeleted      WorkflowState = "Deleted"
	WorkflowStateDisabled     WorkflowState = "Disabled"
	WorkflowStateEnabled      WorkflowState = "Enabled"
	WorkflowStateNotSpecified WorkflowState = "NotSpecified"
	WorkflowStateSuspended    WorkflowState = "Suspended"
)

func PossibleValuesForWorkflowState() []string {
	return []string{
		string(WorkflowStateCompleted),
		string(WorkflowStateDeleted),
		string(WorkflowStateDisabled),
		string(WorkflowStateEnabled),
		string(WorkflowStateNotSpecified),
		string(WorkflowStateSuspended),
	}
}

func parseWorkflowState(input string) (*WorkflowState, error) {
	vals := map[string]WorkflowState{
		"completed":    WorkflowStateCompleted,
		"deleted":      WorkflowStateDeleted,
		"disabled":     WorkflowStateDisabled,
		"enabled":      WorkflowStateEnabled,
		"notspecified": WorkflowStateNotSpecified,
		"suspended":    WorkflowStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowState(input)
	return &out, nil
}
