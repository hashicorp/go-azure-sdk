package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAuthorizationSystemIdentityIdentityType string

const (
	PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Application     PermissionsDefinitionAuthorizationSystemIdentityIdentityType = "application"
	PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ManagedIdentity PermissionsDefinitionAuthorizationSystemIdentityIdentityType = "managedIdentity"
	PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Role            PermissionsDefinitionAuthorizationSystemIdentityIdentityType = "role"
	PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ServiceAccount  PermissionsDefinitionAuthorizationSystemIdentityIdentityType = "serviceAccount"
	PermissionsDefinitionAuthorizationSystemIdentityIdentityType_User            PermissionsDefinitionAuthorizationSystemIdentityIdentityType = "user"
)

func PossibleValuesForPermissionsDefinitionAuthorizationSystemIdentityIdentityType() []string {
	return []string{
		string(PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Application),
		string(PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ManagedIdentity),
		string(PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Role),
		string(PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ServiceAccount),
		string(PermissionsDefinitionAuthorizationSystemIdentityIdentityType_User),
	}
}

func (s *PermissionsDefinitionAuthorizationSystemIdentityIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsDefinitionAuthorizationSystemIdentityIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsDefinitionAuthorizationSystemIdentityIdentityType(input string) (*PermissionsDefinitionAuthorizationSystemIdentityIdentityType, error) {
	vals := map[string]PermissionsDefinitionAuthorizationSystemIdentityIdentityType{
		"application":     PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Application,
		"managedidentity": PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ManagedIdentity,
		"role":            PermissionsDefinitionAuthorizationSystemIdentityIdentityType_Role,
		"serviceaccount":  PermissionsDefinitionAuthorizationSystemIdentityIdentityType_ServiceAccount,
		"user":            PermissionsDefinitionAuthorizationSystemIdentityIdentityType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsDefinitionAuthorizationSystemIdentityIdentityType(input)
	return &out, nil
}
