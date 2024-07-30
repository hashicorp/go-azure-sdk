package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomQuestionAnswerInputType string

const (
	BookingCustomQuestionAnswerInputType_RadioButton BookingCustomQuestionAnswerInputType = "radioButton"
	BookingCustomQuestionAnswerInputType_Text        BookingCustomQuestionAnswerInputType = "text"
)

func PossibleValuesForBookingCustomQuestionAnswerInputType() []string {
	return []string{
		string(BookingCustomQuestionAnswerInputType_RadioButton),
		string(BookingCustomQuestionAnswerInputType_Text),
	}
}

func (s *BookingCustomQuestionAnswerInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingCustomQuestionAnswerInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingCustomQuestionAnswerInputType(input string) (*BookingCustomQuestionAnswerInputType, error) {
	vals := map[string]BookingCustomQuestionAnswerInputType{
		"radiobutton": BookingCustomQuestionAnswerInputType_RadioButton,
		"text":        BookingCustomQuestionAnswerInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingCustomQuestionAnswerInputType(input)
	return &out, nil
}
