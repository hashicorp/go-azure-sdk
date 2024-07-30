package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOrganizationalScopeScopeType string

const (
	SecurityOrganizationalScopeScopeType_DeviceGroup SecurityOrganizationalScopeScopeType = "deviceGroup"
)

func PossibleValuesForSecurityOrganizationalScopeScopeType() []string {
	return []string{
		string(SecurityOrganizationalScopeScopeType_DeviceGroup),
	}
}

func (s *SecurityOrganizationalScopeScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityOrganizationalScopeScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityOrganizationalScopeScopeType(input string) (*SecurityOrganizationalScopeScopeType, error) {
	vals := map[string]SecurityOrganizationalScopeScopeType{
		"devicegroup": SecurityOrganizationalScopeScopeType_DeviceGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOrganizationalScopeScopeType(input)
	return &out, nil
}
