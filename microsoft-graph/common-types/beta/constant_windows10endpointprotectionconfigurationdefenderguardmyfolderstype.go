package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType string

const (
	Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditDiskModification Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType = "auditDiskModification"
	Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditMode             Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_BlockDiskModification Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType = "blockDiskModification"
	Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_Enable                Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType = "enable"
	Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_UserDefined           Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType = "userDefined"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderGuardMyFoldersType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditDiskModification),
		string(Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_BlockDiskModification),
		string(Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_UserDefined),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderGuardMyFoldersType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderGuardMyFoldersType(input string) (*Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType{
		"auditdiskmodification": Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditDiskModification,
		"auditmode":             Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_AuditMode,
		"blockdiskmodification": Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_BlockDiskModification,
		"enable":                Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_Enable,
		"userdefined":           Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderGuardMyFoldersType(input)
	return &out, nil
}
