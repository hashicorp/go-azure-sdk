package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnBlockedUsersLogRowUserBlockMode string

const (
	CallRecordsPstnBlockedUsersLogRowUserBlockModeblocked   CallRecordsPstnBlockedUsersLogRowUserBlockMode = "Blocked"
	CallRecordsPstnBlockedUsersLogRowUserBlockModeunblocked CallRecordsPstnBlockedUsersLogRowUserBlockMode = "Unblocked"
)

func PossibleValuesForCallRecordsPstnBlockedUsersLogRowUserBlockMode() []string {
	return []string{
		string(CallRecordsPstnBlockedUsersLogRowUserBlockModeblocked),
		string(CallRecordsPstnBlockedUsersLogRowUserBlockModeunblocked),
	}
}

func parseCallRecordsPstnBlockedUsersLogRowUserBlockMode(input string) (*CallRecordsPstnBlockedUsersLogRowUserBlockMode, error) {
	vals := map[string]CallRecordsPstnBlockedUsersLogRowUserBlockMode{
		"blocked":   CallRecordsPstnBlockedUsersLogRowUserBlockModeblocked,
		"unblocked": CallRecordsPstnBlockedUsersLogRowUserBlockModeunblocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnBlockedUsersLogRowUserBlockMode(input)
	return &out, nil
}
