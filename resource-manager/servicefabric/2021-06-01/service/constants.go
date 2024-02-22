package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmServicePackageActivationMode string

const (
	ArmServicePackageActivationModeExclusiveProcess ArmServicePackageActivationMode = "ExclusiveProcess"
	ArmServicePackageActivationModeSharedProcess    ArmServicePackageActivationMode = "SharedProcess"
)

func PossibleValuesForArmServicePackageActivationMode() []string {
	return []string{
		string(ArmServicePackageActivationModeExclusiveProcess),
		string(ArmServicePackageActivationModeSharedProcess),
	}
}

func (s *ArmServicePackageActivationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseArmServicePackageActivationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseArmServicePackageActivationMode(input string) (*ArmServicePackageActivationMode, error) {
	vals := map[string]ArmServicePackageActivationMode{
		"exclusiveprocess": ArmServicePackageActivationModeExclusiveProcess,
		"sharedprocess":    ArmServicePackageActivationModeSharedProcess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArmServicePackageActivationMode(input)
	return &out, nil
}

type MoveCost string

const (
	MoveCostHigh   MoveCost = "High"
	MoveCostLow    MoveCost = "Low"
	MoveCostMedium MoveCost = "Medium"
	MoveCostZero   MoveCost = "Zero"
)

func PossibleValuesForMoveCost() []string {
	return []string{
		string(MoveCostHigh),
		string(MoveCostLow),
		string(MoveCostMedium),
		string(MoveCostZero),
	}
}

func (s *MoveCost) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMoveCost(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMoveCost(input string) (*MoveCost, error) {
	vals := map[string]MoveCost{
		"high":   MoveCostHigh,
		"low":    MoveCostLow,
		"medium": MoveCostMedium,
		"zero":   MoveCostZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MoveCost(input)
	return &out, nil
}

type PartitionScheme string

const (
	PartitionSchemeInvalid                PartitionScheme = "Invalid"
	PartitionSchemeNamed                  PartitionScheme = "Named"
	PartitionSchemeSingleton              PartitionScheme = "Singleton"
	PartitionSchemeUniformIntSixFourRange PartitionScheme = "UniformInt64Range"
)

func PossibleValuesForPartitionScheme() []string {
	return []string{
		string(PartitionSchemeInvalid),
		string(PartitionSchemeNamed),
		string(PartitionSchemeSingleton),
		string(PartitionSchemeUniformIntSixFourRange),
	}
}

func (s *PartitionScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartitionScheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartitionScheme(input string) (*PartitionScheme, error) {
	vals := map[string]PartitionScheme{
		"invalid":           PartitionSchemeInvalid,
		"named":             PartitionSchemeNamed,
		"singleton":         PartitionSchemeSingleton,
		"uniformint64range": PartitionSchemeUniformIntSixFourRange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartitionScheme(input)
	return &out, nil
}

type ServiceCorrelationScheme string

const (
	ServiceCorrelationSchemeAffinity           ServiceCorrelationScheme = "Affinity"
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = "AlignedAffinity"
	ServiceCorrelationSchemeInvalid            ServiceCorrelationScheme = "Invalid"
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

func PossibleValuesForServiceCorrelationScheme() []string {
	return []string{
		string(ServiceCorrelationSchemeAffinity),
		string(ServiceCorrelationSchemeAlignedAffinity),
		string(ServiceCorrelationSchemeInvalid),
		string(ServiceCorrelationSchemeNonAlignedAffinity),
	}
}

func (s *ServiceCorrelationScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceCorrelationScheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceCorrelationScheme(input string) (*ServiceCorrelationScheme, error) {
	vals := map[string]ServiceCorrelationScheme{
		"affinity":           ServiceCorrelationSchemeAffinity,
		"alignedaffinity":    ServiceCorrelationSchemeAlignedAffinity,
		"invalid":            ServiceCorrelationSchemeInvalid,
		"nonalignedaffinity": ServiceCorrelationSchemeNonAlignedAffinity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceCorrelationScheme(input)
	return &out, nil
}

type ServiceKind string

const (
	ServiceKindInvalid   ServiceKind = "Invalid"
	ServiceKindStateful  ServiceKind = "Stateful"
	ServiceKindStateless ServiceKind = "Stateless"
)

func PossibleValuesForServiceKind() []string {
	return []string{
		string(ServiceKindInvalid),
		string(ServiceKindStateful),
		string(ServiceKindStateless),
	}
}

func (s *ServiceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceKind(input string) (*ServiceKind, error) {
	vals := map[string]ServiceKind{
		"invalid":   ServiceKindInvalid,
		"stateful":  ServiceKindStateful,
		"stateless": ServiceKindStateless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceKind(input)
	return &out, nil
}

type ServiceLoadMetricWeight string

const (
	ServiceLoadMetricWeightHigh   ServiceLoadMetricWeight = "High"
	ServiceLoadMetricWeightLow    ServiceLoadMetricWeight = "Low"
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = "Medium"
	ServiceLoadMetricWeightZero   ServiceLoadMetricWeight = "Zero"
)

func PossibleValuesForServiceLoadMetricWeight() []string {
	return []string{
		string(ServiceLoadMetricWeightHigh),
		string(ServiceLoadMetricWeightLow),
		string(ServiceLoadMetricWeightMedium),
		string(ServiceLoadMetricWeightZero),
	}
}

func (s *ServiceLoadMetricWeight) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceLoadMetricWeight(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceLoadMetricWeight(input string) (*ServiceLoadMetricWeight, error) {
	vals := map[string]ServiceLoadMetricWeight{
		"high":   ServiceLoadMetricWeightHigh,
		"low":    ServiceLoadMetricWeightLow,
		"medium": ServiceLoadMetricWeightMedium,
		"zero":   ServiceLoadMetricWeightZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceLoadMetricWeight(input)
	return &out, nil
}

type ServicePlacementPolicyType string

const (
	ServicePlacementPolicyTypeInvalid                    ServicePlacementPolicyType = "Invalid"
	ServicePlacementPolicyTypeInvalidDomain              ServicePlacementPolicyType = "InvalidDomain"
	ServicePlacementPolicyTypeNonPartiallyPlaceService   ServicePlacementPolicyType = "NonPartiallyPlaceService"
	ServicePlacementPolicyTypePreferredPrimaryDomain     ServicePlacementPolicyType = "PreferredPrimaryDomain"
	ServicePlacementPolicyTypeRequiredDomain             ServicePlacementPolicyType = "RequiredDomain"
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

func PossibleValuesForServicePlacementPolicyType() []string {
	return []string{
		string(ServicePlacementPolicyTypeInvalid),
		string(ServicePlacementPolicyTypeInvalidDomain),
		string(ServicePlacementPolicyTypeNonPartiallyPlaceService),
		string(ServicePlacementPolicyTypePreferredPrimaryDomain),
		string(ServicePlacementPolicyTypeRequiredDomain),
		string(ServicePlacementPolicyTypeRequiredDomainDistribution),
	}
}

func (s *ServicePlacementPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePlacementPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePlacementPolicyType(input string) (*ServicePlacementPolicyType, error) {
	vals := map[string]ServicePlacementPolicyType{
		"invalid":                    ServicePlacementPolicyTypeInvalid,
		"invaliddomain":              ServicePlacementPolicyTypeInvalidDomain,
		"nonpartiallyplaceservice":   ServicePlacementPolicyTypeNonPartiallyPlaceService,
		"preferredprimarydomain":     ServicePlacementPolicyTypePreferredPrimaryDomain,
		"requireddomain":             ServicePlacementPolicyTypeRequiredDomain,
		"requireddomaindistribution": ServicePlacementPolicyTypeRequiredDomainDistribution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePlacementPolicyType(input)
	return &out, nil
}
