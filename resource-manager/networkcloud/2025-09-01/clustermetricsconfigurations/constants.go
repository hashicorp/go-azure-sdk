package clustermetricsconfigurations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterMetricsConfigurationDetailedStatus string

const (
	ClusterMetricsConfigurationDetailedStatusApplied    ClusterMetricsConfigurationDetailedStatus = "Applied"
	ClusterMetricsConfigurationDetailedStatusError      ClusterMetricsConfigurationDetailedStatus = "Error"
	ClusterMetricsConfigurationDetailedStatusProcessing ClusterMetricsConfigurationDetailedStatus = "Processing"
)

func PossibleValuesForClusterMetricsConfigurationDetailedStatus() []string {
	return []string{
		string(ClusterMetricsConfigurationDetailedStatusApplied),
		string(ClusterMetricsConfigurationDetailedStatusError),
		string(ClusterMetricsConfigurationDetailedStatusProcessing),
	}
}

func (s *ClusterMetricsConfigurationDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterMetricsConfigurationDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterMetricsConfigurationDetailedStatus(input string) (*ClusterMetricsConfigurationDetailedStatus, error) {
	vals := map[string]ClusterMetricsConfigurationDetailedStatus{
		"applied":    ClusterMetricsConfigurationDetailedStatusApplied,
		"error":      ClusterMetricsConfigurationDetailedStatusError,
		"processing": ClusterMetricsConfigurationDetailedStatusProcessing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterMetricsConfigurationDetailedStatus(input)
	return &out, nil
}

type ClusterMetricsConfigurationProvisioningState string

const (
	ClusterMetricsConfigurationProvisioningStateAccepted     ClusterMetricsConfigurationProvisioningState = "Accepted"
	ClusterMetricsConfigurationProvisioningStateCanceled     ClusterMetricsConfigurationProvisioningState = "Canceled"
	ClusterMetricsConfigurationProvisioningStateFailed       ClusterMetricsConfigurationProvisioningState = "Failed"
	ClusterMetricsConfigurationProvisioningStateProvisioning ClusterMetricsConfigurationProvisioningState = "Provisioning"
	ClusterMetricsConfigurationProvisioningStateSucceeded    ClusterMetricsConfigurationProvisioningState = "Succeeded"
)

func PossibleValuesForClusterMetricsConfigurationProvisioningState() []string {
	return []string{
		string(ClusterMetricsConfigurationProvisioningStateAccepted),
		string(ClusterMetricsConfigurationProvisioningStateCanceled),
		string(ClusterMetricsConfigurationProvisioningStateFailed),
		string(ClusterMetricsConfigurationProvisioningStateProvisioning),
		string(ClusterMetricsConfigurationProvisioningStateSucceeded),
	}
}

func (s *ClusterMetricsConfigurationProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterMetricsConfigurationProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterMetricsConfigurationProvisioningState(input string) (*ClusterMetricsConfigurationProvisioningState, error) {
	vals := map[string]ClusterMetricsConfigurationProvisioningState{
		"accepted":     ClusterMetricsConfigurationProvisioningStateAccepted,
		"canceled":     ClusterMetricsConfigurationProvisioningStateCanceled,
		"failed":       ClusterMetricsConfigurationProvisioningStateFailed,
		"provisioning": ClusterMetricsConfigurationProvisioningStateProvisioning,
		"succeeded":    ClusterMetricsConfigurationProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterMetricsConfigurationProvisioningState(input)
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
