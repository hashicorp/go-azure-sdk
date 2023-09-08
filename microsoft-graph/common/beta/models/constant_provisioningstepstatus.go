package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStepStatus string

const (
	ProvisioningStepStatusfailure ProvisioningStepStatus = "Failure"
	ProvisioningStepStatusskipped ProvisioningStepStatus = "Skipped"
	ProvisioningStepStatussuccess ProvisioningStepStatus = "Success"
	ProvisioningStepStatuswarning ProvisioningStepStatus = "Warning"
)

func PossibleValuesForProvisioningStepStatus() []string {
	return []string{
		string(ProvisioningStepStatusfailure),
		string(ProvisioningStepStatusskipped),
		string(ProvisioningStepStatussuccess),
		string(ProvisioningStepStatuswarning),
	}
}

func parseProvisioningStepStatus(input string) (*ProvisioningStepStatus, error) {
	vals := map[string]ProvisioningStepStatus{
		"failure": ProvisioningStepStatusfailure,
		"skipped": ProvisioningStepStatusskipped,
		"success": ProvisioningStepStatussuccess,
		"warning": ProvisioningStepStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStepStatus(input)
	return &out, nil
}
