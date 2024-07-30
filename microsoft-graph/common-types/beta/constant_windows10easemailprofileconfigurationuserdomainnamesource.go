package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationUserDomainNameSource string

const (
	Windows10EasEmailProfileConfigurationUserDomainNameSource_FullDomainName    Windows10EasEmailProfileConfigurationUserDomainNameSource = "fullDomainName"
	Windows10EasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName Windows10EasEmailProfileConfigurationUserDomainNameSource = "netBiosDomainName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationUserDomainNameSource_FullDomainName),
		string(Windows10EasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName),
	}
}

func (s *Windows10EasEmailProfileConfigurationUserDomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EasEmailProfileConfigurationUserDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EasEmailProfileConfigurationUserDomainNameSource(input string) (*Windows10EasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    Windows10EasEmailProfileConfigurationUserDomainNameSource_FullDomainName,
		"netbiosdomainname": Windows10EasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
