package webapplicationfirewallmanagedrulesets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuName string

const (
	SkuNameCustomVerizon                             SkuName = "Custom_Verizon"
	SkuNamePremiumAzureFrontDoor                     SkuName = "Premium_AzureFrontDoor"
	SkuNamePremiumVerizon                            SkuName = "Premium_Verizon"
	SkuNameStandardAkamai                            SkuName = "Standard_Akamai"
	SkuNameStandardAvgBandWidthChinaCdn              SkuName = "Standard_AvgBandWidth_ChinaCdn"
	SkuNameStandardAzureFrontDoor                    SkuName = "Standard_AzureFrontDoor"
	SkuNameStandardChinaCdn                          SkuName = "Standard_ChinaCdn"
	SkuNameStandardMicrosoft                         SkuName = "Standard_Microsoft"
	SkuNameStandardNineFiveFiveBandWidthChinaCdn     SkuName = "Standard_955BandWidth_ChinaCdn"
	SkuNameStandardPlusAvgBandWidthChinaCdn          SkuName = "StandardPlus_AvgBandWidth_ChinaCdn"
	SkuNameStandardPlusChinaCdn                      SkuName = "StandardPlus_ChinaCdn"
	SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn SkuName = "StandardPlus_955BandWidth_ChinaCdn"
	SkuNameStandardVerizon                           SkuName = "Standard_Verizon"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameCustomVerizon),
		string(SkuNamePremiumAzureFrontDoor),
		string(SkuNamePremiumVerizon),
		string(SkuNameStandardAkamai),
		string(SkuNameStandardAvgBandWidthChinaCdn),
		string(SkuNameStandardAzureFrontDoor),
		string(SkuNameStandardChinaCdn),
		string(SkuNameStandardMicrosoft),
		string(SkuNameStandardNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardPlusAvgBandWidthChinaCdn),
		string(SkuNameStandardPlusChinaCdn),
		string(SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn),
		string(SkuNameStandardVerizon),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"custom_verizon":                     SkuNameCustomVerizon,
		"premium_azurefrontdoor":             SkuNamePremiumAzureFrontDoor,
		"premium_verizon":                    SkuNamePremiumVerizon,
		"standard_akamai":                    SkuNameStandardAkamai,
		"standard_avgbandwidth_chinacdn":     SkuNameStandardAvgBandWidthChinaCdn,
		"standard_azurefrontdoor":            SkuNameStandardAzureFrontDoor,
		"standard_chinacdn":                  SkuNameStandardChinaCdn,
		"standard_microsoft":                 SkuNameStandardMicrosoft,
		"standard_955bandwidth_chinacdn":     SkuNameStandardNineFiveFiveBandWidthChinaCdn,
		"standardplus_avgbandwidth_chinacdn": SkuNameStandardPlusAvgBandWidthChinaCdn,
		"standardplus_chinacdn":              SkuNameStandardPlusChinaCdn,
		"standardplus_955bandwidth_chinacdn": SkuNameStandardPlusNineFiveFiveBandWidthChinaCdn,
		"standard_verizon":                   SkuNameStandardVerizon,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}
