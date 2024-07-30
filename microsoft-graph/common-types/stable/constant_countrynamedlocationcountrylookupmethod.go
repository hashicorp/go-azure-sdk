package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountryNamedLocationCountryLookupMethod string

const (
	CountryNamedLocationCountryLookupMethod_AuthenticatorAppGps CountryNamedLocationCountryLookupMethod = "authenticatorAppGps"
	CountryNamedLocationCountryLookupMethod_ClientIpAddress     CountryNamedLocationCountryLookupMethod = "clientIpAddress"
)

func PossibleValuesForCountryNamedLocationCountryLookupMethod() []string {
	return []string{
		string(CountryNamedLocationCountryLookupMethod_AuthenticatorAppGps),
		string(CountryNamedLocationCountryLookupMethod_ClientIpAddress),
	}
}

func (s *CountryNamedLocationCountryLookupMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCountryNamedLocationCountryLookupMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCountryNamedLocationCountryLookupMethod(input string) (*CountryNamedLocationCountryLookupMethod, error) {
	vals := map[string]CountryNamedLocationCountryLookupMethod{
		"authenticatorappgps": CountryNamedLocationCountryLookupMethod_AuthenticatorAppGps,
		"clientipaddress":     CountryNamedLocationCountryLookupMethod_ClientIpAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CountryNamedLocationCountryLookupMethod(input)
	return &out, nil
}
