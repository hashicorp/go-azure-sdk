package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationEapFastConfiguration string

const (
	MacOSEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential                         MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "NoProtectedAccessCredential"
	MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential                        MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredential"
	MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision            MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvision"
	MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously MacOSEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationEapFastConfiguration() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision),
		string(MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func parseMacOSEnterpriseWiFiConfigurationEapFastConfiguration(input string) (*MacOSEnterpriseWiFiConfigurationEapFastConfiguration, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         MacOSEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential,
		"useprotectedaccesscredential":                        MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": MacOSEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationEapFastConfiguration(input)
	return &out, nil
}
