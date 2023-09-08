package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource string

const (
	WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcefullDomainName    WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource = "FullDomainName"
	WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource = "NetBiosDomainName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcefullDomainName),
		string(WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName),
	}
}

func parseWindowsPhoneEASEmailProfileConfigurationUserDomainNameSource(input string) (*WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcefullDomainName,
		"netbiosdomainname": WindowsPhoneEASEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
