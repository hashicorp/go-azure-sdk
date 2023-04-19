package availableskus

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShipmentType string

const (
	ShipmentTypeNotApplicable     ShipmentType = "NotApplicable"
	ShipmentTypeSelfPickup        ShipmentType = "SelfPickup"
	ShipmentTypeShippedToCustomer ShipmentType = "ShippedToCustomer"
)

func PossibleValuesForShipmentType() []string {
	return []string{
		string(ShipmentTypeNotApplicable),
		string(ShipmentTypeSelfPickup),
		string(ShipmentTypeShippedToCustomer),
	}
}

func (s *ShipmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShipmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShipmentType(input string) (*ShipmentType, error) {
	vals := map[string]ShipmentType{
		"notapplicable":     ShipmentTypeNotApplicable,
		"selfpickup":        ShipmentTypeSelfPickup,
		"shippedtocustomer": ShipmentTypeShippedToCustomer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShipmentType(input)
	return &out, nil
}

type SkuAvailability string

const (
	SkuAvailabilityAvailable   SkuAvailability = "Available"
	SkuAvailabilityUnavailable SkuAvailability = "Unavailable"
)

func PossibleValuesForSkuAvailability() []string {
	return []string{
		string(SkuAvailabilityAvailable),
		string(SkuAvailabilityUnavailable),
	}
}

func (s *SkuAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuAvailability(input string) (*SkuAvailability, error) {
	vals := map[string]SkuAvailability{
		"available":   SkuAvailabilityAvailable,
		"unavailable": SkuAvailabilityUnavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuAvailability(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameEdge                 SkuName = "Edge"
	SkuNameEdgeMRMini           SkuName = "EdgeMR_Mini"
	SkuNameEdgePBase            SkuName = "EdgeP_Base"
	SkuNameEdgePHigh            SkuName = "EdgeP_High"
	SkuNameEdgePRBase           SkuName = "EdgePR_Base"
	SkuNameEdgePRBaseUPS        SkuName = "EdgePR_Base_UPS"
	SkuNameGPU                  SkuName = "GPU"
	SkuNameGateway              SkuName = "Gateway"
	SkuNameRCALarge             SkuName = "RCA_Large"
	SkuNameRCASmall             SkuName = "RCA_Small"
	SkuNameRDC                  SkuName = "RDC"
	SkuNameTCALarge             SkuName = "TCA_Large"
	SkuNameTCASmall             SkuName = "TCA_Small"
	SkuNameTDC                  SkuName = "TDC"
	SkuNameTEAFourNodeHeater    SkuName = "TEA_4Node_Heater"
	SkuNameTEAFourNodeUPSHeater SkuName = "TEA_4Node_UPS_Heater"
	SkuNameTEAOneNode           SkuName = "TEA_1Node"
	SkuNameTEAOneNodeHeater     SkuName = "TEA_1Node_Heater"
	SkuNameTEAOneNodeUPS        SkuName = "TEA_1Node_UPS"
	SkuNameTEAOneNodeUPSHeater  SkuName = "TEA_1Node_UPS_Heater"
	SkuNameTMA                  SkuName = "TMA"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEdge),
		string(SkuNameEdgeMRMini),
		string(SkuNameEdgePBase),
		string(SkuNameEdgePHigh),
		string(SkuNameEdgePRBase),
		string(SkuNameEdgePRBaseUPS),
		string(SkuNameGPU),
		string(SkuNameGateway),
		string(SkuNameRCALarge),
		string(SkuNameRCASmall),
		string(SkuNameRDC),
		string(SkuNameTCALarge),
		string(SkuNameTCASmall),
		string(SkuNameTDC),
		string(SkuNameTEAFourNodeHeater),
		string(SkuNameTEAFourNodeUPSHeater),
		string(SkuNameTEAOneNode),
		string(SkuNameTEAOneNodeHeater),
		string(SkuNameTEAOneNodeUPS),
		string(SkuNameTEAOneNodeUPSHeater),
		string(SkuNameTMA),
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
		"edge":                 SkuNameEdge,
		"edgemr_mini":          SkuNameEdgeMRMini,
		"edgep_base":           SkuNameEdgePBase,
		"edgep_high":           SkuNameEdgePHigh,
		"edgepr_base":          SkuNameEdgePRBase,
		"edgepr_base_ups":      SkuNameEdgePRBaseUPS,
		"gpu":                  SkuNameGPU,
		"gateway":              SkuNameGateway,
		"rca_large":            SkuNameRCALarge,
		"rca_small":            SkuNameRCASmall,
		"rdc":                  SkuNameRDC,
		"tca_large":            SkuNameTCALarge,
		"tca_small":            SkuNameTCASmall,
		"tdc":                  SkuNameTDC,
		"tea_4node_heater":     SkuNameTEAFourNodeHeater,
		"tea_4node_ups_heater": SkuNameTEAFourNodeUPSHeater,
		"tea_1node":            SkuNameTEAOneNode,
		"tea_1node_heater":     SkuNameTEAOneNodeHeater,
		"tea_1node_ups":        SkuNameTEAOneNodeUPS,
		"tea_1node_ups_heater": SkuNameTEAOneNodeUPSHeater,
		"tma":                  SkuNameTMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type SkuSignupOption string

const (
	SkuSignupOptionAvailable SkuSignupOption = "Available"
	SkuSignupOptionNone      SkuSignupOption = "None"
)

func PossibleValuesForSkuSignupOption() []string {
	return []string{
		string(SkuSignupOptionAvailable),
		string(SkuSignupOptionNone),
	}
}

func (s *SkuSignupOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuSignupOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuSignupOption(input string) (*SkuSignupOption, error) {
	vals := map[string]SkuSignupOption{
		"available": SkuSignupOptionAvailable,
		"none":      SkuSignupOptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuSignupOption(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
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
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type SkuVersion string

const (
	SkuVersionPreview SkuVersion = "Preview"
	SkuVersionStable  SkuVersion = "Stable"
)

func PossibleValuesForSkuVersion() []string {
	return []string{
		string(SkuVersionPreview),
		string(SkuVersionStable),
	}
}

func (s *SkuVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuVersion(input string) (*SkuVersion, error) {
	vals := map[string]SkuVersion{
		"preview": SkuVersionPreview,
		"stable":  SkuVersionStable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuVersion(input)
	return &out, nil
}
