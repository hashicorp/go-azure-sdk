package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes string

const (
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureNetworkActivityLog AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "captureNetworkActivityLog"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureSecurityLog        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "captureSecurityLog"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CertificateInstall        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "certificateInstall"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_Unspecified               AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "unspecified"
)

func PossibleValuesForAndroidDeviceOwnerDelegatedScopeAppSettingAppScopes() []string {
	return []string{
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureNetworkActivityLog),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureSecurityLog),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CertificateInstall),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_Unspecified),
	}
}

func (s *AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerDelegatedScopeAppSettingAppScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerDelegatedScopeAppSettingAppScopes(input string) (*AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes, error) {
	vals := map[string]AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes{
		"capturenetworkactivitylog": AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureNetworkActivityLog,
		"capturesecuritylog":        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CaptureSecurityLog,
		"certificateinstall":        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_CertificateInstall,
		"unspecified":               AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes(input)
	return &out, nil
}
