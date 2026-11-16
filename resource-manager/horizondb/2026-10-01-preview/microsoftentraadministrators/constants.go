package microsoftentraadministrators

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrincipalTypes string

const (
	PrincipalTypesGroup            PrincipalTypes = "Group"
	PrincipalTypesServicePrincipal PrincipalTypes = "ServicePrincipal"
	PrincipalTypesUnknown          PrincipalTypes = "Unknown"
	PrincipalTypesUser             PrincipalTypes = "User"
)

func PossibleValuesForPrincipalTypes() []string {
	return []string{
		string(PrincipalTypesGroup),
		string(PrincipalTypesServicePrincipal),
		string(PrincipalTypesUnknown),
		string(PrincipalTypesUser),
	}
}

func (s *PrincipalTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrincipalTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrincipalTypes(input string) (*PrincipalTypes, error) {
	vals := map[string]PrincipalTypes{
		"group":            PrincipalTypesGroup,
		"serviceprincipal": PrincipalTypesServicePrincipal,
		"unknown":          PrincipalTypesUnknown,
		"user":             PrincipalTypesUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrincipalTypes(input)
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
