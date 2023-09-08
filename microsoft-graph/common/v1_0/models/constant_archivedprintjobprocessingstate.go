package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJobProcessingState string

const (
	ArchivedPrintJobProcessingStateaborted    ArchivedPrintJobProcessingState = "Aborted"
	ArchivedPrintJobProcessingStatecanceled   ArchivedPrintJobProcessingState = "Canceled"
	ArchivedPrintJobProcessingStatecompleted  ArchivedPrintJobProcessingState = "Completed"
	ArchivedPrintJobProcessingStatepaused     ArchivedPrintJobProcessingState = "Paused"
	ArchivedPrintJobProcessingStatepending    ArchivedPrintJobProcessingState = "Pending"
	ArchivedPrintJobProcessingStateprocessing ArchivedPrintJobProcessingState = "Processing"
	ArchivedPrintJobProcessingStatestopped    ArchivedPrintJobProcessingState = "Stopped"
	ArchivedPrintJobProcessingStateunknown    ArchivedPrintJobProcessingState = "Unknown"
)

func PossibleValuesForArchivedPrintJobProcessingState() []string {
	return []string{
		string(ArchivedPrintJobProcessingStateaborted),
		string(ArchivedPrintJobProcessingStatecanceled),
		string(ArchivedPrintJobProcessingStatecompleted),
		string(ArchivedPrintJobProcessingStatepaused),
		string(ArchivedPrintJobProcessingStatepending),
		string(ArchivedPrintJobProcessingStateprocessing),
		string(ArchivedPrintJobProcessingStatestopped),
		string(ArchivedPrintJobProcessingStateunknown),
	}
}

func parseArchivedPrintJobProcessingState(input string) (*ArchivedPrintJobProcessingState, error) {
	vals := map[string]ArchivedPrintJobProcessingState{
		"aborted":    ArchivedPrintJobProcessingStateaborted,
		"canceled":   ArchivedPrintJobProcessingStatecanceled,
		"completed":  ArchivedPrintJobProcessingStatecompleted,
		"paused":     ArchivedPrintJobProcessingStatepaused,
		"pending":    ArchivedPrintJobProcessingStatepending,
		"processing": ArchivedPrintJobProcessingStateprocessing,
		"stopped":    ArchivedPrintJobProcessingStatestopped,
		"unknown":    ArchivedPrintJobProcessingStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArchivedPrintJobProcessingState(input)
	return &out, nil
}
