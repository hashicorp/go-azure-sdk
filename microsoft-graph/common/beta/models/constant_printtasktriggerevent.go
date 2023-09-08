package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskTriggerEvent string

const (
	PrintTaskTriggerEventjobStarted PrintTaskTriggerEvent = "JobStarted"
)

func PossibleValuesForPrintTaskTriggerEvent() []string {
	return []string{
		string(PrintTaskTriggerEventjobStarted),
	}
}

func parsePrintTaskTriggerEvent(input string) (*PrintTaskTriggerEvent, error) {
	vals := map[string]PrintTaskTriggerEvent{
		"jobstarted": PrintTaskTriggerEventjobStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskTriggerEvent(input)
	return &out, nil
}
