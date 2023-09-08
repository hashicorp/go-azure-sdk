package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsNotificationChangeType string

const (
	CommsNotificationChangeTypecreated CommsNotificationChangeType = "Created"
	CommsNotificationChangeTypedeleted CommsNotificationChangeType = "Deleted"
	CommsNotificationChangeTypeupdated CommsNotificationChangeType = "Updated"
)

func PossibleValuesForCommsNotificationChangeType() []string {
	return []string{
		string(CommsNotificationChangeTypecreated),
		string(CommsNotificationChangeTypedeleted),
		string(CommsNotificationChangeTypeupdated),
	}
}

func parseCommsNotificationChangeType(input string) (*CommsNotificationChangeType, error) {
	vals := map[string]CommsNotificationChangeType{
		"created": CommsNotificationChangeTypecreated,
		"deleted": CommsNotificationChangeTypedeleted,
		"updated": CommsNotificationChangeTypeupdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommsNotificationChangeType(input)
	return &out, nil
}
