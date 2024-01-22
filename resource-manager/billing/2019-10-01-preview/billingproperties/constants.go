package billingproperties

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileSpendingLimit string

const (
	BillingProfileSpendingLimitOff BillingProfileSpendingLimit = "Off"
	BillingProfileSpendingLimitOn  BillingProfileSpendingLimit = "On"
)

func PossibleValuesForBillingProfileSpendingLimit() []string {
	return []string{
		string(BillingProfileSpendingLimitOff),
		string(BillingProfileSpendingLimitOn),
	}
}

func (s *BillingProfileSpendingLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileSpendingLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileSpendingLimit(input string) (*BillingProfileSpendingLimit, error) {
	vals := map[string]BillingProfileSpendingLimit{
		"off": BillingProfileSpendingLimitOff,
		"on":  BillingProfileSpendingLimitOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileSpendingLimit(input)
	return &out, nil
}

type BillingProfileStatus string

const (
	BillingProfileStatusActive   BillingProfileStatus = "Active"
	BillingProfileStatusDisabled BillingProfileStatus = "Disabled"
	BillingProfileStatusWarned   BillingProfileStatus = "Warned"
)

func PossibleValuesForBillingProfileStatus() []string {
	return []string{
		string(BillingProfileStatusActive),
		string(BillingProfileStatusDisabled),
		string(BillingProfileStatusWarned),
	}
}

func (s *BillingProfileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatus(input string) (*BillingProfileStatus, error) {
	vals := map[string]BillingProfileStatus{
		"active":   BillingProfileStatusActive,
		"disabled": BillingProfileStatusDisabled,
		"warned":   BillingProfileStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatus(input)
	return &out, nil
}

type BillingProfileStatusReasonCode string

const (
	BillingProfileStatusReasonCodePastDue              BillingProfileStatusReasonCode = "PastDue"
	BillingProfileStatusReasonCodeSpendingLimitExpired BillingProfileStatusReasonCode = "SpendingLimitExpired"
	BillingProfileStatusReasonCodeSpendingLimitReached BillingProfileStatusReasonCode = "SpendingLimitReached"
)

func PossibleValuesForBillingProfileStatusReasonCode() []string {
	return []string{
		string(BillingProfileStatusReasonCodePastDue),
		string(BillingProfileStatusReasonCodeSpendingLimitExpired),
		string(BillingProfileStatusReasonCodeSpendingLimitReached),
	}
}

func (s *BillingProfileStatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatusReasonCode(input string) (*BillingProfileStatusReasonCode, error) {
	vals := map[string]BillingProfileStatusReasonCode{
		"pastdue":              BillingProfileStatusReasonCodePastDue,
		"spendinglimitexpired": BillingProfileStatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": BillingProfileStatusReasonCodeSpendingLimitReached,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatusReasonCode(input)
	return &out, nil
}
