package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityStatus string

const (
	UserActivityStatus_Active  UserActivityStatus = "active"
	UserActivityStatus_Deleted UserActivityStatus = "deleted"
	UserActivityStatus_Ignored UserActivityStatus = "ignored"
	UserActivityStatus_Updated UserActivityStatus = "updated"
)

func PossibleValuesForUserActivityStatus() []string {
	return []string{
		string(UserActivityStatus_Active),
		string(UserActivityStatus_Deleted),
		string(UserActivityStatus_Ignored),
		string(UserActivityStatus_Updated),
	}
}

func (s *UserActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserActivityStatus(input string) (*UserActivityStatus, error) {
	vals := map[string]UserActivityStatus{
		"active":  UserActivityStatus_Active,
		"deleted": UserActivityStatus_Deleted,
		"ignored": UserActivityStatus_Ignored,
		"updated": UserActivityStatus_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserActivityStatus(input)
	return &out, nil
}
