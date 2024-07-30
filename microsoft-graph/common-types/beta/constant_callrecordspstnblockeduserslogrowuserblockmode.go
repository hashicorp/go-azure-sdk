package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnBlockedUsersLogRowUserBlockMode string

const (
	CallRecordsPstnBlockedUsersLogRowUserBlockMode_Blocked   CallRecordsPstnBlockedUsersLogRowUserBlockMode = "blocked"
	CallRecordsPstnBlockedUsersLogRowUserBlockMode_Unblocked CallRecordsPstnBlockedUsersLogRowUserBlockMode = "unblocked"
)

func PossibleValuesForCallRecordsPstnBlockedUsersLogRowUserBlockMode() []string {
	return []string{
		string(CallRecordsPstnBlockedUsersLogRowUserBlockMode_Blocked),
		string(CallRecordsPstnBlockedUsersLogRowUserBlockMode_Unblocked),
	}
}

func (s *CallRecordsPstnBlockedUsersLogRowUserBlockMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsPstnBlockedUsersLogRowUserBlockMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsPstnBlockedUsersLogRowUserBlockMode(input string) (*CallRecordsPstnBlockedUsersLogRowUserBlockMode, error) {
	vals := map[string]CallRecordsPstnBlockedUsersLogRowUserBlockMode{
		"blocked":   CallRecordsPstnBlockedUsersLogRowUserBlockMode_Blocked,
		"unblocked": CallRecordsPstnBlockedUsersLogRowUserBlockMode_Unblocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnBlockedUsersLogRowUserBlockMode(input)
	return &out, nil
}
