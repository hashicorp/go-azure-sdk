package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingCapabilityAutoAdmittedUsers string

const (
	MeetingCapabilityAutoAdmittedUserseveryone          MeetingCapabilityAutoAdmittedUsers = "Everyone"
	MeetingCapabilityAutoAdmittedUserseveryoneInCompany MeetingCapabilityAutoAdmittedUsers = "EveryoneInCompany"
)

func PossibleValuesForMeetingCapabilityAutoAdmittedUsers() []string {
	return []string{
		string(MeetingCapabilityAutoAdmittedUserseveryone),
		string(MeetingCapabilityAutoAdmittedUserseveryoneInCompany),
	}
}

func parseMeetingCapabilityAutoAdmittedUsers(input string) (*MeetingCapabilityAutoAdmittedUsers, error) {
	vals := map[string]MeetingCapabilityAutoAdmittedUsers{
		"everyone":          MeetingCapabilityAutoAdmittedUserseveryone,
		"everyoneincompany": MeetingCapabilityAutoAdmittedUserseveryoneInCompany,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingCapabilityAutoAdmittedUsers(input)
	return &out, nil
}
