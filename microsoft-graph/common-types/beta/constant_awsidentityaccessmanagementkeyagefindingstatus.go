package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsIdentityAccessManagementKeyAgeFindingStatus string

const (
	AwsIdentityAccessManagementKeyAgeFindingStatus_Active   AwsIdentityAccessManagementKeyAgeFindingStatus = "active"
	AwsIdentityAccessManagementKeyAgeFindingStatus_Disabled AwsIdentityAccessManagementKeyAgeFindingStatus = "disabled"
	AwsIdentityAccessManagementKeyAgeFindingStatus_Inactive AwsIdentityAccessManagementKeyAgeFindingStatus = "inactive"
)

func PossibleValuesForAwsIdentityAccessManagementKeyAgeFindingStatus() []string {
	return []string{
		string(AwsIdentityAccessManagementKeyAgeFindingStatus_Active),
		string(AwsIdentityAccessManagementKeyAgeFindingStatus_Disabled),
		string(AwsIdentityAccessManagementKeyAgeFindingStatus_Inactive),
	}
}

func (s *AwsIdentityAccessManagementKeyAgeFindingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsIdentityAccessManagementKeyAgeFindingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsIdentityAccessManagementKeyAgeFindingStatus(input string) (*AwsIdentityAccessManagementKeyAgeFindingStatus, error) {
	vals := map[string]AwsIdentityAccessManagementKeyAgeFindingStatus{
		"active":   AwsIdentityAccessManagementKeyAgeFindingStatus_Active,
		"disabled": AwsIdentityAccessManagementKeyAgeFindingStatus_Disabled,
		"inactive": AwsIdentityAccessManagementKeyAgeFindingStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsIdentityAccessManagementKeyAgeFindingStatus(input)
	return &out, nil
}
