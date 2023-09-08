package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationEapFastConfiguration string

const (
	IosEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential                         IosEnterpriseWiFiConfigurationEapFastConfiguration = "NoProtectedAccessCredential"
	IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential                        IosEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredential"
	IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision            IosEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvision"
	IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously IosEnterpriseWiFiConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationEapFastConfiguration() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential),
		string(IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential),
		string(IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision),
		string(IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func parseIosEnterpriseWiFiConfigurationEapFastConfiguration(input string) (*IosEnterpriseWiFiConfigurationEapFastConfiguration, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         IosEnterpriseWiFiConfigurationEapFastConfigurationnoProtectedAccessCredential,
		"useprotectedaccesscredential":                        IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": IosEnterpriseWiFiConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationEapFastConfiguration(input)
	return &out, nil
}
