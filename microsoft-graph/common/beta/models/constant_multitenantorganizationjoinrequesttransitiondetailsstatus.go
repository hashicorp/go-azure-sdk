package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetailsStatus string

const (
	MultiTenantOrganizationJoinRequestTransitionDetailsStatusfailed     MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "Failed"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatusnotStarted MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "NotStarted"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatusrunning    MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "Running"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatussucceeded  MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "Succeeded"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestTransitionDetailsStatus() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatusfailed),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatusnotStarted),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatusrunning),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatussucceeded),
	}
}

func parseMultiTenantOrganizationJoinRequestTransitionDetailsStatus(input string) (*MultiTenantOrganizationJoinRequestTransitionDetailsStatus, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestTransitionDetailsStatus{
		"failed":     MultiTenantOrganizationJoinRequestTransitionDetailsStatusfailed,
		"notstarted": MultiTenantOrganizationJoinRequestTransitionDetailsStatusnotStarted,
		"running":    MultiTenantOrganizationJoinRequestTransitionDetailsStatusrunning,
		"succeeded":  MultiTenantOrganizationJoinRequestTransitionDetailsStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestTransitionDetailsStatus(input)
	return &out, nil
}
