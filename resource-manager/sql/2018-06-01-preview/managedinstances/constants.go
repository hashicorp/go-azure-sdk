package managedinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned,UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

func PossibleValuesForIdentityType() []string {
	return []string{
		string(IdentityTypeNone),
		string(IdentityTypeSystemAssigned),
		string(IdentityTypeSystemAssignedUserAssigned),
		string(IdentityTypeUserAssigned),
	}
}

func (s *IdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityType(input string) (*IdentityType, error) {
	vals := map[string]IdentityType{
		"none":                        IdentityTypeNone,
		"systemassigned":              IdentityTypeSystemAssigned,
		"systemassigned,userassigned": IdentityTypeSystemAssignedUserAssigned,
		"userassigned":                IdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityType(input)
	return &out, nil
}

type ManagedInstanceLicenseType string

const (
	ManagedInstanceLicenseTypeBasePrice       ManagedInstanceLicenseType = "BasePrice"
	ManagedInstanceLicenseTypeLicenseIncluded ManagedInstanceLicenseType = "LicenseIncluded"
)

func PossibleValuesForManagedInstanceLicenseType() []string {
	return []string{
		string(ManagedInstanceLicenseTypeBasePrice),
		string(ManagedInstanceLicenseTypeLicenseIncluded),
	}
}

func (s *ManagedInstanceLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedInstanceLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedInstanceLicenseType(input string) (*ManagedInstanceLicenseType, error) {
	vals := map[string]ManagedInstanceLicenseType{
		"baseprice":       ManagedInstanceLicenseTypeBasePrice,
		"licenseincluded": ManagedInstanceLicenseTypeLicenseIncluded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedInstanceLicenseType(input)
	return &out, nil
}

type ManagedInstanceProxyOverride string

const (
	ManagedInstanceProxyOverrideDefault  ManagedInstanceProxyOverride = "Default"
	ManagedInstanceProxyOverrideProxy    ManagedInstanceProxyOverride = "Proxy"
	ManagedInstanceProxyOverrideRedirect ManagedInstanceProxyOverride = "Redirect"
)

func PossibleValuesForManagedInstanceProxyOverride() []string {
	return []string{
		string(ManagedInstanceProxyOverrideDefault),
		string(ManagedInstanceProxyOverrideProxy),
		string(ManagedInstanceProxyOverrideRedirect),
	}
}

func (s *ManagedInstanceProxyOverride) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedInstanceProxyOverride(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedInstanceProxyOverride(input string) (*ManagedInstanceProxyOverride, error) {
	vals := map[string]ManagedInstanceProxyOverride{
		"default":  ManagedInstanceProxyOverrideDefault,
		"proxy":    ManagedInstanceProxyOverrideProxy,
		"redirect": ManagedInstanceProxyOverrideRedirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedInstanceProxyOverride(input)
	return &out, nil
}

type ManagedServerCreateMode string

const (
	ManagedServerCreateModeDefault            ManagedServerCreateMode = "Default"
	ManagedServerCreateModePointInTimeRestore ManagedServerCreateMode = "PointInTimeRestore"
)

func PossibleValuesForManagedServerCreateMode() []string {
	return []string{
		string(ManagedServerCreateModeDefault),
		string(ManagedServerCreateModePointInTimeRestore),
	}
}

func (s *ManagedServerCreateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedServerCreateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedServerCreateMode(input string) (*ManagedServerCreateMode, error) {
	vals := map[string]ManagedServerCreateMode{
		"default":            ManagedServerCreateModeDefault,
		"pointintimerestore": ManagedServerCreateModePointInTimeRestore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedServerCreateMode(input)
	return &out, nil
}
