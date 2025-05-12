package provisionednetworks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedNetworkProvisioningState string

const (
	ProvisionedNetworkProvisioningStateCanceled  ProvisionedNetworkProvisioningState = "Canceled"
	ProvisionedNetworkProvisioningStateFailed    ProvisionedNetworkProvisioningState = "Failed"
	ProvisionedNetworkProvisioningStateSucceeded ProvisionedNetworkProvisioningState = "Succeeded"
)

func PossibleValuesForProvisionedNetworkProvisioningState() []string {
	return []string{
		string(ProvisionedNetworkProvisioningStateCanceled),
		string(ProvisionedNetworkProvisioningStateFailed),
		string(ProvisionedNetworkProvisioningStateSucceeded),
	}
}

func (s *ProvisionedNetworkProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisionedNetworkProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisionedNetworkProvisioningState(input string) (*ProvisionedNetworkProvisioningState, error) {
	vals := map[string]ProvisionedNetworkProvisioningState{
		"canceled":  ProvisionedNetworkProvisioningStateCanceled,
		"failed":    ProvisionedNetworkProvisioningStateFailed,
		"succeeded": ProvisionedNetworkProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisionedNetworkProvisioningState(input)
	return &out, nil
}

type ProvisionedNetworkTypes string

const (
	ProvisionedNetworkTypesEsxManagement     ProvisionedNetworkTypes = "esxManagement"
	ProvisionedNetworkTypesEsxReplication    ProvisionedNetworkTypes = "esxReplication"
	ProvisionedNetworkTypesHcxManagement     ProvisionedNetworkTypes = "hcxManagement"
	ProvisionedNetworkTypesHcxUplink         ProvisionedNetworkTypes = "hcxUplink"
	ProvisionedNetworkTypesVMotion           ProvisionedNetworkTypes = "vmotion"
	ProvisionedNetworkTypesVcenterManagement ProvisionedNetworkTypes = "vcenterManagement"
	ProvisionedNetworkTypesVsan              ProvisionedNetworkTypes = "vsan"
)

func PossibleValuesForProvisionedNetworkTypes() []string {
	return []string{
		string(ProvisionedNetworkTypesEsxManagement),
		string(ProvisionedNetworkTypesEsxReplication),
		string(ProvisionedNetworkTypesHcxManagement),
		string(ProvisionedNetworkTypesHcxUplink),
		string(ProvisionedNetworkTypesVMotion),
		string(ProvisionedNetworkTypesVcenterManagement),
		string(ProvisionedNetworkTypesVsan),
	}
}

func (s *ProvisionedNetworkTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisionedNetworkTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisionedNetworkTypes(input string) (*ProvisionedNetworkTypes, error) {
	vals := map[string]ProvisionedNetworkTypes{
		"esxmanagement":     ProvisionedNetworkTypesEsxManagement,
		"esxreplication":    ProvisionedNetworkTypesEsxReplication,
		"hcxmanagement":     ProvisionedNetworkTypesHcxManagement,
		"hcxuplink":         ProvisionedNetworkTypesHcxUplink,
		"vmotion":           ProvisionedNetworkTypesVMotion,
		"vcentermanagement": ProvisionedNetworkTypesVcenterManagement,
		"vsan":              ProvisionedNetworkTypesVsan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisionedNetworkTypes(input)
	return &out, nil
}
