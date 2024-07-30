package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationMobileDeviceManagementAuthority string

const (
	OrganizationMobileDeviceManagementAuthority_Intune    OrganizationMobileDeviceManagementAuthority = "intune"
	OrganizationMobileDeviceManagementAuthority_Office365 OrganizationMobileDeviceManagementAuthority = "office365"
	OrganizationMobileDeviceManagementAuthority_Sccm      OrganizationMobileDeviceManagementAuthority = "sccm"
	OrganizationMobileDeviceManagementAuthority_Unknown   OrganizationMobileDeviceManagementAuthority = "unknown"
)

func PossibleValuesForOrganizationMobileDeviceManagementAuthority() []string {
	return []string{
		string(OrganizationMobileDeviceManagementAuthority_Intune),
		string(OrganizationMobileDeviceManagementAuthority_Office365),
		string(OrganizationMobileDeviceManagementAuthority_Sccm),
		string(OrganizationMobileDeviceManagementAuthority_Unknown),
	}
}

func (s *OrganizationMobileDeviceManagementAuthority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrganizationMobileDeviceManagementAuthority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrganizationMobileDeviceManagementAuthority(input string) (*OrganizationMobileDeviceManagementAuthority, error) {
	vals := map[string]OrganizationMobileDeviceManagementAuthority{
		"intune":    OrganizationMobileDeviceManagementAuthority_Intune,
		"office365": OrganizationMobileDeviceManagementAuthority_Office365,
		"sccm":      OrganizationMobileDeviceManagementAuthority_Sccm,
		"unknown":   OrganizationMobileDeviceManagementAuthority_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationMobileDeviceManagementAuthority(input)
	return &out, nil
}
