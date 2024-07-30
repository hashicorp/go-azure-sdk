package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationQuestionAnswerInputType string

const (
	MeetingRegistrationQuestionAnswerInputType_RadioButton MeetingRegistrationQuestionAnswerInputType = "radioButton"
	MeetingRegistrationQuestionAnswerInputType_Text        MeetingRegistrationQuestionAnswerInputType = "text"
)

func PossibleValuesForMeetingRegistrationQuestionAnswerInputType() []string {
	return []string{
		string(MeetingRegistrationQuestionAnswerInputType_RadioButton),
		string(MeetingRegistrationQuestionAnswerInputType_Text),
	}
}

func (s *MeetingRegistrationQuestionAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrationQuestionAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrationQuestionAnswerInputType(input string) (*MeetingRegistrationQuestionAnswerInputType, error) {
	vals := map[string]MeetingRegistrationQuestionAnswerInputType{
		"radiobutton": MeetingRegistrationQuestionAnswerInputType_RadioButton,
		"text":        MeetingRegistrationQuestionAnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationQuestionAnswerInputType(input)
	return &out, nil
}
