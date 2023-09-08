package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAccountManagerPolicyAccountDeletionPolicy string

const (
	SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThreshold                    SharedPCAccountManagerPolicyAccountDeletionPolicy = "DiskSpaceThreshold"
	SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThresholdOrInactiveThreshold SharedPCAccountManagerPolicyAccountDeletionPolicy = "DiskSpaceThresholdOrInactiveThreshold"
	SharedPCAccountManagerPolicyAccountDeletionPolicyimmediate                             SharedPCAccountManagerPolicyAccountDeletionPolicy = "Immediate"
)

func PossibleValuesForSharedPCAccountManagerPolicyAccountDeletionPolicy() []string {
	return []string{
		string(SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThreshold),
		string(SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThresholdOrInactiveThreshold),
		string(SharedPCAccountManagerPolicyAccountDeletionPolicyimmediate),
	}
}

func parseSharedPCAccountManagerPolicyAccountDeletionPolicy(input string) (*SharedPCAccountManagerPolicyAccountDeletionPolicy, error) {
	vals := map[string]SharedPCAccountManagerPolicyAccountDeletionPolicy{
		"diskspacethreshold":                    SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThreshold,
		"diskspacethresholdorinactivethreshold": SharedPCAccountManagerPolicyAccountDeletionPolicydiskSpaceThresholdOrInactiveThreshold,
		"immediate":                             SharedPCAccountManagerPolicyAccountDeletionPolicyimmediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCAccountManagerPolicyAccountDeletionPolicy(input)
	return &out, nil
}
