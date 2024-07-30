package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureRoleDefinitionAzureRoleDefinitionType string

const (
	AzureRoleDefinitionAzureRoleDefinitionType_Custom AzureRoleDefinitionAzureRoleDefinitionType = "custom"
	AzureRoleDefinitionAzureRoleDefinitionType_System AzureRoleDefinitionAzureRoleDefinitionType = "system"
)

func PossibleValuesForAzureRoleDefinitionAzureRoleDefinitionType() []string {
	return []string{
		string(AzureRoleDefinitionAzureRoleDefinitionType_Custom),
		string(AzureRoleDefinitionAzureRoleDefinitionType_System),
	}
}

func (s *AzureRoleDefinitionAzureRoleDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureRoleDefinitionAzureRoleDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureRoleDefinitionAzureRoleDefinitionType(input string) (*AzureRoleDefinitionAzureRoleDefinitionType, error) {
	vals := map[string]AzureRoleDefinitionAzureRoleDefinitionType{
		"custom": AzureRoleDefinitionAzureRoleDefinitionType_Custom,
		"system": AzureRoleDefinitionAzureRoleDefinitionType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureRoleDefinitionAzureRoleDefinitionType(input)
	return &out, nil
}
