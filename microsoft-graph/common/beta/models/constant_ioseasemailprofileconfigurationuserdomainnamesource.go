package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUserDomainNameSource string

const (
	IosEasEmailProfileConfigurationUserDomainNameSourcefullDomainName    IosEasEmailProfileConfigurationUserDomainNameSource = "FullDomainName"
	IosEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName IosEasEmailProfileConfigurationUserDomainNameSource = "NetBiosDomainName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUserDomainNameSourcefullDomainName),
		string(IosEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName),
	}
}

func parseIosEasEmailProfileConfigurationUserDomainNameSource(input string) (*IosEasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    IosEasEmailProfileConfigurationUserDomainNameSourcefullDomainName,
		"netbiosdomainname": IosEasEmailProfileConfigurationUserDomainNameSourcenetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
