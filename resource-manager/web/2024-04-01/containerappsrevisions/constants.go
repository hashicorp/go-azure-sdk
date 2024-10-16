package containerappsrevisions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevisionHealthState string

const (
	RevisionHealthStateHealthy   RevisionHealthState = "Healthy"
	RevisionHealthStateNone      RevisionHealthState = "None"
	RevisionHealthStateUnhealthy RevisionHealthState = "Unhealthy"
)

func PossibleValuesForRevisionHealthState() []string {
	return []string{
		string(RevisionHealthStateHealthy),
		string(RevisionHealthStateNone),
		string(RevisionHealthStateUnhealthy),
	}
}

func (s *RevisionHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRevisionHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRevisionHealthState(input string) (*RevisionHealthState, error) {
	vals := map[string]RevisionHealthState{
		"healthy":   RevisionHealthStateHealthy,
		"none":      RevisionHealthStateNone,
		"unhealthy": RevisionHealthStateUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RevisionHealthState(input)
	return &out, nil
}

type RevisionProvisioningState string

const (
	RevisionProvisioningStateDeprovisioned  RevisionProvisioningState = "Deprovisioned"
	RevisionProvisioningStateDeprovisioning RevisionProvisioningState = "Deprovisioning"
	RevisionProvisioningStateFailed         RevisionProvisioningState = "Failed"
	RevisionProvisioningStateProvisioned    RevisionProvisioningState = "Provisioned"
	RevisionProvisioningStateProvisioning   RevisionProvisioningState = "Provisioning"
)

func PossibleValuesForRevisionProvisioningState() []string {
	return []string{
		string(RevisionProvisioningStateDeprovisioned),
		string(RevisionProvisioningStateDeprovisioning),
		string(RevisionProvisioningStateFailed),
		string(RevisionProvisioningStateProvisioned),
		string(RevisionProvisioningStateProvisioning),
	}
}

func (s *RevisionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRevisionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRevisionProvisioningState(input string) (*RevisionProvisioningState, error) {
	vals := map[string]RevisionProvisioningState{
		"deprovisioned":  RevisionProvisioningStateDeprovisioned,
		"deprovisioning": RevisionProvisioningStateDeprovisioning,
		"failed":         RevisionProvisioningStateFailed,
		"provisioned":    RevisionProvisioningStateProvisioned,
		"provisioning":   RevisionProvisioningStateProvisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RevisionProvisioningState(input)
	return &out, nil
}
