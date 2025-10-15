package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateAllocationPolicy string

const (
	AkriConnectorTemplateAllocationPolicyBucketized AkriConnectorTemplateAllocationPolicy = "Bucketized"
)

func PossibleValuesForAkriConnectorTemplateAllocationPolicy() []string {
	return []string{
		string(AkriConnectorTemplateAllocationPolicyBucketized),
	}
}

func (s *AkriConnectorTemplateAllocationPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorTemplateAllocationPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorTemplateAllocationPolicy(input string) (*AkriConnectorTemplateAllocationPolicy, error) {
	vals := map[string]AkriConnectorTemplateAllocationPolicy{
		"bucketized": AkriConnectorTemplateAllocationPolicyBucketized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorTemplateAllocationPolicy(input)
	return &out, nil
}

type AkriConnectorTemplateManagedConfigurationType string

const (
	AkriConnectorTemplateManagedConfigurationTypeImageConfiguration       AkriConnectorTemplateManagedConfigurationType = "ImageConfiguration"
	AkriConnectorTemplateManagedConfigurationTypeStatefulSetConfiguration AkriConnectorTemplateManagedConfigurationType = "StatefulSetConfiguration"
)

func PossibleValuesForAkriConnectorTemplateManagedConfigurationType() []string {
	return []string{
		string(AkriConnectorTemplateManagedConfigurationTypeImageConfiguration),
		string(AkriConnectorTemplateManagedConfigurationTypeStatefulSetConfiguration),
	}
}

func (s *AkriConnectorTemplateManagedConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorTemplateManagedConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorTemplateManagedConfigurationType(input string) (*AkriConnectorTemplateManagedConfigurationType, error) {
	vals := map[string]AkriConnectorTemplateManagedConfigurationType{
		"imageconfiguration":       AkriConnectorTemplateManagedConfigurationTypeImageConfiguration,
		"statefulsetconfiguration": AkriConnectorTemplateManagedConfigurationTypeStatefulSetConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorTemplateManagedConfigurationType(input)
	return &out, nil
}

type AkriConnectorTemplateRuntimeConfigurationType string

const (
	AkriConnectorTemplateRuntimeConfigurationTypeManagedConfiguration AkriConnectorTemplateRuntimeConfigurationType = "ManagedConfiguration"
)

func PossibleValuesForAkriConnectorTemplateRuntimeConfigurationType() []string {
	return []string{
		string(AkriConnectorTemplateRuntimeConfigurationTypeManagedConfiguration),
	}
}

func (s *AkriConnectorTemplateRuntimeConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorTemplateRuntimeConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorTemplateRuntimeConfigurationType(input string) (*AkriConnectorTemplateRuntimeConfigurationType, error) {
	vals := map[string]AkriConnectorTemplateRuntimeConfigurationType{
		"managedconfiguration": AkriConnectorTemplateRuntimeConfigurationTypeManagedConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorTemplateRuntimeConfigurationType(input)
	return &out, nil
}

type AkriConnectorsImagePullPolicy string

const (
	AkriConnectorsImagePullPolicyAlways       AkriConnectorsImagePullPolicy = "Always"
	AkriConnectorsImagePullPolicyIfNotPresent AkriConnectorsImagePullPolicy = "IfNotPresent"
	AkriConnectorsImagePullPolicyNever        AkriConnectorsImagePullPolicy = "Never"
)

func PossibleValuesForAkriConnectorsImagePullPolicy() []string {
	return []string{
		string(AkriConnectorsImagePullPolicyAlways),
		string(AkriConnectorsImagePullPolicyIfNotPresent),
		string(AkriConnectorsImagePullPolicyNever),
	}
}

func (s *AkriConnectorsImagePullPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorsImagePullPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorsImagePullPolicy(input string) (*AkriConnectorsImagePullPolicy, error) {
	vals := map[string]AkriConnectorsImagePullPolicy{
		"always":       AkriConnectorsImagePullPolicyAlways,
		"ifnotpresent": AkriConnectorsImagePullPolicyIfNotPresent,
		"never":        AkriConnectorsImagePullPolicyNever,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorsImagePullPolicy(input)
	return &out, nil
}

type AkriConnectorsMqttAuthenticationMethod string

const (
	AkriConnectorsMqttAuthenticationMethodServiceAccountToken AkriConnectorsMqttAuthenticationMethod = "ServiceAccountToken"
)

func PossibleValuesForAkriConnectorsMqttAuthenticationMethod() []string {
	return []string{
		string(AkriConnectorsMqttAuthenticationMethodServiceAccountToken),
	}
}

func (s *AkriConnectorsMqttAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorsMqttAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorsMqttAuthenticationMethod(input string) (*AkriConnectorsMqttAuthenticationMethod, error) {
	vals := map[string]AkriConnectorsMqttAuthenticationMethod{
		"serviceaccounttoken": AkriConnectorsMqttAuthenticationMethodServiceAccountToken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorsMqttAuthenticationMethod(input)
	return &out, nil
}

type AkriConnectorsMqttProtocolType string

const (
	AkriConnectorsMqttProtocolTypeMqtt AkriConnectorsMqttProtocolType = "Mqtt"
)

func PossibleValuesForAkriConnectorsMqttProtocolType() []string {
	return []string{
		string(AkriConnectorsMqttProtocolTypeMqtt),
	}
}

func (s *AkriConnectorsMqttProtocolType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorsMqttProtocolType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorsMqttProtocolType(input string) (*AkriConnectorsMqttProtocolType, error) {
	vals := map[string]AkriConnectorsMqttProtocolType{
		"mqtt": AkriConnectorsMqttProtocolTypeMqtt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorsMqttProtocolType(input)
	return &out, nil
}

type AkriConnectorsRegistrySettingsType string

const (
	AkriConnectorsRegistrySettingsTypeContainerRegistry   AkriConnectorsRegistrySettingsType = "ContainerRegistry"
	AkriConnectorsRegistrySettingsTypeRegistryEndpointRef AkriConnectorsRegistrySettingsType = "RegistryEndpointRef"
)

func PossibleValuesForAkriConnectorsRegistrySettingsType() []string {
	return []string{
		string(AkriConnectorsRegistrySettingsTypeContainerRegistry),
		string(AkriConnectorsRegistrySettingsTypeRegistryEndpointRef),
	}
}

func (s *AkriConnectorsRegistrySettingsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorsRegistrySettingsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorsRegistrySettingsType(input string) (*AkriConnectorsRegistrySettingsType, error) {
	vals := map[string]AkriConnectorsRegistrySettingsType{
		"containerregistry":   AkriConnectorsRegistrySettingsTypeContainerRegistry,
		"registryendpointref": AkriConnectorsRegistrySettingsTypeRegistryEndpointRef,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorsRegistrySettingsType(input)
	return &out, nil
}

type AkriConnectorsTagDigestType string

const (
	AkriConnectorsTagDigestTypeDigest AkriConnectorsTagDigestType = "Digest"
	AkriConnectorsTagDigestTypeTag    AkriConnectorsTagDigestType = "Tag"
)

func PossibleValuesForAkriConnectorsTagDigestType() []string {
	return []string{
		string(AkriConnectorsTagDigestTypeDigest),
		string(AkriConnectorsTagDigestTypeTag),
	}
}

func (s *AkriConnectorsTagDigestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAkriConnectorsTagDigestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAkriConnectorsTagDigestType(input string) (*AkriConnectorsTagDigestType, error) {
	vals := map[string]AkriConnectorsTagDigestType{
		"digest": AkriConnectorsTagDigestTypeDigest,
		"tag":    AkriConnectorsTagDigestTypeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AkriConnectorsTagDigestType(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
	}
}

func (s *ExtendedLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationType(input string) (*ExtendedLocationType, error) {
	vals := map[string]ExtendedLocationType{
		"customlocation": ExtendedLocationTypeCustomLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}

type OperationalMode string

const (
	OperationalModeDisabled OperationalMode = "Disabled"
	OperationalModeEnabled  OperationalMode = "Enabled"
)

func PossibleValuesForOperationalMode() []string {
	return []string{
		string(OperationalModeDisabled),
		string(OperationalModeEnabled),
	}
}

func (s *OperationalMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationalMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationalMode(input string) (*OperationalMode, error) {
	vals := map[string]OperationalMode{
		"disabled": OperationalModeDisabled,
		"enabled":  OperationalModeEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationalMode(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
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

type ResourceHealthState string

const (
	ResourceHealthStateAvailable   ResourceHealthState = "Available"
	ResourceHealthStateDegraded    ResourceHealthState = "Degraded"
	ResourceHealthStateUnavailable ResourceHealthState = "Unavailable"
	ResourceHealthStateUnknown     ResourceHealthState = "Unknown"
)

func PossibleValuesForResourceHealthState() []string {
	return []string{
		string(ResourceHealthStateAvailable),
		string(ResourceHealthStateDegraded),
		string(ResourceHealthStateUnavailable),
		string(ResourceHealthStateUnknown),
	}
}

func (s *ResourceHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceHealthState(input string) (*ResourceHealthState, error) {
	vals := map[string]ResourceHealthState{
		"available":   ResourceHealthStateAvailable,
		"degraded":    ResourceHealthStateDegraded,
		"unavailable": ResourceHealthStateUnavailable,
		"unknown":     ResourceHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthState(input)
	return &out, nil
}
