package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsUserPurposeV2 string

const (
	MailboxSettingsUserPurposeV2_Equipment MailboxSettingsUserPurposeV2 = "equipment"
	MailboxSettingsUserPurposeV2_Linked    MailboxSettingsUserPurposeV2 = "linked"
	MailboxSettingsUserPurposeV2_Others    MailboxSettingsUserPurposeV2 = "others"
	MailboxSettingsUserPurposeV2_Room      MailboxSettingsUserPurposeV2 = "room"
	MailboxSettingsUserPurposeV2_Shared    MailboxSettingsUserPurposeV2 = "shared"
	MailboxSettingsUserPurposeV2_Unknown   MailboxSettingsUserPurposeV2 = "unknown"
	MailboxSettingsUserPurposeV2_User      MailboxSettingsUserPurposeV2 = "user"
)

func PossibleValuesForMailboxSettingsUserPurposeV2() []string {
	return []string{
		string(MailboxSettingsUserPurposeV2_Equipment),
		string(MailboxSettingsUserPurposeV2_Linked),
		string(MailboxSettingsUserPurposeV2_Others),
		string(MailboxSettingsUserPurposeV2_Room),
		string(MailboxSettingsUserPurposeV2_Shared),
		string(MailboxSettingsUserPurposeV2_Unknown),
		string(MailboxSettingsUserPurposeV2_User),
	}
}

func (s *MailboxSettingsUserPurposeV2) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsUserPurposeV2(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsUserPurposeV2(input string) (*MailboxSettingsUserPurposeV2, error) {
	vals := map[string]MailboxSettingsUserPurposeV2{
		"equipment": MailboxSettingsUserPurposeV2_Equipment,
		"linked":    MailboxSettingsUserPurposeV2_Linked,
		"others":    MailboxSettingsUserPurposeV2_Others,
		"room":      MailboxSettingsUserPurposeV2_Room,
		"shared":    MailboxSettingsUserPurposeV2_Shared,
		"unknown":   MailboxSettingsUserPurposeV2_Unknown,
		"user":      MailboxSettingsUserPurposeV2_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsUserPurposeV2(input)
	return &out, nil
}
