package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingStatusInfoTrainingStatus string

const (
	UserTrainingStatusInfoTrainingStatus_Assigned   UserTrainingStatusInfoTrainingStatus = "assigned"
	UserTrainingStatusInfoTrainingStatus_Completed  UserTrainingStatusInfoTrainingStatus = "completed"
	UserTrainingStatusInfoTrainingStatus_InProgress UserTrainingStatusInfoTrainingStatus = "inProgress"
	UserTrainingStatusInfoTrainingStatus_Overdue    UserTrainingStatusInfoTrainingStatus = "overdue"
	UserTrainingStatusInfoTrainingStatus_Unknown    UserTrainingStatusInfoTrainingStatus = "unknown"
)

func PossibleValuesForUserTrainingStatusInfoTrainingStatus() []string {
	return []string{
		string(UserTrainingStatusInfoTrainingStatus_Assigned),
		string(UserTrainingStatusInfoTrainingStatus_Completed),
		string(UserTrainingStatusInfoTrainingStatus_InProgress),
		string(UserTrainingStatusInfoTrainingStatus_Overdue),
		string(UserTrainingStatusInfoTrainingStatus_Unknown),
	}
}

func (s *UserTrainingStatusInfoTrainingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserTrainingStatusInfoTrainingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserTrainingStatusInfoTrainingStatus(input string) (*UserTrainingStatusInfoTrainingStatus, error) {
	vals := map[string]UserTrainingStatusInfoTrainingStatus{
		"assigned":   UserTrainingStatusInfoTrainingStatus_Assigned,
		"completed":  UserTrainingStatusInfoTrainingStatus_Completed,
		"inprogress": UserTrainingStatusInfoTrainingStatus_InProgress,
		"overdue":    UserTrainingStatusInfoTrainingStatus_Overdue,
		"unknown":    UserTrainingStatusInfoTrainingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserTrainingStatusInfoTrainingStatus(input)
	return &out, nil
}
