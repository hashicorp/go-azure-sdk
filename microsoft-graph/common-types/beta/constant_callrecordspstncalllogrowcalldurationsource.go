package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallLogRowCallDurationSource string

const (
	CallRecordsPstnCallLogRowCallDurationSource_Microsoft CallRecordsPstnCallLogRowCallDurationSource = "microsoft"
	CallRecordsPstnCallLogRowCallDurationSource_Operator  CallRecordsPstnCallLogRowCallDurationSource = "operator"
)

func PossibleValuesForCallRecordsPstnCallLogRowCallDurationSource() []string {
	return []string{
		string(CallRecordsPstnCallLogRowCallDurationSource_Microsoft),
		string(CallRecordsPstnCallLogRowCallDurationSource_Operator),
	}
}

func (s *CallRecordsPstnCallLogRowCallDurationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsPstnCallLogRowCallDurationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsPstnCallLogRowCallDurationSource(input string) (*CallRecordsPstnCallLogRowCallDurationSource, error) {
	vals := map[string]CallRecordsPstnCallLogRowCallDurationSource{
		"microsoft": CallRecordsPstnCallLogRowCallDurationSource_Microsoft,
		"operator":  CallRecordsPstnCallLogRowCallDurationSource_Operator,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnCallLogRowCallDurationSource(input)
	return &out, nil
}
