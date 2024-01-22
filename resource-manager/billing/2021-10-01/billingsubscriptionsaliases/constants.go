package billingsubscriptionsaliases

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoRenew string

const (
	AutoRenewOff AutoRenew = "Off"
	AutoRenewOn  AutoRenew = "On"
)

func PossibleValuesForAutoRenew() []string {
	return []string{
		string(AutoRenewOff),
		string(AutoRenewOn),
	}
}

func (s *AutoRenew) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoRenew(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoRenew(input string) (*AutoRenew, error) {
	vals := map[string]AutoRenew{
		"off": AutoRenewOff,
		"on":  AutoRenewOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoRenew(input)
	return &out, nil
}

type BillingSubscriptionStatus string

const (
	BillingSubscriptionStatusActive    BillingSubscriptionStatus = "Active"
	BillingSubscriptionStatusAutoRenew BillingSubscriptionStatus = "AutoRenew"
	BillingSubscriptionStatusCancelled BillingSubscriptionStatus = "Cancelled"
	BillingSubscriptionStatusDeleted   BillingSubscriptionStatus = "Deleted"
	BillingSubscriptionStatusDisabled  BillingSubscriptionStatus = "Disabled"
	BillingSubscriptionStatusExpired   BillingSubscriptionStatus = "Expired"
	BillingSubscriptionStatusExpiring  BillingSubscriptionStatus = "Expiring"
	BillingSubscriptionStatusSuspended BillingSubscriptionStatus = "Suspended"
	BillingSubscriptionStatusUnknown   BillingSubscriptionStatus = "Unknown"
	BillingSubscriptionStatusWarned    BillingSubscriptionStatus = "Warned"
)

func PossibleValuesForBillingSubscriptionStatus() []string {
	return []string{
		string(BillingSubscriptionStatusActive),
		string(BillingSubscriptionStatusAutoRenew),
		string(BillingSubscriptionStatusCancelled),
		string(BillingSubscriptionStatusDeleted),
		string(BillingSubscriptionStatusDisabled),
		string(BillingSubscriptionStatusExpired),
		string(BillingSubscriptionStatusExpiring),
		string(BillingSubscriptionStatusSuspended),
		string(BillingSubscriptionStatusUnknown),
		string(BillingSubscriptionStatusWarned),
	}
}

func (s *BillingSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSubscriptionStatus(input string) (*BillingSubscriptionStatus, error) {
	vals := map[string]BillingSubscriptionStatus{
		"active":    BillingSubscriptionStatusActive,
		"autorenew": BillingSubscriptionStatusAutoRenew,
		"cancelled": BillingSubscriptionStatusCancelled,
		"deleted":   BillingSubscriptionStatusDeleted,
		"disabled":  BillingSubscriptionStatusDisabled,
		"expired":   BillingSubscriptionStatusExpired,
		"expiring":  BillingSubscriptionStatusExpiring,
		"suspended": BillingSubscriptionStatusSuspended,
		"unknown":   BillingSubscriptionStatusUnknown,
		"warned":    BillingSubscriptionStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSubscriptionStatus(input)
	return &out, nil
}

type SubscriptionEnrollmentAccountStatus string

const (
	SubscriptionEnrollmentAccountStatusActive         SubscriptionEnrollmentAccountStatus = "Active"
	SubscriptionEnrollmentAccountStatusCancelled      SubscriptionEnrollmentAccountStatus = "Cancelled"
	SubscriptionEnrollmentAccountStatusDeleted        SubscriptionEnrollmentAccountStatus = "Deleted"
	SubscriptionEnrollmentAccountStatusExpired        SubscriptionEnrollmentAccountStatus = "Expired"
	SubscriptionEnrollmentAccountStatusTransferredOut SubscriptionEnrollmentAccountStatus = "TransferredOut"
	SubscriptionEnrollmentAccountStatusTransferring   SubscriptionEnrollmentAccountStatus = "Transferring"
)

func PossibleValuesForSubscriptionEnrollmentAccountStatus() []string {
	return []string{
		string(SubscriptionEnrollmentAccountStatusActive),
		string(SubscriptionEnrollmentAccountStatusCancelled),
		string(SubscriptionEnrollmentAccountStatusDeleted),
		string(SubscriptionEnrollmentAccountStatusExpired),
		string(SubscriptionEnrollmentAccountStatusTransferredOut),
		string(SubscriptionEnrollmentAccountStatusTransferring),
	}
}

func (s *SubscriptionEnrollmentAccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionEnrollmentAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionEnrollmentAccountStatus(input string) (*SubscriptionEnrollmentAccountStatus, error) {
	vals := map[string]SubscriptionEnrollmentAccountStatus{
		"active":         SubscriptionEnrollmentAccountStatusActive,
		"cancelled":      SubscriptionEnrollmentAccountStatusCancelled,
		"deleted":        SubscriptionEnrollmentAccountStatusDeleted,
		"expired":        SubscriptionEnrollmentAccountStatusExpired,
		"transferredout": SubscriptionEnrollmentAccountStatusTransferredOut,
		"transferring":   SubscriptionEnrollmentAccountStatusTransferring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionEnrollmentAccountStatus(input)
	return &out, nil
}
