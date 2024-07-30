package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_AuditMode   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Block       Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "block"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Disable     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "disable"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_UserDefined Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Warn        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Block),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Disable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_AuditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Block,
		"disable":     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Disable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_UserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType(input)
	return &out, nil
}
