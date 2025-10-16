package blobservice

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedMethods string

const (
	AllowedMethodsCONNECT AllowedMethods = "CONNECT"
	AllowedMethodsDELETE  AllowedMethods = "DELETE"
	AllowedMethodsGET     AllowedMethods = "GET"
	AllowedMethodsHEAD    AllowedMethods = "HEAD"
	AllowedMethodsMERGE   AllowedMethods = "MERGE"
	AllowedMethodsOPTIONS AllowedMethods = "OPTIONS"
	AllowedMethodsPATCH   AllowedMethods = "PATCH"
	AllowedMethodsPOST    AllowedMethods = "POST"
	AllowedMethodsPUT     AllowedMethods = "PUT"
	AllowedMethodsTRACE   AllowedMethods = "TRACE"
)

func PossibleValuesForAllowedMethods() []string {
	return []string{
		string(AllowedMethodsCONNECT),
		string(AllowedMethodsDELETE),
		string(AllowedMethodsGET),
		string(AllowedMethodsHEAD),
		string(AllowedMethodsMERGE),
		string(AllowedMethodsOPTIONS),
		string(AllowedMethodsPATCH),
		string(AllowedMethodsPOST),
		string(AllowedMethodsPUT),
		string(AllowedMethodsTRACE),
	}
}

func (s *AllowedMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedMethods(input string) (*AllowedMethods, error) {
	vals := map[string]AllowedMethods{
		"connect": AllowedMethodsCONNECT,
		"delete":  AllowedMethodsDELETE,
		"get":     AllowedMethodsGET,
		"head":    AllowedMethodsHEAD,
		"merge":   AllowedMethodsMERGE,
		"options": AllowedMethodsOPTIONS,
		"patch":   AllowedMethodsPATCH,
		"post":    AllowedMethodsPOST,
		"put":     AllowedMethodsPUT,
		"trace":   AllowedMethodsTRACE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedMethods(input)
	return &out, nil
}

type Name string

const (
	NameAccessTimeTracking Name = "AccessTimeTracking"
)

func PossibleValuesForName() []string {
	return []string{
		string(NameAccessTimeTracking),
	}
}

func (s *Name) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseName(input string) (*Name, error) {
	vals := map[string]Name{
		"accesstimetracking": NameAccessTimeTracking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Name(input)
	return &out, nil
}

type SkuName string

const (
	SkuNamePremiumLRS       SkuName = "Premium_LRS"
	SkuNamePremiumVTwoLRS   SkuName = "PremiumV2_LRS"
	SkuNamePremiumVTwoZRS   SkuName = "PremiumV2_ZRS"
	SkuNamePremiumZRS       SkuName = "Premium_ZRS"
	SkuNameStandardGRS      SkuName = "Standard_GRS"
	SkuNameStandardGZRS     SkuName = "Standard_GZRS"
	SkuNameStandardLRS      SkuName = "Standard_LRS"
	SkuNameStandardRAGRS    SkuName = "Standard_RAGRS"
	SkuNameStandardRAGZRS   SkuName = "Standard_RAGZRS"
	SkuNameStandardVTwoGRS  SkuName = "StandardV2_GRS"
	SkuNameStandardVTwoGZRS SkuName = "StandardV2_GZRS"
	SkuNameStandardVTwoLRS  SkuName = "StandardV2_LRS"
	SkuNameStandardVTwoZRS  SkuName = "StandardV2_ZRS"
	SkuNameStandardZRS      SkuName = "Standard_ZRS"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumLRS),
		string(SkuNamePremiumVTwoLRS),
		string(SkuNamePremiumVTwoZRS),
		string(SkuNamePremiumZRS),
		string(SkuNameStandardGRS),
		string(SkuNameStandardGZRS),
		string(SkuNameStandardLRS),
		string(SkuNameStandardRAGRS),
		string(SkuNameStandardRAGZRS),
		string(SkuNameStandardVTwoGRS),
		string(SkuNameStandardVTwoGZRS),
		string(SkuNameStandardVTwoLRS),
		string(SkuNameStandardVTwoZRS),
		string(SkuNameStandardZRS),
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
		"premium_lrs":     SkuNamePremiumLRS,
		"premiumv2_lrs":   SkuNamePremiumVTwoLRS,
		"premiumv2_zrs":   SkuNamePremiumVTwoZRS,
		"premium_zrs":     SkuNamePremiumZRS,
		"standard_grs":    SkuNameStandardGRS,
		"standard_gzrs":   SkuNameStandardGZRS,
		"standard_lrs":    SkuNameStandardLRS,
		"standard_ragrs":  SkuNameStandardRAGRS,
		"standard_ragzrs": SkuNameStandardRAGZRS,
		"standardv2_grs":  SkuNameStandardVTwoGRS,
		"standardv2_gzrs": SkuNameStandardVTwoGZRS,
		"standardv2_lrs":  SkuNameStandardVTwoLRS,
		"standardv2_zrs":  SkuNameStandardVTwoZRS,
		"standard_zrs":    SkuNameStandardZRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
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
