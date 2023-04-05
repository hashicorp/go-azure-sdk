package tenants

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingType string

const (
	BillingTypeAuths              BillingType = "auths"
	BillingTypeMonthlyActiveUsers BillingType = "mau"
)

func PossibleValuesForBillingType() []string {
	return []string{
		string(BillingTypeAuths),
		string(BillingTypeMonthlyActiveUsers),
	}
}

func (s *BillingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForBillingType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = BillingType(decoded)
	return nil
}

type Location string

const (
	LocationAsiaPacific  Location = "Asia Pacific"
	LocationAustralia    Location = "Australia"
	LocationEurope       Location = "Europe"
	LocationGlobal       Location = "Global"
	LocationUnitedStates Location = "United States"
)

func PossibleValuesForLocation() []string {
	return []string{
		string(LocationAsiaPacific),
		string(LocationAustralia),
		string(LocationEurope),
		string(LocationGlobal),
		string(LocationUnitedStates),
	}
}

func (s *Location) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLocation() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Location(decoded)
	return nil
}

type SkuName string

const (
	SkuNamePremiumP1 SkuName = "PremiumP1"
	SkuNamePremiumP2 SkuName = "PremiumP2"
	SkuNameStandard  SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumP1),
		string(SkuNamePremiumP2),
		string(SkuNameStandard),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSkuName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SkuName(decoded)
	return nil
}

type SkuTier string

const (
	SkuTierA0 SkuTier = "A0"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierA0),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSkuTier() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SkuTier(decoded)
	return nil
}
