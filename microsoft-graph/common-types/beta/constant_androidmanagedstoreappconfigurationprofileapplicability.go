package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationProfileApplicability string

const (
	AndroidManagedStoreAppConfigurationProfileApplicability_AndroidDeviceOwner AndroidManagedStoreAppConfigurationProfileApplicability = "androidDeviceOwner"
	AndroidManagedStoreAppConfigurationProfileApplicability_AndroidWorkProfile AndroidManagedStoreAppConfigurationProfileApplicability = "androidWorkProfile"
	AndroidManagedStoreAppConfigurationProfileApplicability_Default            AndroidManagedStoreAppConfigurationProfileApplicability = "default"
)

func PossibleValuesForAndroidManagedStoreAppConfigurationProfileApplicability() []string {
	return []string{
		string(AndroidManagedStoreAppConfigurationProfileApplicability_AndroidDeviceOwner),
		string(AndroidManagedStoreAppConfigurationProfileApplicability_AndroidWorkProfile),
		string(AndroidManagedStoreAppConfigurationProfileApplicability_Default),
	}
}

func (s *AndroidManagedStoreAppConfigurationProfileApplicability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAppConfigurationProfileApplicability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAppConfigurationProfileApplicability(input string) (*AndroidManagedStoreAppConfigurationProfileApplicability, error) {
	vals := map[string]AndroidManagedStoreAppConfigurationProfileApplicability{
		"androiddeviceowner": AndroidManagedStoreAppConfigurationProfileApplicability_AndroidDeviceOwner,
		"androidworkprofile": AndroidManagedStoreAppConfigurationProfileApplicability_AndroidWorkProfile,
		"default":            AndroidManagedStoreAppConfigurationProfileApplicability_Default,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppConfigurationProfileApplicability(input)
	return &out, nil
}
