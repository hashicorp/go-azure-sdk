package lots

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotSource string

const (
	LotSourceConsumptionCommitment LotSource = "ConsumptionCommitment"
	LotSourcePromotionalCredit     LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit       LotSource = "PurchasedCredit"
)

func PossibleValuesForLotSource() []string {
	return []string{
		string(LotSourceConsumptionCommitment),
		string(LotSourcePromotionalCredit),
		string(LotSourcePurchasedCredit),
	}
}

func (s *LotSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLotSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLotSource(input string) (*LotSource, error) {
	vals := map[string]LotSource{
		"consumptioncommitment": LotSourceConsumptionCommitment,
		"promotionalcredit":     LotSourcePromotionalCredit,
		"purchasedcredit":       LotSourcePurchasedCredit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LotSource(input)
	return &out, nil
}

type OrgType string

const (
	OrgTypeContributor OrgType = "Contributor"
	OrgTypePrimary     OrgType = "Primary"
)

func PossibleValuesForOrgType() []string {
	return []string{
		string(OrgTypeContributor),
		string(OrgTypePrimary),
	}
}

func (s *OrgType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrgType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrgType(input string) (*OrgType, error) {
	vals := map[string]OrgType{
		"contributor": OrgTypeContributor,
		"primary":     OrgTypePrimary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrgType(input)
	return &out, nil
}

type Status string

const (
	StatusActive   Status = "Active"
	StatusCanceled Status = "Canceled"
	StatusComplete Status = "Complete"
	StatusExpired  Status = "Expired"
	StatusInactive Status = "Inactive"
	StatusNone     Status = "None"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusActive),
		string(StatusCanceled),
		string(StatusComplete),
		string(StatusExpired),
		string(StatusInactive),
		string(StatusNone),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"active":   StatusActive,
		"canceled": StatusCanceled,
		"complete": StatusComplete,
		"expired":  StatusExpired,
		"inactive": StatusInactive,
		"none":     StatusNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
