package scalingplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

func PossibleValuesForDaysOfWeek() []string {
	return []string{
		string(DaysOfWeekFriday),
		string(DaysOfWeekMonday),
		string(DaysOfWeekSaturday),
		string(DaysOfWeekSunday),
		string(DaysOfWeekThursday),
		string(DaysOfWeekTuesday),
		string(DaysOfWeekWednesday),
	}
}

type ScalingHostPoolType string

const (
	ScalingHostPoolTypePooled ScalingHostPoolType = "Pooled"
)

func PossibleValuesForScalingHostPoolType() []string {
	return []string{
		string(ScalingHostPoolTypePooled),
	}
}

type SessionHostLoadBalancingAlgorithm string

const (
	SessionHostLoadBalancingAlgorithmBreadthFirst SessionHostLoadBalancingAlgorithm = "BreadthFirst"
	SessionHostLoadBalancingAlgorithmDepthFirst   SessionHostLoadBalancingAlgorithm = "DepthFirst"
)

func PossibleValuesForSessionHostLoadBalancingAlgorithm() []string {
	return []string{
		string(SessionHostLoadBalancingAlgorithmBreadthFirst),
		string(SessionHostLoadBalancingAlgorithmDepthFirst),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

type StopHostsWhen string

const (
	StopHostsWhenZeroActiveSessions StopHostsWhen = "ZeroActiveSessions"
	StopHostsWhenZeroSessions       StopHostsWhen = "ZeroSessions"
)

func PossibleValuesForStopHostsWhen() []string {
	return []string{
		string(StopHostsWhenZeroActiveSessions),
		string(StopHostsWhenZeroSessions),
	}
}
