package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType string

const (
	AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypespecificApps AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType = "SpecificApps"
	AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypeuserApproval AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType = "UserApproval"
)

func PossibleValuesForAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypespecificApps),
		string(AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypeuserApproval),
	}
}

func parseAndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType(input string) (*AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType{
		"specificapps": AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypespecificApps,
		"userapproval": AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessTypeuserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationCertificateAccessType(input)
	return &out, nil
}
