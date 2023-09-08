package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationEapFastConfiguration string

const (
	MacOSWiredNetworkConfigurationEapFastConfigurationnoProtectedAccessCredential                         MacOSWiredNetworkConfigurationEapFastConfiguration = "NoProtectedAccessCredential"
	MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredential                        MacOSWiredNetworkConfigurationEapFastConfiguration = "UseProtectedAccessCredential"
	MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision            MacOSWiredNetworkConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvision"
	MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously MacOSWiredNetworkConfigurationEapFastConfiguration = "UseProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForMacOSWiredNetworkConfigurationEapFastConfiguration() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationEapFastConfigurationnoProtectedAccessCredential),
		string(MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredential),
		string(MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision),
		string(MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func parseMacOSWiredNetworkConfigurationEapFastConfiguration(input string) (*MacOSWiredNetworkConfigurationEapFastConfiguration, error) {
	vals := map[string]MacOSWiredNetworkConfigurationEapFastConfiguration{
		"noprotectedaccesscredential":                         MacOSWiredNetworkConfigurationEapFastConfigurationnoProtectedAccessCredential,
		"useprotectedaccesscredential":                        MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": MacOSWiredNetworkConfigurationEapFastConfigurationuseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationEapFastConfiguration(input)
	return &out, nil
}
