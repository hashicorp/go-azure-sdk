package capabilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeProcessorType string

const (
	ComputeProcessorTypeAMD   ComputeProcessorType = "AMD"
	ComputeProcessorTypeIntel ComputeProcessorType = "Intel"
)

func PossibleValuesForComputeProcessorType() []string {
	return []string{
		string(ComputeProcessorTypeAMD),
		string(ComputeProcessorTypeIntel),
	}
}

func (s *ComputeProcessorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeProcessorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeProcessorType(input string) (*ComputeProcessorType, error) {
	vals := map[string]ComputeProcessorType{
		"amd":   ComputeProcessorTypeAMD,
		"intel": ComputeProcessorTypeIntel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeProcessorType(input)
	return &out, nil
}

type PoolComputeModelType string

const (
	PoolComputeModelTypeProvisioned PoolComputeModelType = "Provisioned"
)

func PossibleValuesForPoolComputeModelType() []string {
	return []string{
		string(PoolComputeModelTypeProvisioned),
	}
}

func (s *PoolComputeModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolComputeModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolComputeModelType(input string) (*PoolComputeModelType, error) {
	vals := map[string]PoolComputeModelType{
		"provisioned": PoolComputeModelTypeProvisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolComputeModelType(input)
	return &out, nil
}

type PostgreSqlCapabilityVersion string

const (
	PostgreSqlCapabilityVersionOneEight PostgreSqlCapabilityVersion = "18"
	PostgreSqlCapabilityVersionOneSeven PostgreSqlCapabilityVersion = "17"
)

func PossibleValuesForPostgreSqlCapabilityVersion() []string {
	return []string{
		string(PostgreSqlCapabilityVersionOneEight),
		string(PostgreSqlCapabilityVersionOneSeven),
	}
}

func (s *PostgreSqlCapabilityVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePostgreSqlCapabilityVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePostgreSqlCapabilityVersion(input string) (*PostgreSqlCapabilityVersion, error) {
	vals := map[string]PostgreSqlCapabilityVersion{
		"18": PostgreSqlCapabilityVersionOneEight,
		"17": PostgreSqlCapabilityVersionOneSeven,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostgreSqlCapabilityVersion(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateInProgress   ProvisioningState = "InProgress"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
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
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"inprogress":   ProvisioningStateInProgress,
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
