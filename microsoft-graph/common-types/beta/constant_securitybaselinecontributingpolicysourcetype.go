package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineContributingPolicySourceType string

const (
	SecurityBaselineContributingPolicySourceType_DeviceConfiguration SecurityBaselineContributingPolicySourceType = "deviceConfiguration"
	SecurityBaselineContributingPolicySourceType_DeviceIntent        SecurityBaselineContributingPolicySourceType = "deviceIntent"
)

func PossibleValuesForSecurityBaselineContributingPolicySourceType() []string {
	return []string{
		string(SecurityBaselineContributingPolicySourceType_DeviceConfiguration),
		string(SecurityBaselineContributingPolicySourceType_DeviceIntent),
	}
}

func (s *SecurityBaselineContributingPolicySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineContributingPolicySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineContributingPolicySourceType(input string) (*SecurityBaselineContributingPolicySourceType, error) {
	vals := map[string]SecurityBaselineContributingPolicySourceType{
		"deviceconfiguration": SecurityBaselineContributingPolicySourceType_DeviceConfiguration,
		"deviceintent":        SecurityBaselineContributingPolicySourceType_DeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineContributingPolicySourceType(input)
	return &out, nil
}
