package billingrequest

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRequestStatus string

const (
	BillingRequestStatusApproved  BillingRequestStatus = "Approved"
	BillingRequestStatusCancelled BillingRequestStatus = "Cancelled"
	BillingRequestStatusCompleted BillingRequestStatus = "Completed"
	BillingRequestStatusDeclined  BillingRequestStatus = "Declined"
	BillingRequestStatusExpired   BillingRequestStatus = "Expired"
	BillingRequestStatusOther     BillingRequestStatus = "Other"
	BillingRequestStatusPending   BillingRequestStatus = "Pending"
)

func PossibleValuesForBillingRequestStatus() []string {
	return []string{
		string(BillingRequestStatusApproved),
		string(BillingRequestStatusCancelled),
		string(BillingRequestStatusCompleted),
		string(BillingRequestStatusDeclined),
		string(BillingRequestStatusExpired),
		string(BillingRequestStatusOther),
		string(BillingRequestStatusPending),
	}
}

func (s *BillingRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingRequestStatus(input string) (*BillingRequestStatus, error) {
	vals := map[string]BillingRequestStatus{
		"approved":  BillingRequestStatusApproved,
		"cancelled": BillingRequestStatusCancelled,
		"completed": BillingRequestStatusCompleted,
		"declined":  BillingRequestStatusDeclined,
		"expired":   BillingRequestStatusExpired,
		"other":     BillingRequestStatusOther,
		"pending":   BillingRequestStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRequestStatus(input)
	return &out, nil
}

type BillingRequestType string

const (
	BillingRequestTypeInvoiceAccess       BillingRequestType = "InvoiceAccess"
	BillingRequestTypeOther               BillingRequestType = "Other"
	BillingRequestTypeProvisioningAccess  BillingRequestType = "ProvisioningAccess"
	BillingRequestTypeRoleAssignment      BillingRequestType = "RoleAssignment"
	BillingRequestTypeUpdateBillingPolicy BillingRequestType = "UpdateBillingPolicy"
)

func PossibleValuesForBillingRequestType() []string {
	return []string{
		string(BillingRequestTypeInvoiceAccess),
		string(BillingRequestTypeOther),
		string(BillingRequestTypeProvisioningAccess),
		string(BillingRequestTypeRoleAssignment),
		string(BillingRequestTypeUpdateBillingPolicy),
	}
}

func (s *BillingRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingRequestType(input string) (*BillingRequestType, error) {
	vals := map[string]BillingRequestType{
		"invoiceaccess":       BillingRequestTypeInvoiceAccess,
		"other":               BillingRequestTypeOther,
		"provisioningaccess":  BillingRequestTypeProvisioningAccess,
		"roleassignment":      BillingRequestTypeRoleAssignment,
		"updatebillingpolicy": BillingRequestTypeUpdateBillingPolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRequestType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNew          ProvisioningState = "New"
	ProvisioningStatePending      ProvisioningState = "Pending"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNew),
		string(ProvisioningStatePending),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
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
		"new":          ProvisioningStateNew,
		"pending":      ProvisioningStatePending,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
