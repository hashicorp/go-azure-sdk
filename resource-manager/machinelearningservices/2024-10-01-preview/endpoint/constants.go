package endpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedContentLevel string

const (
	AllowedContentLevelHigh   AllowedContentLevel = "High"
	AllowedContentLevelLow    AllowedContentLevel = "Low"
	AllowedContentLevelMedium AllowedContentLevel = "Medium"
)

func PossibleValuesForAllowedContentLevel() []string {
	return []string{
		string(AllowedContentLevelHigh),
		string(AllowedContentLevelLow),
		string(AllowedContentLevelMedium),
	}
}

func (s *AllowedContentLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedContentLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedContentLevel(input string) (*AllowedContentLevel, error) {
	vals := map[string]AllowedContentLevel{
		"high":   AllowedContentLevelHigh,
		"low":    AllowedContentLevelLow,
		"medium": AllowedContentLevelMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedContentLevel(input)
	return &out, nil
}

type ContentSafetyStatus string

const (
	ContentSafetyStatusDisabled ContentSafetyStatus = "Disabled"
	ContentSafetyStatusEnabled  ContentSafetyStatus = "Enabled"
)

func PossibleValuesForContentSafetyStatus() []string {
	return []string{
		string(ContentSafetyStatusDisabled),
		string(ContentSafetyStatusEnabled),
	}
}

func (s *ContentSafetyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentSafetyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentSafetyStatus(input string) (*ContentSafetyStatus, error) {
	vals := map[string]ContentSafetyStatus{
		"disabled": ContentSafetyStatusDisabled,
		"enabled":  ContentSafetyStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentSafetyStatus(input)
	return &out, nil
}

type DefaultResourceProvisioningState string

const (
	DefaultResourceProvisioningStateAccepted   DefaultResourceProvisioningState = "Accepted"
	DefaultResourceProvisioningStateCanceled   DefaultResourceProvisioningState = "Canceled"
	DefaultResourceProvisioningStateCreating   DefaultResourceProvisioningState = "Creating"
	DefaultResourceProvisioningStateDeleting   DefaultResourceProvisioningState = "Deleting"
	DefaultResourceProvisioningStateDisabled   DefaultResourceProvisioningState = "Disabled"
	DefaultResourceProvisioningStateFailed     DefaultResourceProvisioningState = "Failed"
	DefaultResourceProvisioningStateNotStarted DefaultResourceProvisioningState = "NotStarted"
	DefaultResourceProvisioningStateScaling    DefaultResourceProvisioningState = "Scaling"
	DefaultResourceProvisioningStateSucceeded  DefaultResourceProvisioningState = "Succeeded"
	DefaultResourceProvisioningStateUpdating   DefaultResourceProvisioningState = "Updating"
)

func PossibleValuesForDefaultResourceProvisioningState() []string {
	return []string{
		string(DefaultResourceProvisioningStateAccepted),
		string(DefaultResourceProvisioningStateCanceled),
		string(DefaultResourceProvisioningStateCreating),
		string(DefaultResourceProvisioningStateDeleting),
		string(DefaultResourceProvisioningStateDisabled),
		string(DefaultResourceProvisioningStateFailed),
		string(DefaultResourceProvisioningStateNotStarted),
		string(DefaultResourceProvisioningStateScaling),
		string(DefaultResourceProvisioningStateSucceeded),
		string(DefaultResourceProvisioningStateUpdating),
	}
}

func (s *DefaultResourceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultResourceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultResourceProvisioningState(input string) (*DefaultResourceProvisioningState, error) {
	vals := map[string]DefaultResourceProvisioningState{
		"accepted":   DefaultResourceProvisioningStateAccepted,
		"canceled":   DefaultResourceProvisioningStateCanceled,
		"creating":   DefaultResourceProvisioningStateCreating,
		"deleting":   DefaultResourceProvisioningStateDeleting,
		"disabled":   DefaultResourceProvisioningStateDisabled,
		"failed":     DefaultResourceProvisioningStateFailed,
		"notstarted": DefaultResourceProvisioningStateNotStarted,
		"scaling":    DefaultResourceProvisioningStateScaling,
		"succeeded":  DefaultResourceProvisioningStateSucceeded,
		"updating":   DefaultResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultResourceProvisioningState(input)
	return &out, nil
}

type DeploymentModelVersionUpgradeOption string

const (
	DeploymentModelVersionUpgradeOptionNoAutoUpgrade                  DeploymentModelVersionUpgradeOption = "NoAutoUpgrade"
	DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired      DeploymentModelVersionUpgradeOption = "OnceCurrentVersionExpired"
	DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable DeploymentModelVersionUpgradeOption = "OnceNewDefaultVersionAvailable"
)

func PossibleValuesForDeploymentModelVersionUpgradeOption() []string {
	return []string{
		string(DeploymentModelVersionUpgradeOptionNoAutoUpgrade),
		string(DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired),
		string(DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
	}
}

func (s *DeploymentModelVersionUpgradeOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentModelVersionUpgradeOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentModelVersionUpgradeOption(input string) (*DeploymentModelVersionUpgradeOption, error) {
	vals := map[string]DeploymentModelVersionUpgradeOption{
		"noautoupgrade":                  DeploymentModelVersionUpgradeOptionNoAutoUpgrade,
		"oncecurrentversionexpired":      DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired,
		"oncenewdefaultversionavailable": DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentModelVersionUpgradeOption(input)
	return &out, nil
}

type EndpointAuthMode string

const (
	EndpointAuthModeAADToken EndpointAuthMode = "AADToken"
	EndpointAuthModeAMLToken EndpointAuthMode = "AMLToken"
	EndpointAuthModeKey      EndpointAuthMode = "Key"
)

func PossibleValuesForEndpointAuthMode() []string {
	return []string{
		string(EndpointAuthModeAADToken),
		string(EndpointAuthModeAMLToken),
		string(EndpointAuthModeKey),
	}
}

func (s *EndpointAuthMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointAuthMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointAuthMode(input string) (*EndpointAuthMode, error) {
	vals := map[string]EndpointAuthMode{
		"aadtoken": EndpointAuthModeAADToken,
		"amltoken": EndpointAuthModeAMLToken,
		"key":      EndpointAuthModeKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointAuthMode(input)
	return &out, nil
}

type EndpointComputeType string

const (
	EndpointComputeTypeAzureMLCompute EndpointComputeType = "AzureMLCompute"
	EndpointComputeTypeKubernetes     EndpointComputeType = "Kubernetes"
	EndpointComputeTypeManaged        EndpointComputeType = "Managed"
)

func PossibleValuesForEndpointComputeType() []string {
	return []string{
		string(EndpointComputeTypeAzureMLCompute),
		string(EndpointComputeTypeKubernetes),
		string(EndpointComputeTypeManaged),
	}
}

func (s *EndpointComputeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointComputeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointComputeType(input string) (*EndpointComputeType, error) {
	vals := map[string]EndpointComputeType{
		"azuremlcompute": EndpointComputeTypeAzureMLCompute,
		"kubernetes":     EndpointComputeTypeKubernetes,
		"managed":        EndpointComputeTypeManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointComputeType(input)
	return &out, nil
}

type EndpointType string

const (
	EndpointTypeAzurePointContentSafety EndpointType = "Azure.ContentSafety"
	EndpointTypeAzurePointLlama         EndpointType = "Azure.Llama"
	EndpointTypeAzurePointOpenAI        EndpointType = "Azure.OpenAI"
	EndpointTypeAzurePointSpeech        EndpointType = "Azure.Speech"
	EndpointTypeManagedOnlineEndpoint   EndpointType = "managedOnlineEndpoint"
	EndpointTypeServerlessEndpoint      EndpointType = "serverlessEndpoint"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeAzurePointContentSafety),
		string(EndpointTypeAzurePointLlama),
		string(EndpointTypeAzurePointOpenAI),
		string(EndpointTypeAzurePointSpeech),
		string(EndpointTypeManagedOnlineEndpoint),
		string(EndpointTypeServerlessEndpoint),
	}
}

func (s *EndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointType(input string) (*EndpointType, error) {
	vals := map[string]EndpointType{
		"azure.contentsafety":   EndpointTypeAzurePointContentSafety,
		"azure.llama":           EndpointTypeAzurePointLlama,
		"azure.openai":          EndpointTypeAzurePointOpenAI,
		"azure.speech":          EndpointTypeAzurePointSpeech,
		"managedonlineendpoint": EndpointTypeManagedOnlineEndpoint,
		"serverlessendpoint":    EndpointTypeServerlessEndpoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointType(input)
	return &out, nil
}

type ModelLifecycleStatus string

const (
	ModelLifecycleStatusGenerallyAvailable ModelLifecycleStatus = "GenerallyAvailable"
	ModelLifecycleStatusPreview            ModelLifecycleStatus = "Preview"
)

func PossibleValuesForModelLifecycleStatus() []string {
	return []string{
		string(ModelLifecycleStatusGenerallyAvailable),
		string(ModelLifecycleStatusPreview),
	}
}

func (s *ModelLifecycleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseModelLifecycleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseModelLifecycleStatus(input string) (*ModelLifecycleStatus, error) {
	vals := map[string]ModelLifecycleStatus{
		"generallyavailable": ModelLifecycleStatusGenerallyAvailable,
		"preview":            ModelLifecycleStatusPreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ModelLifecycleStatus(input)
	return &out, nil
}

type RaiPolicyContentSource string

const (
	RaiPolicyContentSourceCompletion RaiPolicyContentSource = "Completion"
	RaiPolicyContentSourcePrompt     RaiPolicyContentSource = "Prompt"
)

func PossibleValuesForRaiPolicyContentSource() []string {
	return []string{
		string(RaiPolicyContentSourceCompletion),
		string(RaiPolicyContentSourcePrompt),
	}
}

func (s *RaiPolicyContentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyContentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyContentSource(input string) (*RaiPolicyContentSource, error) {
	vals := map[string]RaiPolicyContentSource{
		"completion": RaiPolicyContentSourceCompletion,
		"prompt":     RaiPolicyContentSourcePrompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyContentSource(input)
	return &out, nil
}

type RaiPolicyMode string

const (
	RaiPolicyModeBlocking RaiPolicyMode = "Blocking"
	RaiPolicyModeDefault  RaiPolicyMode = "Default"
	RaiPolicyModeDeferred RaiPolicyMode = "Deferred"
)

func PossibleValuesForRaiPolicyMode() []string {
	return []string{
		string(RaiPolicyModeBlocking),
		string(RaiPolicyModeDefault),
		string(RaiPolicyModeDeferred),
	}
}

func (s *RaiPolicyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyMode(input string) (*RaiPolicyMode, error) {
	vals := map[string]RaiPolicyMode{
		"blocking": RaiPolicyModeBlocking,
		"default":  RaiPolicyModeDefault,
		"deferred": RaiPolicyModeDeferred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyMode(input)
	return &out, nil
}

type RaiPolicyType string

const (
	RaiPolicyTypeSystemManaged RaiPolicyType = "SystemManaged"
	RaiPolicyTypeUserManaged   RaiPolicyType = "UserManaged"
)

func PossibleValuesForRaiPolicyType() []string {
	return []string{
		string(RaiPolicyTypeSystemManaged),
		string(RaiPolicyTypeUserManaged),
	}
}

func (s *RaiPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyType(input string) (*RaiPolicyType, error) {
	vals := map[string]RaiPolicyType{
		"systemmanaged": RaiPolicyTypeSystemManaged,
		"usermanaged":   RaiPolicyTypeUserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyType(input)
	return &out, nil
}

type ServerlessEndpointState string

const (
	ServerlessEndpointStateCreating       ServerlessEndpointState = "Creating"
	ServerlessEndpointStateCreationFailed ServerlessEndpointState = "CreationFailed"
	ServerlessEndpointStateDeleting       ServerlessEndpointState = "Deleting"
	ServerlessEndpointStateDeletionFailed ServerlessEndpointState = "DeletionFailed"
	ServerlessEndpointStateOnline         ServerlessEndpointState = "Online"
	ServerlessEndpointStateReinstating    ServerlessEndpointState = "Reinstating"
	ServerlessEndpointStateSuspended      ServerlessEndpointState = "Suspended"
	ServerlessEndpointStateSuspending     ServerlessEndpointState = "Suspending"
	ServerlessEndpointStateUnknown        ServerlessEndpointState = "Unknown"
)

func PossibleValuesForServerlessEndpointState() []string {
	return []string{
		string(ServerlessEndpointStateCreating),
		string(ServerlessEndpointStateCreationFailed),
		string(ServerlessEndpointStateDeleting),
		string(ServerlessEndpointStateDeletionFailed),
		string(ServerlessEndpointStateOnline),
		string(ServerlessEndpointStateReinstating),
		string(ServerlessEndpointStateSuspended),
		string(ServerlessEndpointStateSuspending),
		string(ServerlessEndpointStateUnknown),
	}
}

func (s *ServerlessEndpointState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerlessEndpointState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerlessEndpointState(input string) (*ServerlessEndpointState, error) {
	vals := map[string]ServerlessEndpointState{
		"creating":       ServerlessEndpointStateCreating,
		"creationfailed": ServerlessEndpointStateCreationFailed,
		"deleting":       ServerlessEndpointStateDeleting,
		"deletionfailed": ServerlessEndpointStateDeletionFailed,
		"online":         ServerlessEndpointStateOnline,
		"reinstating":    ServerlessEndpointStateReinstating,
		"suspended":      ServerlessEndpointStateSuspended,
		"suspending":     ServerlessEndpointStateSuspending,
		"unknown":        ServerlessEndpointStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerlessEndpointState(input)
	return &out, nil
}

type ServerlessInferenceEndpointAuthMode string

const (
	ServerlessInferenceEndpointAuthModeAAD ServerlessInferenceEndpointAuthMode = "AAD"
	ServerlessInferenceEndpointAuthModeKey ServerlessInferenceEndpointAuthMode = "Key"
)

func PossibleValuesForServerlessInferenceEndpointAuthMode() []string {
	return []string{
		string(ServerlessInferenceEndpointAuthModeAAD),
		string(ServerlessInferenceEndpointAuthModeKey),
	}
}

func (s *ServerlessInferenceEndpointAuthMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerlessInferenceEndpointAuthMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerlessInferenceEndpointAuthMode(input string) (*ServerlessInferenceEndpointAuthMode, error) {
	vals := map[string]ServerlessInferenceEndpointAuthMode{
		"aad": ServerlessInferenceEndpointAuthModeAAD,
		"key": ServerlessInferenceEndpointAuthModeKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerlessInferenceEndpointAuthMode(input)
	return &out, nil
}

type ServiceAccountKeyName string

const (
	ServiceAccountKeyNameKeyOne ServiceAccountKeyName = "Key1"
	ServiceAccountKeyNameKeyTwo ServiceAccountKeyName = "Key2"
)

func PossibleValuesForServiceAccountKeyName() []string {
	return []string{
		string(ServiceAccountKeyNameKeyOne),
		string(ServiceAccountKeyNameKeyTwo),
	}
}

func (s *ServiceAccountKeyName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceAccountKeyName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceAccountKeyName(input string) (*ServiceAccountKeyName, error) {
	vals := map[string]ServiceAccountKeyName{
		"key1": ServiceAccountKeyNameKeyOne,
		"key2": ServiceAccountKeyNameKeyTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceAccountKeyName(input)
	return &out, nil
}
