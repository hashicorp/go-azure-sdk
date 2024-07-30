package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkMobileAppConfigurationProfileApplicability string

const (
	AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidDeviceOwner AndroidForWorkMobileAppConfigurationProfileApplicability = "androidDeviceOwner"
	AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidWorkProfile AndroidForWorkMobileAppConfigurationProfileApplicability = "androidWorkProfile"
	AndroidForWorkMobileAppConfigurationProfileApplicability_Default            AndroidForWorkMobileAppConfigurationProfileApplicability = "default"
)

func PossibleValuesForAndroidForWorkMobileAppConfigurationProfileApplicability() []string {
	return []string{
		string(AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidDeviceOwner),
		string(AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidWorkProfile),
		string(AndroidForWorkMobileAppConfigurationProfileApplicability_Default),
	}
}

func (s *AndroidForWorkMobileAppConfigurationProfileApplicability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkMobileAppConfigurationProfileApplicability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkMobileAppConfigurationProfileApplicability(input string) (*AndroidForWorkMobileAppConfigurationProfileApplicability, error) {
	vals := map[string]AndroidForWorkMobileAppConfigurationProfileApplicability{
		"androiddeviceowner": AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidDeviceOwner,
		"androidworkprofile": AndroidForWorkMobileAppConfigurationProfileApplicability_AndroidWorkProfile,
		"default":            AndroidForWorkMobileAppConfigurationProfileApplicability_Default,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkMobileAppConfigurationProfileApplicability(input)
	return &out, nil
}
