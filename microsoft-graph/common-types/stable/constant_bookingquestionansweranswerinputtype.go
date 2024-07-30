package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAnswerAnswerInputType string

const (
	BookingQuestionAnswerAnswerInputType_RadioButton BookingQuestionAnswerAnswerInputType = "radioButton"
	BookingQuestionAnswerAnswerInputType_Text        BookingQuestionAnswerAnswerInputType = "text"
)

func PossibleValuesForBookingQuestionAnswerAnswerInputType() []string {
	return []string{
		string(BookingQuestionAnswerAnswerInputType_RadioButton),
		string(BookingQuestionAnswerAnswerInputType_Text),
	}
}

func (s *BookingQuestionAnswerAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingQuestionAnswerAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingQuestionAnswerAnswerInputType(input string) (*BookingQuestionAnswerAnswerInputType, error) {
	vals := map[string]BookingQuestionAnswerAnswerInputType{
		"radiobutton": BookingQuestionAnswerAnswerInputType_RadioButton,
		"text":        BookingQuestionAnswerAnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingQuestionAnswerAnswerInputType(input)
	return &out, nil
}
