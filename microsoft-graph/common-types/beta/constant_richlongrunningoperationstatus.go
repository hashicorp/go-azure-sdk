package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RichLongRunningOperationStatus string

const (
	RichLongRunningOperationStatus_Failed     RichLongRunningOperationStatus = "failed"
	RichLongRunningOperationStatus_NotStarted RichLongRunningOperationStatus = "notStarted"
	RichLongRunningOperationStatus_Running    RichLongRunningOperationStatus = "running"
	RichLongRunningOperationStatus_Succeeded  RichLongRunningOperationStatus = "succeeded"
)

func PossibleValuesForRichLongRunningOperationStatus() []string {
	return []string{
		string(RichLongRunningOperationStatus_Failed),
		string(RichLongRunningOperationStatus_NotStarted),
		string(RichLongRunningOperationStatus_Running),
		string(RichLongRunningOperationStatus_Succeeded),
	}
}

func (s *RichLongRunningOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRichLongRunningOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRichLongRunningOperationStatus(input string) (*RichLongRunningOperationStatus, error) {
	vals := map[string]RichLongRunningOperationStatus{
		"failed":     RichLongRunningOperationStatus_Failed,
		"notstarted": RichLongRunningOperationStatus_NotStarted,
		"running":    RichLongRunningOperationStatus_Running,
		"succeeded":  RichLongRunningOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RichLongRunningOperationStatus(input)
	return &out, nil
}
