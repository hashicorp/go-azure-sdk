package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsUserPurpose string

const (
	MailboxSettingsUserPurpose_Equipment MailboxSettingsUserPurpose = "equipment"
	MailboxSettingsUserPurpose_Linked    MailboxSettingsUserPurpose = "linked"
	MailboxSettingsUserPurpose_Others    MailboxSettingsUserPurpose = "others"
	MailboxSettingsUserPurpose_Room      MailboxSettingsUserPurpose = "room"
	MailboxSettingsUserPurpose_Shared    MailboxSettingsUserPurpose = "shared"
	MailboxSettingsUserPurpose_Unknown   MailboxSettingsUserPurpose = "unknown"
	MailboxSettingsUserPurpose_User      MailboxSettingsUserPurpose = "user"
)

func PossibleValuesForMailboxSettingsUserPurpose() []string {
	return []string{
		string(MailboxSettingsUserPurpose_Equipment),
		string(MailboxSettingsUserPurpose_Linked),
		string(MailboxSettingsUserPurpose_Others),
		string(MailboxSettingsUserPurpose_Room),
		string(MailboxSettingsUserPurpose_Shared),
		string(MailboxSettingsUserPurpose_Unknown),
		string(MailboxSettingsUserPurpose_User),
	}
}

func (s *MailboxSettingsUserPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsUserPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsUserPurpose(input string) (*MailboxSettingsUserPurpose, error) {
	vals := map[string]MailboxSettingsUserPurpose{
		"equipment": MailboxSettingsUserPurpose_Equipment,
		"linked":    MailboxSettingsUserPurpose_Linked,
		"others":    MailboxSettingsUserPurpose_Others,
		"room":      MailboxSettingsUserPurpose_Room,
		"shared":    MailboxSettingsUserPurpose_Shared,
		"unknown":   MailboxSettingsUserPurpose_Unknown,
		"user":      MailboxSettingsUserPurpose_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurpose(input)
	return &out, nil
}
