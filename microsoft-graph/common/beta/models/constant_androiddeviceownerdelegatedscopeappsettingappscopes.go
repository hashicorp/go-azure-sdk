package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes string

const (
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureNetworkActivityLog AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "CaptureNetworkActivityLog"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureSecurityLog        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "CaptureSecurityLog"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescertificateInstall        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "CertificateInstall"
	AndroidDeviceOwnerDelegatedScopeAppSettingAppScopesunspecified               AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes = "Unspecified"
)

func PossibleValuesForAndroidDeviceOwnerDelegatedScopeAppSettingAppScopes() []string {
	return []string{
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureNetworkActivityLog),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureSecurityLog),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescertificateInstall),
		string(AndroidDeviceOwnerDelegatedScopeAppSettingAppScopesunspecified),
	}
}

func parseAndroidDeviceOwnerDelegatedScopeAppSettingAppScopes(input string) (*AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes, error) {
	vals := map[string]AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes{
		"capturenetworkactivitylog": AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureNetworkActivityLog,
		"capturesecuritylog":        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescaptureSecurityLog,
		"certificateinstall":        AndroidDeviceOwnerDelegatedScopeAppSettingAppScopescertificateInstall,
		"unspecified":               AndroidDeviceOwnerDelegatedScopeAppSettingAppScopesunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes(input)
	return &out, nil
}
