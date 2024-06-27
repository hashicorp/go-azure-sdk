package associatedtenant

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingManagementTenantState string

const (
	BillingManagementTenantStateActive     BillingManagementTenantState = "Active"
	BillingManagementTenantStateNotAllowed BillingManagementTenantState = "NotAllowed"
	BillingManagementTenantStateOther      BillingManagementTenantState = "Other"
	BillingManagementTenantStateRevoked    BillingManagementTenantState = "Revoked"
)

func PossibleValuesForBillingManagementTenantState() []string {
	return []string{
		string(BillingManagementTenantStateActive),
		string(BillingManagementTenantStateNotAllowed),
		string(BillingManagementTenantStateOther),
		string(BillingManagementTenantStateRevoked),
	}
}

func (s *BillingManagementTenantState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingManagementTenantState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingManagementTenantState(input string) (*BillingManagementTenantState, error) {
	vals := map[string]BillingManagementTenantState{
		"active":     BillingManagementTenantStateActive,
		"notallowed": BillingManagementTenantStateNotAllowed,
		"other":      BillingManagementTenantStateOther,
		"revoked":    BillingManagementTenantStateRevoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingManagementTenantState(input)
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

type ProvisioningTenantState string

const (
	ProvisioningTenantStateActive                 ProvisioningTenantState = "Active"
	ProvisioningTenantStateBillingRequestDeclined ProvisioningTenantState = "BillingRequestDeclined"
	ProvisioningTenantStateBillingRequestExpired  ProvisioningTenantState = "BillingRequestExpired"
	ProvisioningTenantStateNotRequested           ProvisioningTenantState = "NotRequested"
	ProvisioningTenantStateOther                  ProvisioningTenantState = "Other"
	ProvisioningTenantStatePending                ProvisioningTenantState = "Pending"
	ProvisioningTenantStateRevoked                ProvisioningTenantState = "Revoked"
)

func PossibleValuesForProvisioningTenantState() []string {
	return []string{
		string(ProvisioningTenantStateActive),
		string(ProvisioningTenantStateBillingRequestDeclined),
		string(ProvisioningTenantStateBillingRequestExpired),
		string(ProvisioningTenantStateNotRequested),
		string(ProvisioningTenantStateOther),
		string(ProvisioningTenantStatePending),
		string(ProvisioningTenantStateRevoked),
	}
}

func (s *ProvisioningTenantState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningTenantState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningTenantState(input string) (*ProvisioningTenantState, error) {
	vals := map[string]ProvisioningTenantState{
		"active":                 ProvisioningTenantStateActive,
		"billingrequestdeclined": ProvisioningTenantStateBillingRequestDeclined,
		"billingrequestexpired":  ProvisioningTenantStateBillingRequestExpired,
		"notrequested":           ProvisioningTenantStateNotRequested,
		"other":                  ProvisioningTenantStateOther,
		"pending":                ProvisioningTenantStatePending,
		"revoked":                ProvisioningTenantStateRevoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningTenantState(input)
	return &out, nil
}
