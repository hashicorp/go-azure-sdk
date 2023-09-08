package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationClientAuthenticationType string

const (
	IosikEv2VpnConfigurationClientAuthenticationTypedeviceAuthentication IosikEv2VpnConfigurationClientAuthenticationType = "DeviceAuthentication"
	IosikEv2VpnConfigurationClientAuthenticationTypeuserAuthentication   IosikEv2VpnConfigurationClientAuthenticationType = "UserAuthentication"
)

func PossibleValuesForIosikEv2VpnConfigurationClientAuthenticationType() []string {
	return []string{
		string(IosikEv2VpnConfigurationClientAuthenticationTypedeviceAuthentication),
		string(IosikEv2VpnConfigurationClientAuthenticationTypeuserAuthentication),
	}
}

func parseIosikEv2VpnConfigurationClientAuthenticationType(input string) (*IosikEv2VpnConfigurationClientAuthenticationType, error) {
	vals := map[string]IosikEv2VpnConfigurationClientAuthenticationType{
		"deviceauthentication": IosikEv2VpnConfigurationClientAuthenticationTypedeviceAuthentication,
		"userauthentication":   IosikEv2VpnConfigurationClientAuthenticationTypeuserAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationClientAuthenticationType(input)
	return &out, nil
}
