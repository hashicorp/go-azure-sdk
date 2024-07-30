package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAuthorizationSystemTypeActionActionType string

const (
	AzureAuthorizationSystemTypeActionActionType_Delete AzureAuthorizationSystemTypeActionActionType = "delete"
	AzureAuthorizationSystemTypeActionActionType_Read   AzureAuthorizationSystemTypeActionActionType = "read"
)

func PossibleValuesForAzureAuthorizationSystemTypeActionActionType() []string {
	return []string{
		string(AzureAuthorizationSystemTypeActionActionType_Delete),
		string(AzureAuthorizationSystemTypeActionActionType_Read),
	}
}

func (s *AzureAuthorizationSystemTypeActionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureAuthorizationSystemTypeActionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureAuthorizationSystemTypeActionActionType(input string) (*AzureAuthorizationSystemTypeActionActionType, error) {
	vals := map[string]AzureAuthorizationSystemTypeActionActionType{
		"delete": AzureAuthorizationSystemTypeActionActionType_Delete,
		"read":   AzureAuthorizationSystemTypeActionActionType_Read,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureAuthorizationSystemTypeActionActionType(input)
	return &out, nil
}
