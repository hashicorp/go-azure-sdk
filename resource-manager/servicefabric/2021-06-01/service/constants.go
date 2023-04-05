package service

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
