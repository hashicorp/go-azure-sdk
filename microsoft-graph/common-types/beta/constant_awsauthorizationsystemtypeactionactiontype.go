package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAuthorizationSystemTypeActionActionType string

const (
	AwsAuthorizationSystemTypeActionActionType_Delete AwsAuthorizationSystemTypeActionActionType = "delete"
	AwsAuthorizationSystemTypeActionActionType_Read   AwsAuthorizationSystemTypeActionActionType = "read"
)

func PossibleValuesForAwsAuthorizationSystemTypeActionActionType() []string {
	return []string{
		string(AwsAuthorizationSystemTypeActionActionType_Delete),
		string(AwsAuthorizationSystemTypeActionActionType_Read),
	}
}

func (s *AwsAuthorizationSystemTypeActionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsAuthorizationSystemTypeActionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsAuthorizationSystemTypeActionActionType(input string) (*AwsAuthorizationSystemTypeActionActionType, error) {
	vals := map[string]AwsAuthorizationSystemTypeActionActionType{
		"delete": AwsAuthorizationSystemTypeActionActionType_Delete,
		"read":   AwsAuthorizationSystemTypeActionActionType_Read,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsAuthorizationSystemTypeActionActionType(input)
	return &out, nil
}
