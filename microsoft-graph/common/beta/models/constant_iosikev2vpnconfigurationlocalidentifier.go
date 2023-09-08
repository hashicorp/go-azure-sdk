package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationLocalIdentifier string

const (
	IosikEv2VpnConfigurationLocalIdentifierclientCertificateSubjectName IosikEv2VpnConfigurationLocalIdentifier = "ClientCertificateSubjectName"
	IosikEv2VpnConfigurationLocalIdentifierdeviceFQDN                   IosikEv2VpnConfigurationLocalIdentifier = "DeviceFQDN"
	IosikEv2VpnConfigurationLocalIdentifierempty                        IosikEv2VpnConfigurationLocalIdentifier = "Empty"
)

func PossibleValuesForIosikEv2VpnConfigurationLocalIdentifier() []string {
	return []string{
		string(IosikEv2VpnConfigurationLocalIdentifierclientCertificateSubjectName),
		string(IosikEv2VpnConfigurationLocalIdentifierdeviceFQDN),
		string(IosikEv2VpnConfigurationLocalIdentifierempty),
	}
}

func parseIosikEv2VpnConfigurationLocalIdentifier(input string) (*IosikEv2VpnConfigurationLocalIdentifier, error) {
	vals := map[string]IosikEv2VpnConfigurationLocalIdentifier{
		"clientcertificatesubjectname": IosikEv2VpnConfigurationLocalIdentifierclientCertificateSubjectName,
		"devicefqdn":                   IosikEv2VpnConfigurationLocalIdentifierdeviceFQDN,
		"empty":                        IosikEv2VpnConfigurationLocalIdentifierempty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationLocalIdentifier(input)
	return &out, nil
}
