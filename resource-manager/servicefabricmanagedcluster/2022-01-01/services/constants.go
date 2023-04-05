package services

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
	PartitionSchemeNamed                  PartitionScheme = "Named"
	PartitionSchemeSingleton              PartitionScheme = "Singleton"
	PartitionSchemeUniformIntSixFourRange PartitionScheme = "UniformInt64Range"
)

func PossibleValuesForPartitionScheme() []string {
	return []string{
		string(PartitionSchemeNamed),
		string(PartitionSchemeSingleton),
		string(PartitionSchemeUniformIntSixFourRange),
	}
}

type ServiceCorrelationScheme string

const (
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = "AlignedAffinity"
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

func PossibleValuesForServiceCorrelationScheme() []string {
	return []string{
		string(ServiceCorrelationSchemeAlignedAffinity),
		string(ServiceCorrelationSchemeNonAlignedAffinity),
	}
}

type ServiceKind string

const (
	ServiceKindStateful  ServiceKind = "Stateful"
	ServiceKindStateless ServiceKind = "Stateless"
)

func PossibleValuesForServiceKind() []string {
	return []string{
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

type ServicePackageActivationMode string

const (
	ServicePackageActivationModeExclusiveProcess ServicePackageActivationMode = "ExclusiveProcess"
	ServicePackageActivationModeSharedProcess    ServicePackageActivationMode = "SharedProcess"
)

func PossibleValuesForServicePackageActivationMode() []string {
	return []string{
		string(ServicePackageActivationModeExclusiveProcess),
		string(ServicePackageActivationModeSharedProcess),
	}
}

type ServicePlacementPolicyType string

const (
	ServicePlacementPolicyTypeInvalidDomain              ServicePlacementPolicyType = "InvalidDomain"
	ServicePlacementPolicyTypeNonPartiallyPlaceService   ServicePlacementPolicyType = "NonPartiallyPlaceService"
	ServicePlacementPolicyTypePreferredPrimaryDomain     ServicePlacementPolicyType = "PreferredPrimaryDomain"
	ServicePlacementPolicyTypeRequiredDomain             ServicePlacementPolicyType = "RequiredDomain"
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

func PossibleValuesForServicePlacementPolicyType() []string {
	return []string{
		string(ServicePlacementPolicyTypeInvalidDomain),
		string(ServicePlacementPolicyTypeNonPartiallyPlaceService),
		string(ServicePlacementPolicyTypePreferredPrimaryDomain),
		string(ServicePlacementPolicyTypeRequiredDomain),
		string(ServicePlacementPolicyTypeRequiredDomainDistribution),
	}
}

type ServiceScalingMechanismKind string

const (
	ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition ServiceScalingMechanismKind = "AddRemoveIncrementalNamedPartition"
	ServiceScalingMechanismKindScalePartitionInstanceCount        ServiceScalingMechanismKind = "ScalePartitionInstanceCount"
)

func PossibleValuesForServiceScalingMechanismKind() []string {
	return []string{
		string(ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition),
		string(ServiceScalingMechanismKindScalePartitionInstanceCount),
	}
}

type ServiceScalingTriggerKind string

const (
	ServiceScalingTriggerKindAveragePartitionLoadTrigger ServiceScalingTriggerKind = "AveragePartitionLoadTrigger"
	ServiceScalingTriggerKindAverageServiceLoadTrigger   ServiceScalingTriggerKind = "AverageServiceLoadTrigger"
)

func PossibleValuesForServiceScalingTriggerKind() []string {
	return []string{
		string(ServiceScalingTriggerKindAveragePartitionLoadTrigger),
		string(ServiceScalingTriggerKindAverageServiceLoadTrigger),
	}
}
