package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAccountManagerPolicyAccountDeletionPolicy string

const (
	SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThreshold                    SharedPCAccountManagerPolicyAccountDeletionPolicy = "diskSpaceThreshold"
	SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThresholdOrInactiveThreshold SharedPCAccountManagerPolicyAccountDeletionPolicy = "diskSpaceThresholdOrInactiveThreshold"
	SharedPCAccountManagerPolicyAccountDeletionPolicy_Immediate                             SharedPCAccountManagerPolicyAccountDeletionPolicy = "immediate"
)

func PossibleValuesForSharedPCAccountManagerPolicyAccountDeletionPolicy() []string {
	return []string{
		string(SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThreshold),
		string(SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThresholdOrInactiveThreshold),
		string(SharedPCAccountManagerPolicyAccountDeletionPolicy_Immediate),
	}
}

func (s *SharedPCAccountManagerPolicyAccountDeletionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedPCAccountManagerPolicyAccountDeletionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedPCAccountManagerPolicyAccountDeletionPolicy(input string) (*SharedPCAccountManagerPolicyAccountDeletionPolicy, error) {
	vals := map[string]SharedPCAccountManagerPolicyAccountDeletionPolicy{
		"diskspacethreshold":                    SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThreshold,
		"diskspacethresholdorinactivethreshold": SharedPCAccountManagerPolicyAccountDeletionPolicy_DiskSpaceThresholdOrInactiveThreshold,
		"immediate":                             SharedPCAccountManagerPolicyAccountDeletionPolicy_Immediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCAccountManagerPolicyAccountDeletionPolicy(input)
	return &out, nil
}
