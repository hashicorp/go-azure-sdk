package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsUserPurposeV2 string

const (
	MailboxSettingsUserPurposeV2equipment MailboxSettingsUserPurposeV2 = "Equipment"
	MailboxSettingsUserPurposeV2linked    MailboxSettingsUserPurposeV2 = "Linked"
	MailboxSettingsUserPurposeV2others    MailboxSettingsUserPurposeV2 = "Others"
	MailboxSettingsUserPurposeV2room      MailboxSettingsUserPurposeV2 = "Room"
	MailboxSettingsUserPurposeV2shared    MailboxSettingsUserPurposeV2 = "Shared"
	MailboxSettingsUserPurposeV2unknown   MailboxSettingsUserPurposeV2 = "Unknown"
	MailboxSettingsUserPurposeV2user      MailboxSettingsUserPurposeV2 = "User"
)

func PossibleValuesForMailboxSettingsUserPurposeV2() []string {
	return []string{
		string(MailboxSettingsUserPurposeV2equipment),
		string(MailboxSettingsUserPurposeV2linked),
		string(MailboxSettingsUserPurposeV2others),
		string(MailboxSettingsUserPurposeV2room),
		string(MailboxSettingsUserPurposeV2shared),
		string(MailboxSettingsUserPurposeV2unknown),
		string(MailboxSettingsUserPurposeV2user),
	}
}

func parseMailboxSettingsUserPurposeV2(input string) (*MailboxSettingsUserPurposeV2, error) {
	vals := map[string]MailboxSettingsUserPurposeV2{
		"equipment": MailboxSettingsUserPurposeV2equipment,
		"linked":    MailboxSettingsUserPurposeV2linked,
		"others":    MailboxSettingsUserPurposeV2others,
		"room":      MailboxSettingsUserPurposeV2room,
		"shared":    MailboxSettingsUserPurposeV2shared,
		"unknown":   MailboxSettingsUserPurposeV2unknown,
		"user":      MailboxSettingsUserPurposeV2user,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurposeV2(input)
	return &out, nil
}
