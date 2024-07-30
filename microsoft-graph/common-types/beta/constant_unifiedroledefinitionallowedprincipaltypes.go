package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleDefinitionAllowedPrincipalTypes string

const (
	UnifiedRoleDefinitionAllowedPrincipalTypes_Group            UnifiedRoleDefinitionAllowedPrincipalTypes = "group"
	UnifiedRoleDefinitionAllowedPrincipalTypes_ServicePrincipal UnifiedRoleDefinitionAllowedPrincipalTypes = "servicePrincipal"
	UnifiedRoleDefinitionAllowedPrincipalTypes_User             UnifiedRoleDefinitionAllowedPrincipalTypes = "user"
)

func PossibleValuesForUnifiedRoleDefinitionAllowedPrincipalTypes() []string {
	return []string{
		string(UnifiedRoleDefinitionAllowedPrincipalTypes_Group),
		string(UnifiedRoleDefinitionAllowedPrincipalTypes_ServicePrincipal),
		string(UnifiedRoleDefinitionAllowedPrincipalTypes_User),
	}
}

func (s *UnifiedRoleDefinitionAllowedPrincipalTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleDefinitionAllowedPrincipalTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleDefinitionAllowedPrincipalTypes(input string) (*UnifiedRoleDefinitionAllowedPrincipalTypes, error) {
	vals := map[string]UnifiedRoleDefinitionAllowedPrincipalTypes{
		"group":            UnifiedRoleDefinitionAllowedPrincipalTypes_Group,
		"serviceprincipal": UnifiedRoleDefinitionAllowedPrincipalTypes_ServicePrincipal,
		"user":             UnifiedRoleDefinitionAllowedPrincipalTypes_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleDefinitionAllowedPrincipalTypes(input)
	return &out, nil
}
