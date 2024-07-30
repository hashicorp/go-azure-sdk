package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallActivityStatisticsActivity string

const (
	CallActivityStatisticsActivity_Call    CallActivityStatisticsActivity = "Call"
	CallActivityStatisticsActivity_Chat    CallActivityStatisticsActivity = "Chat"
	CallActivityStatisticsActivity_Email   CallActivityStatisticsActivity = "Email"
	CallActivityStatisticsActivity_Focus   CallActivityStatisticsActivity = "Focus"
	CallActivityStatisticsActivity_Meeting CallActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForCallActivityStatisticsActivity() []string {
	return []string{
		string(CallActivityStatisticsActivity_Call),
		string(CallActivityStatisticsActivity_Chat),
		string(CallActivityStatisticsActivity_Email),
		string(CallActivityStatisticsActivity_Focus),
		string(CallActivityStatisticsActivity_Meeting),
	}
}

func (s *CallActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallActivityStatisticsActivity(input string) (*CallActivityStatisticsActivity, error) {
	vals := map[string]CallActivityStatisticsActivity{
		"call":    CallActivityStatisticsActivity_Call,
		"chat":    CallActivityStatisticsActivity_Chat,
		"email":   CallActivityStatisticsActivity_Email,
		"focus":   CallActivityStatisticsActivity_Focus,
		"meeting": CallActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallActivityStatisticsActivity(input)
	return &out, nil
}
