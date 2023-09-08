package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineContributingPolicySourceType string

const (
	SecurityBaselineContributingPolicySourceTypedeviceConfiguration SecurityBaselineContributingPolicySourceType = "DeviceConfiguration"
	SecurityBaselineContributingPolicySourceTypedeviceIntent        SecurityBaselineContributingPolicySourceType = "DeviceIntent"
)

func PossibleValuesForSecurityBaselineContributingPolicySourceType() []string {
	return []string{
		string(SecurityBaselineContributingPolicySourceTypedeviceConfiguration),
		string(SecurityBaselineContributingPolicySourceTypedeviceIntent),
	}
}

func parseSecurityBaselineContributingPolicySourceType(input string) (*SecurityBaselineContributingPolicySourceType, error) {
	vals := map[string]SecurityBaselineContributingPolicySourceType{
		"deviceconfiguration": SecurityBaselineContributingPolicySourceTypedeviceConfiguration,
		"deviceintent":        SecurityBaselineContributingPolicySourceTypedeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineContributingPolicySourceType(input)
	return &out, nil
}
