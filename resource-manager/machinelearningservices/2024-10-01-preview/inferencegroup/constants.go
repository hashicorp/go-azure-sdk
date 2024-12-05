package inferencegroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrderString string

const (
	OrderStringCreatedAtAsc  OrderString = "CreatedAtAsc"
	OrderStringCreatedAtDesc OrderString = "CreatedAtDesc"
	OrderStringUpdatedAtAsc  OrderString = "UpdatedAtAsc"
	OrderStringUpdatedAtDesc OrderString = "UpdatedAtDesc"
)

func PossibleValuesForOrderString() []string {
	return []string{
		string(OrderStringCreatedAtAsc),
		string(OrderStringCreatedAtDesc),
		string(OrderStringUpdatedAtAsc),
		string(OrderStringUpdatedAtDesc),
	}
}

func (s *OrderString) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrderString(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrderString(input string) (*OrderString, error) {
	vals := map[string]OrderString{
		"createdatasc":  OrderStringCreatedAtAsc,
		"createdatdesc": OrderStringCreatedAtDesc,
		"updatedatasc":  OrderStringUpdatedAtAsc,
		"updatedatdesc": OrderStringUpdatedAtDesc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrderString(input)
	return &out, nil
}

type PoolProvisioningState string

const (
	PoolProvisioningStateCanceled  PoolProvisioningState = "Canceled"
	PoolProvisioningStateCreating  PoolProvisioningState = "Creating"
	PoolProvisioningStateDeleting  PoolProvisioningState = "Deleting"
	PoolProvisioningStateFailed    PoolProvisioningState = "Failed"
	PoolProvisioningStateSucceeded PoolProvisioningState = "Succeeded"
	PoolProvisioningStateUpdating  PoolProvisioningState = "Updating"
)

func PossibleValuesForPoolProvisioningState() []string {
	return []string{
		string(PoolProvisioningStateCanceled),
		string(PoolProvisioningStateCreating),
		string(PoolProvisioningStateDeleting),
		string(PoolProvisioningStateFailed),
		string(PoolProvisioningStateSucceeded),
		string(PoolProvisioningStateUpdating),
	}
}

func (s *PoolProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePoolProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePoolProvisioningState(input string) (*PoolProvisioningState, error) {
	vals := map[string]PoolProvisioningState{
		"canceled":  PoolProvisioningStateCanceled,
		"creating":  PoolProvisioningStateCreating,
		"deleting":  PoolProvisioningStateDeleting,
		"failed":    PoolProvisioningStateFailed,
		"succeeded": PoolProvisioningStateSucceeded,
		"updating":  PoolProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PoolProvisioningState(input)
	return &out, nil
}

type SkuScaleType string

const (
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	SkuScaleTypeManual    SkuScaleType = "Manual"
	SkuScaleTypeNone      SkuScaleType = "None"
)

func PossibleValuesForSkuScaleType() []string {
	return []string{
		string(SkuScaleTypeAutomatic),
		string(SkuScaleTypeManual),
		string(SkuScaleTypeNone),
	}
}

func (s *SkuScaleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuScaleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuScaleType(input string) (*SkuScaleType, error) {
	vals := map[string]SkuScaleType{
		"automatic": SkuScaleTypeAutomatic,
		"manual":    SkuScaleTypeManual,
		"none":      SkuScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuScaleType(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
