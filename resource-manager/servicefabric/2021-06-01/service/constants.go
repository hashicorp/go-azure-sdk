package service

import "strings"

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
