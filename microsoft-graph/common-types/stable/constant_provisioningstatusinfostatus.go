package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStatusInfoStatus string

const (
	ProvisioningStatusInfoStatus_Failure ProvisioningStatusInfoStatus = "failure"
	ProvisioningStatusInfoStatus_Skipped ProvisioningStatusInfoStatus = "skipped"
	ProvisioningStatusInfoStatus_Success ProvisioningStatusInfoStatus = "success"
	ProvisioningStatusInfoStatus_Warning ProvisioningStatusInfoStatus = "warning"
)

func PossibleValuesForProvisioningStatusInfoStatus() []string {
	return []string{
		string(ProvisioningStatusInfoStatus_Failure),
		string(ProvisioningStatusInfoStatus_Skipped),
		string(ProvisioningStatusInfoStatus_Success),
		string(ProvisioningStatusInfoStatus_Warning),
	}
}

func (s *ProvisioningStatusInfoStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStatusInfoStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStatusInfoStatus(input string) (*ProvisioningStatusInfoStatus, error) {
	vals := map[string]ProvisioningStatusInfoStatus{
		"failure": ProvisioningStatusInfoStatus_Failure,
		"skipped": ProvisioningStatusInfoStatus_Skipped,
		"success": ProvisioningStatusInfoStatus_Success,
		"warning": ProvisioningStatusInfoStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStatusInfoStatus(input)
	return &out, nil
}
