package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingActionAllowedForwarders string

const (
	ProtectOnlineMeetingActionAllowedForwarderseveryone  ProtectOnlineMeetingActionAllowedForwarders = "Everyone"
	ProtectOnlineMeetingActionAllowedForwardersorganizer ProtectOnlineMeetingActionAllowedForwarders = "Organizer"
)

func PossibleValuesForProtectOnlineMeetingActionAllowedForwarders() []string {
	return []string{
		string(ProtectOnlineMeetingActionAllowedForwarderseveryone),
		string(ProtectOnlineMeetingActionAllowedForwardersorganizer),
	}
}

func parseProtectOnlineMeetingActionAllowedForwarders(input string) (*ProtectOnlineMeetingActionAllowedForwarders, error) {
	vals := map[string]ProtectOnlineMeetingActionAllowedForwarders{
		"everyone":  ProtectOnlineMeetingActionAllowedForwarderseveryone,
		"organizer": ProtectOnlineMeetingActionAllowedForwardersorganizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectOnlineMeetingActionAllowedForwarders(input)
	return &out, nil
}
