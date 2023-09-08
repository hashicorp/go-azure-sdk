package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfileStatusStatus string

const (
	EducationSynchronizationProfileStatusStatuserror           EducationSynchronizationProfileStatusStatus = "Error"
	EducationSynchronizationProfileStatusStatusextracting      EducationSynchronizationProfileStatusStatus = "Extracting"
	EducationSynchronizationProfileStatusStatusinProgress      EducationSynchronizationProfileStatusStatus = "InProgress"
	EducationSynchronizationProfileStatusStatuspaused          EducationSynchronizationProfileStatusStatus = "Paused"
	EducationSynchronizationProfileStatusStatusquarantined     EducationSynchronizationProfileStatusStatus = "Quarantined"
	EducationSynchronizationProfileStatusStatussuccess         EducationSynchronizationProfileStatusStatus = "Success"
	EducationSynchronizationProfileStatusStatusvalidating      EducationSynchronizationProfileStatusStatus = "Validating"
	EducationSynchronizationProfileStatusStatusvalidationError EducationSynchronizationProfileStatusStatus = "ValidationError"
)

func PossibleValuesForEducationSynchronizationProfileStatusStatus() []string {
	return []string{
		string(EducationSynchronizationProfileStatusStatuserror),
		string(EducationSynchronizationProfileStatusStatusextracting),
		string(EducationSynchronizationProfileStatusStatusinProgress),
		string(EducationSynchronizationProfileStatusStatuspaused),
		string(EducationSynchronizationProfileStatusStatusquarantined),
		string(EducationSynchronizationProfileStatusStatussuccess),
		string(EducationSynchronizationProfileStatusStatusvalidating),
		string(EducationSynchronizationProfileStatusStatusvalidationError),
	}
}

func parseEducationSynchronizationProfileStatusStatus(input string) (*EducationSynchronizationProfileStatusStatus, error) {
	vals := map[string]EducationSynchronizationProfileStatusStatus{
		"error":           EducationSynchronizationProfileStatusStatuserror,
		"extracting":      EducationSynchronizationProfileStatusStatusextracting,
		"inprogress":      EducationSynchronizationProfileStatusStatusinProgress,
		"paused":          EducationSynchronizationProfileStatusStatuspaused,
		"quarantined":     EducationSynchronizationProfileStatusStatusquarantined,
		"success":         EducationSynchronizationProfileStatusStatussuccess,
		"validating":      EducationSynchronizationProfileStatusStatusvalidating,
		"validationerror": EducationSynchronizationProfileStatusStatusvalidationError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationProfileStatusStatus(input)
	return &out, nil
}
