package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskTriggerEvent string

const (
	PrintTaskTriggerEvent_JobStarted PrintTaskTriggerEvent = "jobStarted"
)

func PossibleValuesForPrintTaskTriggerEvent() []string {
	return []string{
		string(PrintTaskTriggerEvent_JobStarted),
	}
}

func (s *PrintTaskTriggerEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintTaskTriggerEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintTaskTriggerEvent(input string) (*PrintTaskTriggerEvent, error) {
	vals := map[string]PrintTaskTriggerEvent{
		"jobstarted": PrintTaskTriggerEvent_JobStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskTriggerEvent(input)
	return &out, nil
}
