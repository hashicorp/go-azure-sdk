package activesessionhostconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainJoinType string

const (
	DomainJoinTypeActiveDirectory      DomainJoinType = "ActiveDirectory"
	DomainJoinTypeAzureActiveDirectory DomainJoinType = "AzureActiveDirectory"
)

func PossibleValuesForDomainJoinType() []string {
	return []string{
		string(DomainJoinTypeActiveDirectory),
		string(DomainJoinTypeAzureActiveDirectory),
	}
}

func (s *DomainJoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDomainJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDomainJoinType(input string) (*DomainJoinType, error) {
	vals := map[string]DomainJoinType{
		"activedirectory":      DomainJoinTypeActiveDirectory,
		"azureactivedirectory": DomainJoinTypeAzureActiveDirectory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DomainJoinType(input)
	return &out, nil
}

type Type string

const (
	TypeCustom      Type = "Custom"
	TypeMarketplace Type = "Marketplace"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeCustom),
		string(TypeMarketplace),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"custom":      TypeCustom,
		"marketplace": TypeMarketplace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}

type VirtualMachineDiskType string

const (
	VirtualMachineDiskTypePremiumLRS     VirtualMachineDiskType = "Premium_LRS"
	VirtualMachineDiskTypeStandardLRS    VirtualMachineDiskType = "Standard_LRS"
	VirtualMachineDiskTypeStandardSSDLRS VirtualMachineDiskType = "StandardSSD_LRS"
)

func PossibleValuesForVirtualMachineDiskType() []string {
	return []string{
		string(VirtualMachineDiskTypePremiumLRS),
		string(VirtualMachineDiskTypeStandardLRS),
		string(VirtualMachineDiskTypeStandardSSDLRS),
	}
}

func (s *VirtualMachineDiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineDiskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineDiskType(input string) (*VirtualMachineDiskType, error) {
	vals := map[string]VirtualMachineDiskType{
		"premium_lrs":     VirtualMachineDiskTypePremiumLRS,
		"standard_lrs":    VirtualMachineDiskTypeStandardLRS,
		"standardssd_lrs": VirtualMachineDiskTypeStandardSSDLRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineDiskType(input)
	return &out, nil
}

type VirtualMachineSecurityType string

const (
	VirtualMachineSecurityTypeConfidentialVM VirtualMachineSecurityType = "ConfidentialVM"
	VirtualMachineSecurityTypeStandard       VirtualMachineSecurityType = "Standard"
	VirtualMachineSecurityTypeTrustedLaunch  VirtualMachineSecurityType = "TrustedLaunch"
)

func PossibleValuesForVirtualMachineSecurityType() []string {
	return []string{
		string(VirtualMachineSecurityTypeConfidentialVM),
		string(VirtualMachineSecurityTypeStandard),
		string(VirtualMachineSecurityTypeTrustedLaunch),
	}
}

func (s *VirtualMachineSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualMachineSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualMachineSecurityType(input string) (*VirtualMachineSecurityType, error) {
	vals := map[string]VirtualMachineSecurityType{
		"confidentialvm": VirtualMachineSecurityTypeConfidentialVM,
		"standard":       VirtualMachineSecurityTypeStandard,
		"trustedlaunch":  VirtualMachineSecurityTypeTrustedLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineSecurityType(input)
	return &out, nil
}
