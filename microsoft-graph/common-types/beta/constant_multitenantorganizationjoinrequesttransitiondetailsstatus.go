package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetailsStatus string

const (
	MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Failed     MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "failed"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatus_NotStarted MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "notStarted"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Running    MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "running"
	MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Succeeded  MultiTenantOrganizationJoinRequestTransitionDetailsStatus = "succeeded"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestTransitionDetailsStatus() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Failed),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatus_NotStarted),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Running),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Succeeded),
	}
}

func (s *MultiTenantOrganizationJoinRequestTransitionDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationJoinRequestTransitionDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationJoinRequestTransitionDetailsStatus(input string) (*MultiTenantOrganizationJoinRequestTransitionDetailsStatus, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestTransitionDetailsStatus{
		"failed":     MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Failed,
		"notstarted": MultiTenantOrganizationJoinRequestTransitionDetailsStatus_NotStarted,
		"running":    MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Running,
		"succeeded":  MultiTenantOrganizationJoinRequestTransitionDetailsStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestTransitionDetailsStatus(input)
	return &out, nil
}
