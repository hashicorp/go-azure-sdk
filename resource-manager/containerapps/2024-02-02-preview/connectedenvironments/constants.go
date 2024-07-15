package connectedenvironments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateProvisioningState string

const (
	CertificateProvisioningStateCanceled     CertificateProvisioningState = "Canceled"
	CertificateProvisioningStateDeleteFailed CertificateProvisioningState = "DeleteFailed"
	CertificateProvisioningStateFailed       CertificateProvisioningState = "Failed"
	CertificateProvisioningStatePending      CertificateProvisioningState = "Pending"
	CertificateProvisioningStateSucceeded    CertificateProvisioningState = "Succeeded"
)

func PossibleValuesForCertificateProvisioningState() []string {
	return []string{
		string(CertificateProvisioningStateCanceled),
		string(CertificateProvisioningStateDeleteFailed),
		string(CertificateProvisioningStateFailed),
		string(CertificateProvisioningStatePending),
		string(CertificateProvisioningStateSucceeded),
	}
}

func (s *CertificateProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateProvisioningState(input string) (*CertificateProvisioningState, error) {
	vals := map[string]CertificateProvisioningState{
		"canceled":     CertificateProvisioningStateCanceled,
		"deletefailed": CertificateProvisioningStateDeleteFailed,
		"failed":       CertificateProvisioningStateFailed,
		"pending":      CertificateProvisioningStatePending,
		"succeeded":    CertificateProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateProvisioningState(input)
	return &out, nil
}

type CertificateType string

const (
	CertificateTypeImagePullTrustedCA   CertificateType = "ImagePullTrustedCA"
	CertificateTypeServerSSLCertificate CertificateType = "ServerSSLCertificate"
)

func PossibleValuesForCertificateType() []string {
	return []string{
		string(CertificateTypeImagePullTrustedCA),
		string(CertificateTypeServerSSLCertificate),
	}
}

func (s *CertificateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateType(input string) (*CertificateType, error) {
	vals := map[string]CertificateType{
		"imagepulltrustedca":   CertificateTypeImagePullTrustedCA,
		"serversslcertificate": CertificateTypeServerSSLCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateType(input)
	return &out, nil
}

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

func (s *CheckNameAvailabilityReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckNameAvailabilityReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckNameAvailabilityReason(input string) (*CheckNameAvailabilityReason, error) {
	vals := map[string]CheckNameAvailabilityReason{
		"alreadyexists": CheckNameAvailabilityReasonAlreadyExists,
		"invalid":       CheckNameAvailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityReason(input)
	return &out, nil
}

type ConnectedEnvironmentProvisioningState string

const (
	ConnectedEnvironmentProvisioningStateCanceled                      ConnectedEnvironmentProvisioningState = "Canceled"
	ConnectedEnvironmentProvisioningStateFailed                        ConnectedEnvironmentProvisioningState = "Failed"
	ConnectedEnvironmentProvisioningStateInfrastructureSetupComplete   ConnectedEnvironmentProvisioningState = "InfrastructureSetupComplete"
	ConnectedEnvironmentProvisioningStateInfrastructureSetupInProgress ConnectedEnvironmentProvisioningState = "InfrastructureSetupInProgress"
	ConnectedEnvironmentProvisioningStateInitializationInProgress      ConnectedEnvironmentProvisioningState = "InitializationInProgress"
	ConnectedEnvironmentProvisioningStateScheduledForDelete            ConnectedEnvironmentProvisioningState = "ScheduledForDelete"
	ConnectedEnvironmentProvisioningStateSucceeded                     ConnectedEnvironmentProvisioningState = "Succeeded"
	ConnectedEnvironmentProvisioningStateWaiting                       ConnectedEnvironmentProvisioningState = "Waiting"
)

func PossibleValuesForConnectedEnvironmentProvisioningState() []string {
	return []string{
		string(ConnectedEnvironmentProvisioningStateCanceled),
		string(ConnectedEnvironmentProvisioningStateFailed),
		string(ConnectedEnvironmentProvisioningStateInfrastructureSetupComplete),
		string(ConnectedEnvironmentProvisioningStateInfrastructureSetupInProgress),
		string(ConnectedEnvironmentProvisioningStateInitializationInProgress),
		string(ConnectedEnvironmentProvisioningStateScheduledForDelete),
		string(ConnectedEnvironmentProvisioningStateSucceeded),
		string(ConnectedEnvironmentProvisioningStateWaiting),
	}
}

func (s *ConnectedEnvironmentProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectedEnvironmentProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectedEnvironmentProvisioningState(input string) (*ConnectedEnvironmentProvisioningState, error) {
	vals := map[string]ConnectedEnvironmentProvisioningState{
		"canceled":                      ConnectedEnvironmentProvisioningStateCanceled,
		"failed":                        ConnectedEnvironmentProvisioningStateFailed,
		"infrastructuresetupcomplete":   ConnectedEnvironmentProvisioningStateInfrastructureSetupComplete,
		"infrastructuresetupinprogress": ConnectedEnvironmentProvisioningStateInfrastructureSetupInProgress,
		"initializationinprogress":      ConnectedEnvironmentProvisioningStateInitializationInProgress,
		"scheduledfordelete":            ConnectedEnvironmentProvisioningStateScheduledForDelete,
		"succeeded":                     ConnectedEnvironmentProvisioningStateSucceeded,
		"waiting":                       ConnectedEnvironmentProvisioningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectedEnvironmentProvisioningState(input)
	return &out, nil
}

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation ExtendedLocationTypes = "CustomLocation"
)

func PossibleValuesForExtendedLocationTypes() []string {
	return []string{
		string(ExtendedLocationTypesCustomLocation),
	}
}

func (s *ExtendedLocationTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExtendedLocationTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExtendedLocationTypes(input string) (*ExtendedLocationTypes, error) {
	vals := map[string]ExtendedLocationTypes{
		"customlocation": ExtendedLocationTypesCustomLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationTypes(input)
	return &out, nil
}
