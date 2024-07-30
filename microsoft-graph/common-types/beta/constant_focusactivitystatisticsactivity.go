package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FocusActivityStatisticsActivity string

const (
	FocusActivityStatisticsActivity_Call    FocusActivityStatisticsActivity = "Call"
	FocusActivityStatisticsActivity_Chat    FocusActivityStatisticsActivity = "Chat"
	FocusActivityStatisticsActivity_Email   FocusActivityStatisticsActivity = "Email"
	FocusActivityStatisticsActivity_Focus   FocusActivityStatisticsActivity = "Focus"
	FocusActivityStatisticsActivity_Meeting FocusActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForFocusActivityStatisticsActivity() []string {
	return []string{
		string(FocusActivityStatisticsActivity_Call),
		string(FocusActivityStatisticsActivity_Chat),
		string(FocusActivityStatisticsActivity_Email),
		string(FocusActivityStatisticsActivity_Focus),
		string(FocusActivityStatisticsActivity_Meeting),
	}
}

func (s *FocusActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFocusActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFocusActivityStatisticsActivity(input string) (*FocusActivityStatisticsActivity, error) {
	vals := map[string]FocusActivityStatisticsActivity{
		"call":    FocusActivityStatisticsActivity_Call,
		"chat":    FocusActivityStatisticsActivity_Chat,
		"email":   FocusActivityStatisticsActivity_Email,
		"focus":   FocusActivityStatisticsActivity_Focus,
		"meeting": FocusActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FocusActivityStatisticsActivity(input)
	return &out, nil
}
