package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource string

const (
	WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_FullDomainName    WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource = "fullDomainName"
	WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource = "netBiosDomainName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationUserDomainNameSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_FullDomainName),
		string(WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName),
	}
}

func (s *WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhoneEASEmailProfileConfigurationUserDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhoneEASEmailProfileConfigurationUserDomainNameSource(input string) (*WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource{
		"fulldomainname":    WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_FullDomainName,
		"netbiosdomainname": WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationUserDomainNameSource(input)
	return &out, nil
}
