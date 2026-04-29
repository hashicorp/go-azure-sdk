package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Filter string

const (
	FilterResourceTypeEqMicrosoftPointKeyVaultVaults Filter = "resourceType eq 'Microsoft.KeyVault/vaults'"
)

func PossibleValuesForFilter() []string {
	return []string{
		string(FilterResourceTypeEqMicrosoftPointKeyVaultVaults),
	}
}

func (s *Filter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilter(input string) (*Filter, error) {
	vals := map[string]Filter{
		"resourcetype eq 'microsoft.keyvault/vaults'": FilterResourceTypeEqMicrosoftPointKeyVaultVaults,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Filter(input)
	return &out, nil
}

type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAccountNameInvalid),
		string(ReasonAlreadyExists),
	}
}

func (s *Reason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"accountnameinvalid": ReasonAccountNameInvalid,
		"alreadyexists":      ReasonAlreadyExists,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type VaultCheckNameAvailabilityParametersType string

const (
	VaultCheckNameAvailabilityParametersTypeMicrosoftPointKeyVaultVaults VaultCheckNameAvailabilityParametersType = "Microsoft.KeyVault/vaults"
)

func PossibleValuesForVaultCheckNameAvailabilityParametersType() []string {
	return []string{
		string(VaultCheckNameAvailabilityParametersTypeMicrosoftPointKeyVaultVaults),
	}
}

func (s *VaultCheckNameAvailabilityParametersType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVaultCheckNameAvailabilityParametersType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVaultCheckNameAvailabilityParametersType(input string) (*VaultCheckNameAvailabilityParametersType, error) {
	vals := map[string]VaultCheckNameAvailabilityParametersType{
		"microsoft.keyvault/vaults": VaultCheckNameAvailabilityParametersTypeMicrosoftPointKeyVaultVaults,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VaultCheckNameAvailabilityParametersType(input)
	return &out, nil
}
