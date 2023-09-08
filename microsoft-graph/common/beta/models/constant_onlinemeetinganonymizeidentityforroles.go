package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAnonymizeIdentityForRoles string

const (
	OnlineMeetingAnonymizeIdentityForRolesattendee    OnlineMeetingAnonymizeIdentityForRoles = "Attendee"
	OnlineMeetingAnonymizeIdentityForRolescoorganizer OnlineMeetingAnonymizeIdentityForRoles = "Coorganizer"
	OnlineMeetingAnonymizeIdentityForRolespresenter   OnlineMeetingAnonymizeIdentityForRoles = "Presenter"
	OnlineMeetingAnonymizeIdentityForRolesproducer    OnlineMeetingAnonymizeIdentityForRoles = "Producer"
)

func PossibleValuesForOnlineMeetingAnonymizeIdentityForRoles() []string {
	return []string{
		string(OnlineMeetingAnonymizeIdentityForRolesattendee),
		string(OnlineMeetingAnonymizeIdentityForRolescoorganizer),
		string(OnlineMeetingAnonymizeIdentityForRolespresenter),
		string(OnlineMeetingAnonymizeIdentityForRolesproducer),
	}
}

func parseOnlineMeetingAnonymizeIdentityForRoles(input string) (*OnlineMeetingAnonymizeIdentityForRoles, error) {
	vals := map[string]OnlineMeetingAnonymizeIdentityForRoles{
		"attendee":    OnlineMeetingAnonymizeIdentityForRolesattendee,
		"coorganizer": OnlineMeetingAnonymizeIdentityForRolescoorganizer,
		"presenter":   OnlineMeetingAnonymizeIdentityForRolespresenter,
		"producer":    OnlineMeetingAnonymizeIdentityForRolesproducer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAnonymizeIdentityForRoles(input)
	return &out, nil
}
