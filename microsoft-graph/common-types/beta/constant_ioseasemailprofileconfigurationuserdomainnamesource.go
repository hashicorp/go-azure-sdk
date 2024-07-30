package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUserDomainNameSource string

const (
	IosEasEmailProfileConfigurationUserDomainNameSource_FullDomainName    IosEasEmailProfileConfigurationUserDomainNameSource = "fullDomainName"
	IosEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName IosEasEmailProfileConfigurationUserDomainNameSource = "netBiosDomainName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUserDomainNameSource_FullDomainName),
		string(IosEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName),
	}
}

func (s *IosEasEmailProfileConfigurationUserDomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationUserDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationUserDomainNameSource(input string) (*IosEasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    IosEasEmailProfileConfigurationUserDomainNameSource_FullDomainName,
		"netbiosdomainname": IosEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
