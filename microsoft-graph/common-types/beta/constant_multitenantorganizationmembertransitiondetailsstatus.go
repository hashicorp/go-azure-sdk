package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetailsStatus string

const (
	MultiTenantOrganizationMemberTransitionDetailsStatus_Failed     MultiTenantOrganizationMemberTransitionDetailsStatus = "failed"
	MultiTenantOrganizationMemberTransitionDetailsStatus_NotStarted MultiTenantOrganizationMemberTransitionDetailsStatus = "notStarted"
	MultiTenantOrganizationMemberTransitionDetailsStatus_Running    MultiTenantOrganizationMemberTransitionDetailsStatus = "running"
	MultiTenantOrganizationMemberTransitionDetailsStatus_Succeeded  MultiTenantOrganizationMemberTransitionDetailsStatus = "succeeded"
)

func PossibleValuesForMultiTenantOrganizationMemberTransitionDetailsStatus() []string {
	return []string{
		string(MultiTenantOrganizationMemberTransitionDetailsStatus_Failed),
		string(MultiTenantOrganizationMemberTransitionDetailsStatus_NotStarted),
		string(MultiTenantOrganizationMemberTransitionDetailsStatus_Running),
		string(MultiTenantOrganizationMemberTransitionDetailsStatus_Succeeded),
	}
}

func (s *MultiTenantOrganizationMemberTransitionDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberTransitionDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberTransitionDetailsStatus(input string) (*MultiTenantOrganizationMemberTransitionDetailsStatus, error) {
	vals := map[string]MultiTenantOrganizationMemberTransitionDetailsStatus{
		"failed":     MultiTenantOrganizationMemberTransitionDetailsStatus_Failed,
		"notstarted": MultiTenantOrganizationMemberTransitionDetailsStatus_NotStarted,
		"running":    MultiTenantOrganizationMemberTransitionDetailsStatus_Running,
		"succeeded":  MultiTenantOrganizationMemberTransitionDetailsStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberTransitionDetailsStatus(input)
	return &out, nil
}
