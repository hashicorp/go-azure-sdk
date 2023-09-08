package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStatusInfoStatus string

const (
	ProvisioningStatusInfoStatusfailure ProvisioningStatusInfoStatus = "Failure"
	ProvisioningStatusInfoStatusskipped ProvisioningStatusInfoStatus = "Skipped"
	ProvisioningStatusInfoStatussuccess ProvisioningStatusInfoStatus = "Success"
	ProvisioningStatusInfoStatuswarning ProvisioningStatusInfoStatus = "Warning"
)

func PossibleValuesForProvisioningStatusInfoStatus() []string {
	return []string{
		string(ProvisioningStatusInfoStatusfailure),
		string(ProvisioningStatusInfoStatusskipped),
		string(ProvisioningStatusInfoStatussuccess),
		string(ProvisioningStatusInfoStatuswarning),
	}
}

func parseProvisioningStatusInfoStatus(input string) (*ProvisioningStatusInfoStatus, error) {
	vals := map[string]ProvisioningStatusInfoStatus{
		"failure": ProvisioningStatusInfoStatusfailure,
		"skipped": ProvisioningStatusInfoStatusskipped,
		"success": ProvisioningStatusInfoStatussuccess,
		"warning": ProvisioningStatusInfoStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStatusInfoStatus(input)
	return &out, nil
}
