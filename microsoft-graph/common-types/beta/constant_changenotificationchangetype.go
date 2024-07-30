package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationChangeType string

const (
	ChangeNotificationChangeType_Created ChangeNotificationChangeType = "created"
	ChangeNotificationChangeType_Deleted ChangeNotificationChangeType = "deleted"
	ChangeNotificationChangeType_Updated ChangeNotificationChangeType = "updated"
)

func PossibleValuesForChangeNotificationChangeType() []string {
	return []string{
		string(ChangeNotificationChangeType_Created),
		string(ChangeNotificationChangeType_Deleted),
		string(ChangeNotificationChangeType_Updated),
	}
}

func (s *ChangeNotificationChangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeNotificationChangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeNotificationChangeType(input string) (*ChangeNotificationChangeType, error) {
	vals := map[string]ChangeNotificationChangeType{
		"created": ChangeNotificationChangeType_Created,
		"deleted": ChangeNotificationChangeType_Deleted,
		"updated": ChangeNotificationChangeType_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeNotificationChangeType(input)
	return &out, nil
}
