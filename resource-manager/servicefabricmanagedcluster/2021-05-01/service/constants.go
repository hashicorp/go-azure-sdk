package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = "AlignedAffinity"
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

func PossibleValuesForServiceCorrelationScheme() []string {
	return []string{
		string(ServiceCorrelationSchemeAlignedAffinity),
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
		"alignedaffinity":    ServiceCorrelationSchemeAlignedAffinity,
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
	ServiceKindStateful  ServiceKind = "Stateful"
	ServiceKindStateless ServiceKind = "Stateless"
)

func PossibleValuesForServiceKind() []string {
	return []string{
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

func (s *ServicePackageActivationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePackageActivationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePackageActivationMode(input string) (*ServicePackageActivationMode, error) {
	vals := map[string]ServicePackageActivationMode{
		"exclusiveprocess": ServicePackageActivationModeExclusiveProcess,
		"sharedprocess":    ServicePackageActivationModeSharedProcess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePackageActivationMode(input)
	return &out, nil
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

func (s *ServiceScalingMechanismKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceScalingMechanismKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceScalingMechanismKind(input string) (*ServiceScalingMechanismKind, error) {
	vals := map[string]ServiceScalingMechanismKind{
		"addremoveincrementalnamedpartition": ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition,
		"scalepartitioninstancecount":        ServiceScalingMechanismKindScalePartitionInstanceCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceScalingMechanismKind(input)
	return &out, nil
}

type ServiceScalingTriggerKind string

const (
	ServiceScalingTriggerKindAveragePartitionLoad ServiceScalingTriggerKind = "AveragePartitionLoad"
	ServiceScalingTriggerKindAverageServiceLoad   ServiceScalingTriggerKind = "AverageServiceLoad"
)

func PossibleValuesForServiceScalingTriggerKind() []string {
	return []string{
		string(ServiceScalingTriggerKindAveragePartitionLoad),
		string(ServiceScalingTriggerKindAverageServiceLoad),
	}
}

func (s *ServiceScalingTriggerKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceScalingTriggerKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceScalingTriggerKind(input string) (*ServiceScalingTriggerKind, error) {
	vals := map[string]ServiceScalingTriggerKind{
		"averagepartitionload": ServiceScalingTriggerKindAveragePartitionLoad,
		"averageserviceload":   ServiceScalingTriggerKindAverageServiceLoad,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceScalingTriggerKind(input)
	return &out, nil
}
