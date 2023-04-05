package scalingplanpooledschedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
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
