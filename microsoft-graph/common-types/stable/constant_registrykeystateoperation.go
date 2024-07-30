package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateOperation string

const (
	RegistryKeyStateOperation_Create  RegistryKeyStateOperation = "create"
	RegistryKeyStateOperation_Delete  RegistryKeyStateOperation = "delete"
	RegistryKeyStateOperation_Modify  RegistryKeyStateOperation = "modify"
	RegistryKeyStateOperation_Unknown RegistryKeyStateOperation = "unknown"
)

func PossibleValuesForRegistryKeyStateOperation() []string {
	return []string{
		string(RegistryKeyStateOperation_Create),
		string(RegistryKeyStateOperation_Delete),
		string(RegistryKeyStateOperation_Modify),
		string(RegistryKeyStateOperation_Unknown),
	}
}

func (s *RegistryKeyStateOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryKeyStateOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryKeyStateOperation(input string) (*RegistryKeyStateOperation, error) {
	vals := map[string]RegistryKeyStateOperation{
		"create":  RegistryKeyStateOperation_Create,
		"delete":  RegistryKeyStateOperation_Delete,
		"modify":  RegistryKeyStateOperation_Modify,
		"unknown": RegistryKeyStateOperation_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateOperation(input)
	return &out, nil
}
