package invoicesections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceSectionState string

const (
	InvoiceSectionStateActive     InvoiceSectionState = "Active"
	InvoiceSectionStateRestricted InvoiceSectionState = "Restricted"
)

func PossibleValuesForInvoiceSectionState() []string {
	return []string{
		string(InvoiceSectionStateActive),
		string(InvoiceSectionStateRestricted),
	}
}

func (s *InvoiceSectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvoiceSectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvoiceSectionState(input string) (*InvoiceSectionState, error) {
	vals := map[string]InvoiceSectionState{
		"active":     InvoiceSectionStateActive,
		"restricted": InvoiceSectionStateRestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvoiceSectionState(input)
	return &out, nil
}

type TargetCloud string

const (
	TargetCloudUSGov TargetCloud = "USGov"
	TargetCloudUSNat TargetCloud = "USNat"
	TargetCloudUSSec TargetCloud = "USSec"
)

func PossibleValuesForTargetCloud() []string {
	return []string{
		string(TargetCloudUSGov),
		string(TargetCloudUSNat),
		string(TargetCloudUSSec),
	}
}

func (s *TargetCloud) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetCloud(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetCloud(input string) (*TargetCloud, error) {
	vals := map[string]TargetCloud{
		"usgov": TargetCloudUSGov,
		"usnat": TargetCloudUSNat,
		"ussec": TargetCloudUSSec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetCloud(input)
	return &out, nil
}
