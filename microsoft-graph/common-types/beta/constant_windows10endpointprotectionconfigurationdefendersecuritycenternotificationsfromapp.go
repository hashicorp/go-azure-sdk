package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp string

const (
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockAllNotifications         Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "blockAllNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockNoncriticalNotifications Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "blockNoncriticalNotifications"
	Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_NotConfigured                 Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockAllNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockNoncriticalNotifications),
		string(Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp(input string) (*Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp{
		"blockallnotifications":         Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockAllNotifications,
		"blocknoncriticalnotifications": Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_BlockNoncriticalNotifications,
		"notconfigured":                 Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderSecurityCenterNotificationsFromApp(input)
	return &out, nil
}
