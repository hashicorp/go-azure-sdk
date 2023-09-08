package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfileState string

const (
	EducationSynchronizationProfileStatedeleting           EducationSynchronizationProfileState = "Deleting"
	EducationSynchronizationProfileStatedeletionFailed     EducationSynchronizationProfileState = "DeletionFailed"
	EducationSynchronizationProfileStateprovisioned        EducationSynchronizationProfileState = "Provisioned"
	EducationSynchronizationProfileStateprovisioning       EducationSynchronizationProfileState = "Provisioning"
	EducationSynchronizationProfileStateprovisioningFailed EducationSynchronizationProfileState = "ProvisioningFailed"
)

func PossibleValuesForEducationSynchronizationProfileState() []string {
	return []string{
		string(EducationSynchronizationProfileStatedeleting),
		string(EducationSynchronizationProfileStatedeletionFailed),
		string(EducationSynchronizationProfileStateprovisioned),
		string(EducationSynchronizationProfileStateprovisioning),
		string(EducationSynchronizationProfileStateprovisioningFailed),
	}
}

func parseEducationSynchronizationProfileState(input string) (*EducationSynchronizationProfileState, error) {
	vals := map[string]EducationSynchronizationProfileState{
		"deleting":           EducationSynchronizationProfileStatedeleting,
		"deletionfailed":     EducationSynchronizationProfileStatedeletionFailed,
		"provisioned":        EducationSynchronizationProfileStateprovisioned,
		"provisioning":       EducationSynchronizationProfileStateprovisioning,
		"provisioningfailed": EducationSynchronizationProfileStateprovisioningFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationProfileState(input)
	return &out, nil
}
