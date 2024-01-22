package policies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplacePurchasesPolicy string

const (
	MarketplacePurchasesPolicyAllAllowed      MarketplacePurchasesPolicy = "AllAllowed"
	MarketplacePurchasesPolicyNotAllowed      MarketplacePurchasesPolicy = "NotAllowed"
	MarketplacePurchasesPolicyOnlyFreeAllowed MarketplacePurchasesPolicy = "OnlyFreeAllowed"
)

func PossibleValuesForMarketplacePurchasesPolicy() []string {
	return []string{
		string(MarketplacePurchasesPolicyAllAllowed),
		string(MarketplacePurchasesPolicyNotAllowed),
		string(MarketplacePurchasesPolicyOnlyFreeAllowed),
	}
}

func (s *MarketplacePurchasesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMarketplacePurchasesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMarketplacePurchasesPolicy(input string) (*MarketplacePurchasesPolicy, error) {
	vals := map[string]MarketplacePurchasesPolicy{
		"allallowed":      MarketplacePurchasesPolicyAllAllowed,
		"notallowed":      MarketplacePurchasesPolicyNotAllowed,
		"onlyfreeallowed": MarketplacePurchasesPolicyOnlyFreeAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MarketplacePurchasesPolicy(input)
	return &out, nil
}

type ReservationPurchasesPolicy string

const (
	ReservationPurchasesPolicyAllowed    ReservationPurchasesPolicy = "Allowed"
	ReservationPurchasesPolicyNotAllowed ReservationPurchasesPolicy = "NotAllowed"
)

func PossibleValuesForReservationPurchasesPolicy() []string {
	return []string{
		string(ReservationPurchasesPolicyAllowed),
		string(ReservationPurchasesPolicyNotAllowed),
	}
}

func (s *ReservationPurchasesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationPurchasesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationPurchasesPolicy(input string) (*ReservationPurchasesPolicy, error) {
	vals := map[string]ReservationPurchasesPolicy{
		"allowed":    ReservationPurchasesPolicyAllowed,
		"notallowed": ReservationPurchasesPolicyNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationPurchasesPolicy(input)
	return &out, nil
}

type ViewCharges string

const (
	ViewChargesAllowed    ViewCharges = "Allowed"
	ViewChargesNotAllowed ViewCharges = "NotAllowed"
)

func PossibleValuesForViewCharges() []string {
	return []string{
		string(ViewChargesAllowed),
		string(ViewChargesNotAllowed),
	}
}

func (s *ViewCharges) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseViewCharges(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseViewCharges(input string) (*ViewCharges, error) {
	vals := map[string]ViewCharges{
		"allowed":    ViewChargesAllowed,
		"notallowed": ViewChargesNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ViewCharges(input)
	return &out, nil
}

type ViewChargesPolicy string

const (
	ViewChargesPolicyAllowed    ViewChargesPolicy = "Allowed"
	ViewChargesPolicyNotAllowed ViewChargesPolicy = "NotAllowed"
)

func PossibleValuesForViewChargesPolicy() []string {
	return []string{
		string(ViewChargesPolicyAllowed),
		string(ViewChargesPolicyNotAllowed),
	}
}

func (s *ViewChargesPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseViewChargesPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseViewChargesPolicy(input string) (*ViewChargesPolicy, error) {
	vals := map[string]ViewChargesPolicy{
		"allowed":    ViewChargesPolicyAllowed,
		"notallowed": ViewChargesPolicyNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ViewChargesPolicy(input)
	return &out, nil
}
