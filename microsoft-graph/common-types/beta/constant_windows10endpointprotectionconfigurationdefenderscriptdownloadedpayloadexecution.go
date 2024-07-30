package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_AuditMode     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution = "auditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Enable        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution = "enable"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_NotConfigured Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution = "notConfigured"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_UserDefined   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Warn          Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution = "warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_AuditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Enable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_NotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Warn),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_AuditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_NotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_UserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecution(input)
	return &out, nil
}
