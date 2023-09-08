package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingAvailabilityStatus string

const (
	TrainingAvailabilityStatusarchive      TrainingAvailabilityStatus = "Archive"
	TrainingAvailabilityStatusavailable    TrainingAvailabilityStatus = "Available"
	TrainingAvailabilityStatusdelete       TrainingAvailabilityStatus = "Delete"
	TrainingAvailabilityStatusnotAvailable TrainingAvailabilityStatus = "NotAvailable"
	TrainingAvailabilityStatusunknown      TrainingAvailabilityStatus = "Unknown"
)

func PossibleValuesForTrainingAvailabilityStatus() []string {
	return []string{
		string(TrainingAvailabilityStatusarchive),
		string(TrainingAvailabilityStatusavailable),
		string(TrainingAvailabilityStatusdelete),
		string(TrainingAvailabilityStatusnotAvailable),
		string(TrainingAvailabilityStatusunknown),
	}
}

func parseTrainingAvailabilityStatus(input string) (*TrainingAvailabilityStatus, error) {
	vals := map[string]TrainingAvailabilityStatus{
		"archive":      TrainingAvailabilityStatusarchive,
		"available":    TrainingAvailabilityStatusavailable,
		"delete":       TrainingAvailabilityStatusdelete,
		"notavailable": TrainingAvailabilityStatusnotAvailable,
		"unknown":      TrainingAvailabilityStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingAvailabilityStatus(input)
	return &out, nil
}
