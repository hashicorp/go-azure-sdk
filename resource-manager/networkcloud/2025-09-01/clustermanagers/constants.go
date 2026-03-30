package clustermanagers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterManagerDetailedStatus string

const (
	ClusterManagerDetailedStatusAvailable          ClusterManagerDetailedStatus = "Available"
	ClusterManagerDetailedStatusError              ClusterManagerDetailedStatus = "Error"
	ClusterManagerDetailedStatusProvisioning       ClusterManagerDetailedStatus = "Provisioning"
	ClusterManagerDetailedStatusProvisioningFailed ClusterManagerDetailedStatus = "ProvisioningFailed"
	ClusterManagerDetailedStatusUpdateFailed       ClusterManagerDetailedStatus = "UpdateFailed"
	ClusterManagerDetailedStatusUpdating           ClusterManagerDetailedStatus = "Updating"
)

func PossibleValuesForClusterManagerDetailedStatus() []string {
	return []string{
		string(ClusterManagerDetailedStatusAvailable),
		string(ClusterManagerDetailedStatusError),
		string(ClusterManagerDetailedStatusProvisioning),
		string(ClusterManagerDetailedStatusProvisioningFailed),
		string(ClusterManagerDetailedStatusUpdateFailed),
		string(ClusterManagerDetailedStatusUpdating),
	}
}

func (s *ClusterManagerDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterManagerDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterManagerDetailedStatus(input string) (*ClusterManagerDetailedStatus, error) {
	vals := map[string]ClusterManagerDetailedStatus{
		"available":          ClusterManagerDetailedStatusAvailable,
		"error":              ClusterManagerDetailedStatusError,
		"provisioning":       ClusterManagerDetailedStatusProvisioning,
		"provisioningfailed": ClusterManagerDetailedStatusProvisioningFailed,
		"updatefailed":       ClusterManagerDetailedStatusUpdateFailed,
		"updating":           ClusterManagerDetailedStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerDetailedStatus(input)
	return &out, nil
}

type ClusterManagerProvisioningState string

const (
	ClusterManagerProvisioningStateAccepted     ClusterManagerProvisioningState = "Accepted"
	ClusterManagerProvisioningStateCanceled     ClusterManagerProvisioningState = "Canceled"
	ClusterManagerProvisioningStateFailed       ClusterManagerProvisioningState = "Failed"
	ClusterManagerProvisioningStateProvisioning ClusterManagerProvisioningState = "Provisioning"
	ClusterManagerProvisioningStateSucceeded    ClusterManagerProvisioningState = "Succeeded"
	ClusterManagerProvisioningStateUpdating     ClusterManagerProvisioningState = "Updating"
)

func PossibleValuesForClusterManagerProvisioningState() []string {
	return []string{
		string(ClusterManagerProvisioningStateAccepted),
		string(ClusterManagerProvisioningStateCanceled),
		string(ClusterManagerProvisioningStateFailed),
		string(ClusterManagerProvisioningStateProvisioning),
		string(ClusterManagerProvisioningStateSucceeded),
		string(ClusterManagerProvisioningStateUpdating),
	}
}

func (s *ClusterManagerProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterManagerProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterManagerProvisioningState(input string) (*ClusterManagerProvisioningState, error) {
	vals := map[string]ClusterManagerProvisioningState{
		"accepted":     ClusterManagerProvisioningStateAccepted,
		"canceled":     ClusterManagerProvisioningStateCanceled,
		"failed":       ClusterManagerProvisioningStateFailed,
		"provisioning": ClusterManagerProvisioningStateProvisioning,
		"succeeded":    ClusterManagerProvisioningStateSucceeded,
		"updating":     ClusterManagerProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerProvisioningState(input)
	return &out, nil
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeCustomLocation ExtendedLocationType = "CustomLocation"
	ExtendedLocationTypeEdgeZone       ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeCustomLocation),
		string(ExtendedLocationTypeEdgeZone),
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
		"edgezone":       ExtendedLocationTypeEdgeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtendedLocationType(input)
	return &out, nil
}
