package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUserDomainNameSource string

const (
	EasEmailProfileConfigurationBaseUserDomainNameSource_FullDomainName    EasEmailProfileConfigurationBaseUserDomainNameSource = "fullDomainName"
	EasEmailProfileConfigurationBaseUserDomainNameSource_NetBiosDomainName EasEmailProfileConfigurationBaseUserDomainNameSource = "netBiosDomainName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUserDomainNameSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUserDomainNameSource_FullDomainName),
		string(EasEmailProfileConfigurationBaseUserDomainNameSource_NetBiosDomainName),
	}
}

func (s *EasEmailProfileConfigurationBaseUserDomainNameSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEasEmailProfileConfigurationBaseUserDomainNameSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEasEmailProfileConfigurationBaseUserDomainNameSource(input string) (*EasEmailProfileConfigurationBaseUserDomainNameSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUserDomainNameSource{
		"fulldomainname":    EasEmailProfileConfigurationBaseUserDomainNameSource_FullDomainName,
		"netbiosdomainname": EasEmailProfileConfigurationBaseUserDomainNameSource_NetBiosDomainName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUserDomainNameSource(input)
	return &out, nil
}
