package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlFrequencyInterval string

const (
	SignInFrequencySessionControlFrequencyIntervaleveryTime SignInFrequencySessionControlFrequencyInterval = "EveryTime"
	SignInFrequencySessionControlFrequencyIntervaltimeBased SignInFrequencySessionControlFrequencyInterval = "TimeBased"
)

func PossibleValuesForSignInFrequencySessionControlFrequencyInterval() []string {
	return []string{
		string(SignInFrequencySessionControlFrequencyIntervaleveryTime),
		string(SignInFrequencySessionControlFrequencyIntervaltimeBased),
	}
}

func parseSignInFrequencySessionControlFrequencyInterval(input string) (*SignInFrequencySessionControlFrequencyInterval, error) {
	vals := map[string]SignInFrequencySessionControlFrequencyInterval{
		"everytime": SignInFrequencySessionControlFrequencyIntervaleveryTime,
		"timebased": SignInFrequencySessionControlFrequencyIntervaltimeBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlFrequencyInterval(input)
	return &out, nil
}
