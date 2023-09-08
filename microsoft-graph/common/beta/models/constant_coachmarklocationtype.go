package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CoachmarkLocationType string

const (
	CoachmarkLocationTypedisplayName CoachmarkLocationType = "DisplayName"
	CoachmarkLocationTypeexternalTag CoachmarkLocationType = "ExternalTag"
	CoachmarkLocationTypefromEmail   CoachmarkLocationType = "FromEmail"
	CoachmarkLocationTypemessageBody CoachmarkLocationType = "MessageBody"
	CoachmarkLocationTypesubject     CoachmarkLocationType = "Subject"
	CoachmarkLocationTypeunknown     CoachmarkLocationType = "Unknown"
)

func PossibleValuesForCoachmarkLocationType() []string {
	return []string{
		string(CoachmarkLocationTypedisplayName),
		string(CoachmarkLocationTypeexternalTag),
		string(CoachmarkLocationTypefromEmail),
		string(CoachmarkLocationTypemessageBody),
		string(CoachmarkLocationTypesubject),
		string(CoachmarkLocationTypeunknown),
	}
}

func parseCoachmarkLocationType(input string) (*CoachmarkLocationType, error) {
	vals := map[string]CoachmarkLocationType{
		"displayname": CoachmarkLocationTypedisplayName,
		"externaltag": CoachmarkLocationTypeexternalTag,
		"fromemail":   CoachmarkLocationTypefromEmail,
		"messagebody": CoachmarkLocationTypemessageBody,
		"subject":     CoachmarkLocationTypesubject,
		"unknown":     CoachmarkLocationTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CoachmarkLocationType(input)
	return &out, nil
}
