package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsNotificationChangeType string

const (
	CommsNotificationChangeType_Created CommsNotificationChangeType = "created"
	CommsNotificationChangeType_Deleted CommsNotificationChangeType = "deleted"
	CommsNotificationChangeType_Updated CommsNotificationChangeType = "updated"
)

func PossibleValuesForCommsNotificationChangeType() []string {
	return []string{
		string(CommsNotificationChangeType_Created),
		string(CommsNotificationChangeType_Deleted),
		string(CommsNotificationChangeType_Updated),
	}
}

func (s *CommsNotificationChangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommsNotificationChangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommsNotificationChangeType(input string) (*CommsNotificationChangeType, error) {
	vals := map[string]CommsNotificationChangeType{
		"created": CommsNotificationChangeType_Created,
		"deleted": CommsNotificationChangeType_Deleted,
		"updated": CommsNotificationChangeType_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommsNotificationChangeType(input)
	return &out, nil
}
