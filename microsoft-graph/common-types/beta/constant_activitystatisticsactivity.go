package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityStatisticsActivity string

const (
	ActivityStatisticsActivity_Call    ActivityStatisticsActivity = "Call"
	ActivityStatisticsActivity_Chat    ActivityStatisticsActivity = "Chat"
	ActivityStatisticsActivity_Email   ActivityStatisticsActivity = "Email"
	ActivityStatisticsActivity_Focus   ActivityStatisticsActivity = "Focus"
	ActivityStatisticsActivity_Meeting ActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForActivityStatisticsActivity() []string {
	return []string{
		string(ActivityStatisticsActivity_Call),
		string(ActivityStatisticsActivity_Chat),
		string(ActivityStatisticsActivity_Email),
		string(ActivityStatisticsActivity_Focus),
		string(ActivityStatisticsActivity_Meeting),
	}
}

func (s *ActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityStatisticsActivity(input string) (*ActivityStatisticsActivity, error) {
	vals := map[string]ActivityStatisticsActivity{
		"call":    ActivityStatisticsActivity_Call,
		"chat":    ActivityStatisticsActivity_Chat,
		"email":   ActivityStatisticsActivity_Email,
		"focus":   ActivityStatisticsActivity_Focus,
		"meeting": ActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityStatisticsActivity(input)
	return &out, nil
}
