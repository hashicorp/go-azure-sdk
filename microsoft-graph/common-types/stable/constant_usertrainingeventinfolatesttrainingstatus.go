package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingEventInfoLatestTrainingStatus string

const (
	UserTrainingEventInfoLatestTrainingStatus_Assigned   UserTrainingEventInfoLatestTrainingStatus = "assigned"
	UserTrainingEventInfoLatestTrainingStatus_Completed  UserTrainingEventInfoLatestTrainingStatus = "completed"
	UserTrainingEventInfoLatestTrainingStatus_InProgress UserTrainingEventInfoLatestTrainingStatus = "inProgress"
	UserTrainingEventInfoLatestTrainingStatus_Overdue    UserTrainingEventInfoLatestTrainingStatus = "overdue"
	UserTrainingEventInfoLatestTrainingStatus_Unknown    UserTrainingEventInfoLatestTrainingStatus = "unknown"
)

func PossibleValuesForUserTrainingEventInfoLatestTrainingStatus() []string {
	return []string{
		string(UserTrainingEventInfoLatestTrainingStatus_Assigned),
		string(UserTrainingEventInfoLatestTrainingStatus_Completed),
		string(UserTrainingEventInfoLatestTrainingStatus_InProgress),
		string(UserTrainingEventInfoLatestTrainingStatus_Overdue),
		string(UserTrainingEventInfoLatestTrainingStatus_Unknown),
	}
}

func (s *UserTrainingEventInfoLatestTrainingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserTrainingEventInfoLatestTrainingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserTrainingEventInfoLatestTrainingStatus(input string) (*UserTrainingEventInfoLatestTrainingStatus, error) {
	vals := map[string]UserTrainingEventInfoLatestTrainingStatus{
		"assigned":   UserTrainingEventInfoLatestTrainingStatus_Assigned,
		"completed":  UserTrainingEventInfoLatestTrainingStatus_Completed,
		"inprogress": UserTrainingEventInfoLatestTrainingStatus_InProgress,
		"overdue":    UserTrainingEventInfoLatestTrainingStatus_Overdue,
		"unknown":    UserTrainingEventInfoLatestTrainingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserTrainingEventInfoLatestTrainingStatus(input)
	return &out, nil
}
