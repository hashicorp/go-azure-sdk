package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityStatus string

const (
	UserActivityStatusactive  UserActivityStatus = "Active"
	UserActivityStatusdeleted UserActivityStatus = "Deleted"
	UserActivityStatusignored UserActivityStatus = "Ignored"
	UserActivityStatusupdated UserActivityStatus = "Updated"
)

func PossibleValuesForUserActivityStatus() []string {
	return []string{
		string(UserActivityStatusactive),
		string(UserActivityStatusdeleted),
		string(UserActivityStatusignored),
		string(UserActivityStatusupdated),
	}
}

func parseUserActivityStatus(input string) (*UserActivityStatus, error) {
	vals := map[string]UserActivityStatus{
		"active":  UserActivityStatusactive,
		"deleted": UserActivityStatusdeleted,
		"ignored": UserActivityStatusignored,
		"updated": UserActivityStatusupdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserActivityStatus(input)
	return &out, nil
}
