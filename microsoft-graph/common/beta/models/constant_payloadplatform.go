package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadPlatform string

const (
	PayloadPlatformemail   PayloadPlatform = "Email"
	PayloadPlatformsms     PayloadPlatform = "Sms"
	PayloadPlatformteams   PayloadPlatform = "Teams"
	PayloadPlatformunknown PayloadPlatform = "Unknown"
)

func PossibleValuesForPayloadPlatform() []string {
	return []string{
		string(PayloadPlatformemail),
		string(PayloadPlatformsms),
		string(PayloadPlatformteams),
		string(PayloadPlatformunknown),
	}
}

func parsePayloadPlatform(input string) (*PayloadPlatform, error) {
	vals := map[string]PayloadPlatform{
		"email":   PayloadPlatformemail,
		"sms":     PayloadPlatformsms,
		"teams":   PayloadPlatformteams,
		"unknown": PayloadPlatformunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadPlatform(input)
	return &out, nil
}
