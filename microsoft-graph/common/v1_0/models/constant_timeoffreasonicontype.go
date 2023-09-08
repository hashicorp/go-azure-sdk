package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffReasonIconType string

const (
	TimeOffReasonIconTypecake        TimeOffReasonIconType = "Cake"
	TimeOffReasonIconTypecalendar    TimeOffReasonIconType = "Calendar"
	TimeOffReasonIconTypecar         TimeOffReasonIconType = "Car"
	TimeOffReasonIconTypeclock       TimeOffReasonIconType = "Clock"
	TimeOffReasonIconTypecup         TimeOffReasonIconType = "Cup"
	TimeOffReasonIconTypedoctor      TimeOffReasonIconType = "Doctor"
	TimeOffReasonIconTypedog         TimeOffReasonIconType = "Dog"
	TimeOffReasonIconTypefirstAid    TimeOffReasonIconType = "FirstAid"
	TimeOffReasonIconTypeglobe       TimeOffReasonIconType = "Globe"
	TimeOffReasonIconTypejuryDuty    TimeOffReasonIconType = "JuryDuty"
	TimeOffReasonIconTypenone        TimeOffReasonIconType = "None"
	TimeOffReasonIconTypenotWorking  TimeOffReasonIconType = "NotWorking"
	TimeOffReasonIconTypephone       TimeOffReasonIconType = "Phone"
	TimeOffReasonIconTypepiggyBank   TimeOffReasonIconType = "PiggyBank"
	TimeOffReasonIconTypepin         TimeOffReasonIconType = "Pin"
	TimeOffReasonIconTypeplane       TimeOffReasonIconType = "Plane"
	TimeOffReasonIconTyperunning     TimeOffReasonIconType = "Running"
	TimeOffReasonIconTypesunny       TimeOffReasonIconType = "Sunny"
	TimeOffReasonIconTypetrafficCone TimeOffReasonIconType = "TrafficCone"
	TimeOffReasonIconTypeumbrella    TimeOffReasonIconType = "Umbrella"
	TimeOffReasonIconTypeweather     TimeOffReasonIconType = "Weather"
)

func PossibleValuesForTimeOffReasonIconType() []string {
	return []string{
		string(TimeOffReasonIconTypecake),
		string(TimeOffReasonIconTypecalendar),
		string(TimeOffReasonIconTypecar),
		string(TimeOffReasonIconTypeclock),
		string(TimeOffReasonIconTypecup),
		string(TimeOffReasonIconTypedoctor),
		string(TimeOffReasonIconTypedog),
		string(TimeOffReasonIconTypefirstAid),
		string(TimeOffReasonIconTypeglobe),
		string(TimeOffReasonIconTypejuryDuty),
		string(TimeOffReasonIconTypenone),
		string(TimeOffReasonIconTypenotWorking),
		string(TimeOffReasonIconTypephone),
		string(TimeOffReasonIconTypepiggyBank),
		string(TimeOffReasonIconTypepin),
		string(TimeOffReasonIconTypeplane),
		string(TimeOffReasonIconTyperunning),
		string(TimeOffReasonIconTypesunny),
		string(TimeOffReasonIconTypetrafficCone),
		string(TimeOffReasonIconTypeumbrella),
		string(TimeOffReasonIconTypeweather),
	}
}

func parseTimeOffReasonIconType(input string) (*TimeOffReasonIconType, error) {
	vals := map[string]TimeOffReasonIconType{
		"cake":        TimeOffReasonIconTypecake,
		"calendar":    TimeOffReasonIconTypecalendar,
		"car":         TimeOffReasonIconTypecar,
		"clock":       TimeOffReasonIconTypeclock,
		"cup":         TimeOffReasonIconTypecup,
		"doctor":      TimeOffReasonIconTypedoctor,
		"dog":         TimeOffReasonIconTypedog,
		"firstaid":    TimeOffReasonIconTypefirstAid,
		"globe":       TimeOffReasonIconTypeglobe,
		"juryduty":    TimeOffReasonIconTypejuryDuty,
		"none":        TimeOffReasonIconTypenone,
		"notworking":  TimeOffReasonIconTypenotWorking,
		"phone":       TimeOffReasonIconTypephone,
		"piggybank":   TimeOffReasonIconTypepiggyBank,
		"pin":         TimeOffReasonIconTypepin,
		"plane":       TimeOffReasonIconTypeplane,
		"running":     TimeOffReasonIconTyperunning,
		"sunny":       TimeOffReasonIconTypesunny,
		"trafficcone": TimeOffReasonIconTypetrafficCone,
		"umbrella":    TimeOffReasonIconTypeumbrella,
		"weather":     TimeOffReasonIconTypeweather,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffReasonIconType(input)
	return &out, nil
}
