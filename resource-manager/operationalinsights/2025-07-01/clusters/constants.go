package clusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingType string

const (
	BillingTypeCluster    BillingType = "Cluster"
	BillingTypeWorkspaces BillingType = "Workspaces"
)

func PossibleValuesForBillingType() []string {
	return []string{
		string(BillingTypeCluster),
		string(BillingTypeWorkspaces),
	}
}

func (s *BillingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingType(input string) (*BillingType, error) {
	vals := map[string]BillingType{
		"cluster":    BillingTypeCluster,
		"workspaces": BillingTypeWorkspaces,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingType(input)
	return &out, nil
}

type ClusterEntityStatus string

const (
	ClusterEntityStatusCanceled            ClusterEntityStatus = "Canceled"
	ClusterEntityStatusCreating            ClusterEntityStatus = "Creating"
	ClusterEntityStatusDeleting            ClusterEntityStatus = "Deleting"
	ClusterEntityStatusFailed              ClusterEntityStatus = "Failed"
	ClusterEntityStatusProvisioningAccount ClusterEntityStatus = "ProvisioningAccount"
	ClusterEntityStatusSucceeded           ClusterEntityStatus = "Succeeded"
	ClusterEntityStatusUpdating            ClusterEntityStatus = "Updating"
)

func PossibleValuesForClusterEntityStatus() []string {
	return []string{
		string(ClusterEntityStatusCanceled),
		string(ClusterEntityStatusCreating),
		string(ClusterEntityStatusDeleting),
		string(ClusterEntityStatusFailed),
		string(ClusterEntityStatusProvisioningAccount),
		string(ClusterEntityStatusSucceeded),
		string(ClusterEntityStatusUpdating),
	}
}

func (s *ClusterEntityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterEntityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterEntityStatus(input string) (*ClusterEntityStatus, error) {
	vals := map[string]ClusterEntityStatus{
		"canceled":            ClusterEntityStatusCanceled,
		"creating":            ClusterEntityStatusCreating,
		"deleting":            ClusterEntityStatusDeleting,
		"failed":              ClusterEntityStatusFailed,
		"provisioningaccount": ClusterEntityStatusProvisioningAccount,
		"succeeded":           ClusterEntityStatusSucceeded,
		"updating":            ClusterEntityStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterEntityStatus(input)
	return &out, nil
}

type ClusterReplicationState string

const (
	ClusterReplicationStateCanceled          ClusterReplicationState = "Canceled"
	ClusterReplicationStateDisableRequested  ClusterReplicationState = "DisableRequested"
	ClusterReplicationStateDisabling         ClusterReplicationState = "Disabling"
	ClusterReplicationStateEnableRequested   ClusterReplicationState = "EnableRequested"
	ClusterReplicationStateEnabling          ClusterReplicationState = "Enabling"
	ClusterReplicationStateFailed            ClusterReplicationState = "Failed"
	ClusterReplicationStateRollbackRequested ClusterReplicationState = "RollbackRequested"
	ClusterReplicationStateRollingBack       ClusterReplicationState = "RollingBack"
	ClusterReplicationStateSucceeded         ClusterReplicationState = "Succeeded"
)

func PossibleValuesForClusterReplicationState() []string {
	return []string{
		string(ClusterReplicationStateCanceled),
		string(ClusterReplicationStateDisableRequested),
		string(ClusterReplicationStateDisabling),
		string(ClusterReplicationStateEnableRequested),
		string(ClusterReplicationStateEnabling),
		string(ClusterReplicationStateFailed),
		string(ClusterReplicationStateRollbackRequested),
		string(ClusterReplicationStateRollingBack),
		string(ClusterReplicationStateSucceeded),
	}
}

func (s *ClusterReplicationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterReplicationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterReplicationState(input string) (*ClusterReplicationState, error) {
	vals := map[string]ClusterReplicationState{
		"canceled":          ClusterReplicationStateCanceled,
		"disablerequested":  ClusterReplicationStateDisableRequested,
		"disabling":         ClusterReplicationStateDisabling,
		"enablerequested":   ClusterReplicationStateEnableRequested,
		"enabling":          ClusterReplicationStateEnabling,
		"failed":            ClusterReplicationStateFailed,
		"rollbackrequested": ClusterReplicationStateRollbackRequested,
		"rollingback":       ClusterReplicationStateRollingBack,
		"succeeded":         ClusterReplicationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterReplicationState(input)
	return &out, nil
}

type ClusterSkuNameEnum string

const (
	ClusterSkuNameEnumCapacityReservation ClusterSkuNameEnum = "CapacityReservation"
)

func PossibleValuesForClusterSkuNameEnum() []string {
	return []string{
		string(ClusterSkuNameEnumCapacityReservation),
	}
}

func (s *ClusterSkuNameEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClusterSkuNameEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClusterSkuNameEnum(input string) (*ClusterSkuNameEnum, error) {
	vals := map[string]ClusterSkuNameEnum{
		"capacityreservation": ClusterSkuNameEnumCapacityReservation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterSkuNameEnum(input)
	return &out, nil
}
