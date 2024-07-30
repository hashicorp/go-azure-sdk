package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfileStatusStatus string

const (
	EducationSynchronizationProfileStatusStatus_Error           EducationSynchronizationProfileStatusStatus = "error"
	EducationSynchronizationProfileStatusStatus_Extracting      EducationSynchronizationProfileStatusStatus = "extracting"
	EducationSynchronizationProfileStatusStatus_InProgress      EducationSynchronizationProfileStatusStatus = "inProgress"
	EducationSynchronizationProfileStatusStatus_Paused          EducationSynchronizationProfileStatusStatus = "paused"
	EducationSynchronizationProfileStatusStatus_Quarantined     EducationSynchronizationProfileStatusStatus = "quarantined"
	EducationSynchronizationProfileStatusStatus_Success         EducationSynchronizationProfileStatusStatus = "success"
	EducationSynchronizationProfileStatusStatus_Validating      EducationSynchronizationProfileStatusStatus = "validating"
	EducationSynchronizationProfileStatusStatus_ValidationError EducationSynchronizationProfileStatusStatus = "validationError"
)

func PossibleValuesForEducationSynchronizationProfileStatusStatus() []string {
	return []string{
		string(EducationSynchronizationProfileStatusStatus_Error),
		string(EducationSynchronizationProfileStatusStatus_Extracting),
		string(EducationSynchronizationProfileStatusStatus_InProgress),
		string(EducationSynchronizationProfileStatusStatus_Paused),
		string(EducationSynchronizationProfileStatusStatus_Quarantined),
		string(EducationSynchronizationProfileStatusStatus_Success),
		string(EducationSynchronizationProfileStatusStatus_Validating),
		string(EducationSynchronizationProfileStatusStatus_ValidationError),
	}
}

func (s *EducationSynchronizationProfileStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSynchronizationProfileStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSynchronizationProfileStatusStatus(input string) (*EducationSynchronizationProfileStatusStatus, error) {
	vals := map[string]EducationSynchronizationProfileStatusStatus{
		"error":           EducationSynchronizationProfileStatusStatus_Error,
		"extracting":      EducationSynchronizationProfileStatusStatus_Extracting,
		"inprogress":      EducationSynchronizationProfileStatusStatus_InProgress,
		"paused":          EducationSynchronizationProfileStatusStatus_Paused,
		"quarantined":     EducationSynchronizationProfileStatusStatus_Quarantined,
		"success":         EducationSynchronizationProfileStatusStatus_Success,
		"validating":      EducationSynchronizationProfileStatusStatus_Validating,
		"validationerror": EducationSynchronizationProfileStatusStatus_ValidationError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationProfileStatusStatus(input)
	return &out, nil
}
