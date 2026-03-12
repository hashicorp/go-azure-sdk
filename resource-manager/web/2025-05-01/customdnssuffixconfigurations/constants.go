package customdnssuffixconfigurations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDnsSuffixProvisioningState string

const (
	CustomDnsSuffixProvisioningStateDegraded   CustomDnsSuffixProvisioningState = "Degraded"
	CustomDnsSuffixProvisioningStateFailed     CustomDnsSuffixProvisioningState = "Failed"
	CustomDnsSuffixProvisioningStateInProgress CustomDnsSuffixProvisioningState = "InProgress"
	CustomDnsSuffixProvisioningStateSucceeded  CustomDnsSuffixProvisioningState = "Succeeded"
)

func PossibleValuesForCustomDnsSuffixProvisioningState() []string {
	return []string{
		string(CustomDnsSuffixProvisioningStateDegraded),
		string(CustomDnsSuffixProvisioningStateFailed),
		string(CustomDnsSuffixProvisioningStateInProgress),
		string(CustomDnsSuffixProvisioningStateSucceeded),
	}
}

func (s *CustomDnsSuffixProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomDnsSuffixProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomDnsSuffixProvisioningState(input string) (*CustomDnsSuffixProvisioningState, error) {
	vals := map[string]CustomDnsSuffixProvisioningState{
		"degraded":   CustomDnsSuffixProvisioningStateDegraded,
		"failed":     CustomDnsSuffixProvisioningStateFailed,
		"inprogress": CustomDnsSuffixProvisioningStateInProgress,
		"succeeded":  CustomDnsSuffixProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomDnsSuffixProvisioningState(input)
	return &out, nil
}
