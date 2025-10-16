package brokerauthentication

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerAuthenticationMethod string

const (
	BrokerAuthenticationMethodCustom              BrokerAuthenticationMethod = "Custom"
	BrokerAuthenticationMethodServiceAccountToken BrokerAuthenticationMethod = "ServiceAccountToken"
	BrokerAuthenticationMethodXFiveZeroNine       BrokerAuthenticationMethod = "X509"
)

func PossibleValuesForBrokerAuthenticationMethod() []string {
	return []string{
		string(BrokerAuthenticationMethodCustom),
		string(BrokerAuthenticationMethodServiceAccountToken),
		string(BrokerAuthenticationMethodXFiveZeroNine),
	}
}

func (s *BrokerAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrokerAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrokerAuthenticationMethod(input string) (*BrokerAuthenticationMethod, error) {
	vals := map[string]BrokerAuthenticationMethod{
		"custom":              BrokerAuthenticationMethodCustom,
		"serviceaccounttoken": BrokerAuthenticationMethodServiceAccountToken,
		"x509":                BrokerAuthenticationMethodXFiveZeroNine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrokerAuthenticationMethod(input)
	return &out, nil
}

type BrokerAuthenticatorValidationMethods string

const (
	BrokerAuthenticatorValidationMethodsAzureDeviceRegistry BrokerAuthenticatorValidationMethods = "AzureDeviceRegistry"
	BrokerAuthenticatorValidationMethodsNone                BrokerAuthenticatorValidationMethods = "None"
)

func PossibleValuesForBrokerAuthenticatorValidationMethods() []string {
	return []string{
		string(BrokerAuthenticatorValidationMethodsAzureDeviceRegistry),
		string(BrokerAuthenticatorValidationMethodsNone),
	}
}

func (s *BrokerAuthenticatorValidationMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrokerAuthenticatorValidationMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrokerAuthenticatorValidationMethods(input string) (*BrokerAuthenticatorValidationMethods, error) {
	vals := map[string]BrokerAuthenticatorValidationMethods{
		"azuredeviceregistry": BrokerAuthenticatorValidationMethodsAzureDeviceRegistry,
		"none":                BrokerAuthenticatorValidationMethodsNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrokerAuthenticatorValidationMethods(input)
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
