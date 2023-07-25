package autoscalevcores

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCoreProvisioningState string

const (
	VCoreProvisioningStateSucceeded VCoreProvisioningState = "Succeeded"
)

func PossibleValuesForVCoreProvisioningState() []string {
	return []string{
		string(VCoreProvisioningStateSucceeded),
	}
}

func (s *VCoreProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVCoreProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVCoreProvisioningState(input string) (*VCoreProvisioningState, error) {
	vals := map[string]VCoreProvisioningState{
		"succeeded": VCoreProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VCoreProvisioningState(input)
	return &out, nil
}

type VCoreSkuTier string

const (
	VCoreSkuTierAutoScale VCoreSkuTier = "AutoScale"
)

func PossibleValuesForVCoreSkuTier() []string {
	return []string{
		string(VCoreSkuTierAutoScale),
	}
}

func (s *VCoreSkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVCoreSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVCoreSkuTier(input string) (*VCoreSkuTier, error) {
	vals := map[string]VCoreSkuTier{
		"autoscale": VCoreSkuTierAutoScale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VCoreSkuTier(input)
	return &out, nil
}
