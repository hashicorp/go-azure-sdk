package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InitiatorInitiatorType string

const (
	InitiatorInitiatorType_Application InitiatorInitiatorType = "application"
	InitiatorInitiatorType_System      InitiatorInitiatorType = "system"
	InitiatorInitiatorType_User        InitiatorInitiatorType = "user"
)

func PossibleValuesForInitiatorInitiatorType() []string {
	return []string{
		string(InitiatorInitiatorType_Application),
		string(InitiatorInitiatorType_System),
		string(InitiatorInitiatorType_User),
	}
}

func (s *InitiatorInitiatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInitiatorInitiatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInitiatorInitiatorType(input string) (*InitiatorInitiatorType, error) {
	vals := map[string]InitiatorInitiatorType{
		"application": InitiatorInitiatorType_Application,
		"system":      InitiatorInitiatorType_System,
		"user":        InitiatorInitiatorType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InitiatorInitiatorType(input)
	return &out, nil
}
