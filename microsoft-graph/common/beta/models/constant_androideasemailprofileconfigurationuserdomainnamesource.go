package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationUserDomainNameSource string

const (
	AndroidEasEmailProfileConfigurationUserDomainNameSourcefullDomainName    AndroidEasEmailProfileConfigurationUserDomainNameSource = "FullDomainName"
	AndroidEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName AndroidEasEmailProfileConfigurationUserDomainNameSource = "NetBiosDomainName"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationUserDomainNameSourcefullDomainName),
		string(AndroidEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName),
	}
}

func parseAndroidEasEmailProfileConfigurationUserDomainNameSource(input string) (*AndroidEasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    AndroidEasEmailProfileConfigurationUserDomainNameSourcefullDomainName,
		"netbiosdomainname": AndroidEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
