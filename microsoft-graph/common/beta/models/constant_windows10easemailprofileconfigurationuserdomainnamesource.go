package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationUserDomainNameSource string

const (
	Windows10EasEmailProfileConfigurationUserDomainNameSourcefullDomainName    Windows10EasEmailProfileConfigurationUserDomainNameSource = "FullDomainName"
	Windows10EasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName Windows10EasEmailProfileConfigurationUserDomainNameSource = "NetBiosDomainName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationUserDomainNameSourcefullDomainName),
		string(Windows10EasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName),
	}
}

func parseWindows10EasEmailProfileConfigurationUserDomainNameSource(input string) (*Windows10EasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    Windows10EasEmailProfileConfigurationUserDomainNameSourcefullDomainName,
		"netbiosdomainname": Windows10EasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
