package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailActivityStatisticsActivity string

const (
	EmailActivityStatisticsActivity_Call    EmailActivityStatisticsActivity = "Call"
	EmailActivityStatisticsActivity_Chat    EmailActivityStatisticsActivity = "Chat"
	EmailActivityStatisticsActivity_Email   EmailActivityStatisticsActivity = "Email"
	EmailActivityStatisticsActivity_Focus   EmailActivityStatisticsActivity = "Focus"
	EmailActivityStatisticsActivity_Meeting EmailActivityStatisticsActivity = "Meeting"
)

func PossibleValuesForEmailActivityStatisticsActivity() []string {
	return []string{
		string(EmailActivityStatisticsActivity_Call),
		string(EmailActivityStatisticsActivity_Chat),
		string(EmailActivityStatisticsActivity_Email),
		string(EmailActivityStatisticsActivity_Focus),
		string(EmailActivityStatisticsActivity_Meeting),
	}
}

func (s *EmailActivityStatisticsActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailActivityStatisticsActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailActivityStatisticsActivity(input string) (*EmailActivityStatisticsActivity, error) {
	vals := map[string]EmailActivityStatisticsActivity{
		"call":    EmailActivityStatisticsActivity_Call,
		"chat":    EmailActivityStatisticsActivity_Chat,
		"email":   EmailActivityStatisticsActivity_Email,
		"focus":   EmailActivityStatisticsActivity_Focus,
		"meeting": EmailActivityStatisticsActivity_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailActivityStatisticsActivity(input)
	return &out, nil
}
