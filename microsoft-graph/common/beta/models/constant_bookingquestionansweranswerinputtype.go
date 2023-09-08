package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAnswerAnswerInputType string

const (
	BookingQuestionAnswerAnswerInputTyperadioButton BookingQuestionAnswerAnswerInputType = "RadioButton"
	BookingQuestionAnswerAnswerInputTypetext        BookingQuestionAnswerAnswerInputType = "Text"
)

func PossibleValuesForBookingQuestionAnswerAnswerInputType() []string {
	return []string{
		string(BookingQuestionAnswerAnswerInputTyperadioButton),
		string(BookingQuestionAnswerAnswerInputTypetext),
	}
}

func parseBookingQuestionAnswerAnswerInputType(input string) (*BookingQuestionAnswerAnswerInputType, error) {
	vals := map[string]BookingQuestionAnswerAnswerInputType{
		"radiobutton": BookingQuestionAnswerAnswerInputTyperadioButton,
		"text":        BookingQuestionAnswerAnswerInputTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingQuestionAnswerAnswerInputType(input)
	return &out, nil
}
