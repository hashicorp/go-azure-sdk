package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledPermissionsRequestStatusDetail string

const (
	ScheduledPermissionsRequestStatusDetail_Approved  ScheduledPermissionsRequestStatusDetail = "approved"
	ScheduledPermissionsRequestStatusDetail_Canceled  ScheduledPermissionsRequestStatusDetail = "canceled"
	ScheduledPermissionsRequestStatusDetail_Completed ScheduledPermissionsRequestStatusDetail = "completed"
	ScheduledPermissionsRequestStatusDetail_Rejected  ScheduledPermissionsRequestStatusDetail = "rejected"
	ScheduledPermissionsRequestStatusDetail_Submitted ScheduledPermissionsRequestStatusDetail = "submitted"
)

func PossibleValuesForScheduledPermissionsRequestStatusDetail() []string {
	return []string{
		string(ScheduledPermissionsRequestStatusDetail_Approved),
		string(ScheduledPermissionsRequestStatusDetail_Canceled),
		string(ScheduledPermissionsRequestStatusDetail_Completed),
		string(ScheduledPermissionsRequestStatusDetail_Rejected),
		string(ScheduledPermissionsRequestStatusDetail_Submitted),
	}
}

func (s *ScheduledPermissionsRequestStatusDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduledPermissionsRequestStatusDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduledPermissionsRequestStatusDetail(input string) (*ScheduledPermissionsRequestStatusDetail, error) {
	vals := map[string]ScheduledPermissionsRequestStatusDetail{
		"approved":  ScheduledPermissionsRequestStatusDetail_Approved,
		"canceled":  ScheduledPermissionsRequestStatusDetail_Canceled,
		"completed": ScheduledPermissionsRequestStatusDetail_Completed,
		"rejected":  ScheduledPermissionsRequestStatusDetail_Rejected,
		"submitted": ScheduledPermissionsRequestStatusDetail_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduledPermissionsRequestStatusDetail(input)
	return &out, nil
}
