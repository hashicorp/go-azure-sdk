package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsIdentityAccessManagementKeyUsageFindingStatus string

const (
	AwsIdentityAccessManagementKeyUsageFindingStatus_Active   AwsIdentityAccessManagementKeyUsageFindingStatus = "active"
	AwsIdentityAccessManagementKeyUsageFindingStatus_Disabled AwsIdentityAccessManagementKeyUsageFindingStatus = "disabled"
	AwsIdentityAccessManagementKeyUsageFindingStatus_Inactive AwsIdentityAccessManagementKeyUsageFindingStatus = "inactive"
)

func PossibleValuesForAwsIdentityAccessManagementKeyUsageFindingStatus() []string {
	return []string{
		string(AwsIdentityAccessManagementKeyUsageFindingStatus_Active),
		string(AwsIdentityAccessManagementKeyUsageFindingStatus_Disabled),
		string(AwsIdentityAccessManagementKeyUsageFindingStatus_Inactive),
	}
}

func (s *AwsIdentityAccessManagementKeyUsageFindingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsIdentityAccessManagementKeyUsageFindingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsIdentityAccessManagementKeyUsageFindingStatus(input string) (*AwsIdentityAccessManagementKeyUsageFindingStatus, error) {
	vals := map[string]AwsIdentityAccessManagementKeyUsageFindingStatus{
		"active":   AwsIdentityAccessManagementKeyUsageFindingStatus_Active,
		"disabled": AwsIdentityAccessManagementKeyUsageFindingStatus_Disabled,
		"inactive": AwsIdentityAccessManagementKeyUsageFindingStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsIdentityAccessManagementKeyUsageFindingStatus(input)
	return &out, nil
}
