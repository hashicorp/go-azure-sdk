package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationChangeType string

const (
	ChangeNotificationChangeTypecreated ChangeNotificationChangeType = "Created"
	ChangeNotificationChangeTypedeleted ChangeNotificationChangeType = "Deleted"
	ChangeNotificationChangeTypeupdated ChangeNotificationChangeType = "Updated"
)

func PossibleValuesForChangeNotificationChangeType() []string {
	return []string{
		string(ChangeNotificationChangeTypecreated),
		string(ChangeNotificationChangeTypedeleted),
		string(ChangeNotificationChangeTypeupdated),
	}
}

func parseChangeNotificationChangeType(input string) (*ChangeNotificationChangeType, error) {
	vals := map[string]ChangeNotificationChangeType{
		"created": ChangeNotificationChangeTypecreated,
		"deleted": ChangeNotificationChangeTypedeleted,
		"updated": ChangeNotificationChangeTypeupdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeNotificationChangeType(input)
	return &out, nil
}
