package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsFailureInfoStage string

const (
	CallRecordsFailureInfoStage_CallSetup CallRecordsFailureInfoStage = "callSetup"
	CallRecordsFailureInfoStage_Midcall   CallRecordsFailureInfoStage = "midcall"
	CallRecordsFailureInfoStage_Unknown   CallRecordsFailureInfoStage = "unknown"
)

func PossibleValuesForCallRecordsFailureInfoStage() []string {
	return []string{
		string(CallRecordsFailureInfoStage_CallSetup),
		string(CallRecordsFailureInfoStage_Midcall),
		string(CallRecordsFailureInfoStage_Unknown),
	}
}

func (s *CallRecordsFailureInfoStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsFailureInfoStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsFailureInfoStage(input string) (*CallRecordsFailureInfoStage, error) {
	vals := map[string]CallRecordsFailureInfoStage{
		"callsetup": CallRecordsFailureInfoStage_CallSetup,
		"midcall":   CallRecordsFailureInfoStage_Midcall,
		"unknown":   CallRecordsFailureInfoStage_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsFailureInfoStage(input)
	return &out, nil
}
