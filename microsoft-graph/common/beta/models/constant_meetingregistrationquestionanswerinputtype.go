package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationQuestionAnswerInputType string

const (
	MeetingRegistrationQuestionAnswerInputTyperadioButton MeetingRegistrationQuestionAnswerInputType = "RadioButton"
	MeetingRegistrationQuestionAnswerInputTypetext        MeetingRegistrationQuestionAnswerInputType = "Text"
)

func PossibleValuesForMeetingRegistrationQuestionAnswerInputType() []string {
	return []string{
		string(MeetingRegistrationQuestionAnswerInputTyperadioButton),
		string(MeetingRegistrationQuestionAnswerInputTypetext),
	}
}

func parseMeetingRegistrationQuestionAnswerInputType(input string) (*MeetingRegistrationQuestionAnswerInputType, error) {
	vals := map[string]MeetingRegistrationQuestionAnswerInputType{
		"radiobutton": MeetingRegistrationQuestionAnswerInputTyperadioButton,
		"text":        MeetingRegistrationQuestionAnswerInputTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationQuestionAnswerInputType(input)
	return &out, nil
}
