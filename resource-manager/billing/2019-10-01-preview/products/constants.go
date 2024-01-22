package products

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingFrequency string

const (
	BillingFrequencyMonthly    BillingFrequency = "Monthly"
	BillingFrequencyOneTime    BillingFrequency = "OneTime"
	BillingFrequencyUsageBased BillingFrequency = "UsageBased"
)

func PossibleValuesForBillingFrequency() []string {
	return []string{
		string(BillingFrequencyMonthly),
		string(BillingFrequencyOneTime),
		string(BillingFrequencyUsageBased),
	}
}

func (s *BillingFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingFrequency(input string) (*BillingFrequency, error) {
	vals := map[string]BillingFrequency{
		"monthly":    BillingFrequencyMonthly,
		"onetime":    BillingFrequencyOneTime,
		"usagebased": BillingFrequencyUsageBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingFrequency(input)
	return &out, nil
}

type ProductStatusType string

const (
	ProductStatusTypeActive    ProductStatusType = "Active"
	ProductStatusTypeAutoRenew ProductStatusType = "AutoRenew"
	ProductStatusTypeCancelled ProductStatusType = "Cancelled"
	ProductStatusTypeDisabled  ProductStatusType = "Disabled"
	ProductStatusTypeExpired   ProductStatusType = "Expired"
	ProductStatusTypeExpiring  ProductStatusType = "Expiring"
	ProductStatusTypeInactive  ProductStatusType = "Inactive"
	ProductStatusTypePastDue   ProductStatusType = "PastDue"
)

func PossibleValuesForProductStatusType() []string {
	return []string{
		string(ProductStatusTypeActive),
		string(ProductStatusTypeAutoRenew),
		string(ProductStatusTypeCancelled),
		string(ProductStatusTypeDisabled),
		string(ProductStatusTypeExpired),
		string(ProductStatusTypeExpiring),
		string(ProductStatusTypeInactive),
		string(ProductStatusTypePastDue),
	}
}

func (s *ProductStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProductStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProductStatusType(input string) (*ProductStatusType, error) {
	vals := map[string]ProductStatusType{
		"active":    ProductStatusTypeActive,
		"autorenew": ProductStatusTypeAutoRenew,
		"cancelled": ProductStatusTypeCancelled,
		"disabled":  ProductStatusTypeDisabled,
		"expired":   ProductStatusTypeExpired,
		"expiring":  ProductStatusTypeExpiring,
		"inactive":  ProductStatusTypeInactive,
		"pastdue":   ProductStatusTypePastDue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductStatusType(input)
	return &out, nil
}

type UpdateAutoRenew string

const (
	UpdateAutoRenewFalse UpdateAutoRenew = "false"
	UpdateAutoRenewTrue  UpdateAutoRenew = "true"
)

func PossibleValuesForUpdateAutoRenew() []string {
	return []string{
		string(UpdateAutoRenewFalse),
		string(UpdateAutoRenewTrue),
	}
}

func (s *UpdateAutoRenew) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateAutoRenew(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateAutoRenew(input string) (*UpdateAutoRenew, error) {
	vals := map[string]UpdateAutoRenew{
		"false": UpdateAutoRenewFalse,
		"true":  UpdateAutoRenewTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateAutoRenew(input)
	return &out, nil
}
