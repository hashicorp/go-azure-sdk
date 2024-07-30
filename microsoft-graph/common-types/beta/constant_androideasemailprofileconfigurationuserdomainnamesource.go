package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationUserDomainNameSource string

const (
	AndroidEasEmailProfileConfigurationUserDomainNameSource_FullDomainName    AndroidEasEmailProfileConfigurationUserDomainNameSource = "fullDomainName"
	AndroidEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName AndroidEasEmailProfileConfigurationUserDomainNameSource = "netBiosDomainName"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationUserDomainNameSource_FullDomainName),
		string(AndroidEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName),
	}
}

func (s *AndroidEasEmailProfileConfigurationUserDomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationUserDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationUserDomainNameSource(input string) (*AndroidEasEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    AndroidEasEmailProfileConfigurationUserDomainNameSource_FullDomainName,
		"netbiosdomainname": AndroidEasEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
