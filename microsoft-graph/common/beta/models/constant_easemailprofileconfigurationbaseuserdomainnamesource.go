package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUserDomainNameSource string

const (
	EasEmailProfileConfigurationBaseUserDomainNameSourcefullDomainName    EasEmailProfileConfigurationBaseUserDomainNameSource = "FullDomainName"
	EasEmailProfileConfigurationBaseUserDomainNameSourcenetBiosDomainName EasEmailProfileConfigurationBaseUserDomainNameSource = "NetBiosDomainName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUserDomainNameSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUserDomainNameSourcefullDomainName),
		string(EasEmailProfileConfigurationBaseUserDomainNameSourcenetBiosDomainName),
	}
}

func parseEasEmailProfileConfigurationBaseUserDomainNameSource(input string) (*EasEmailProfileConfigurationBaseUserDomainNameSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUserDomainNameSource{
		"fulldomainname":    EasEmailProfileConfigurationBaseUserDomainNameSourcefullDomainName,
		"netbiosdomainname": EasEmailProfileConfigurationBaseUserDomainNameSourcenetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUserDomainNameSource(input)
	return &out, nil
}
