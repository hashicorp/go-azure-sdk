package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsRequestChangeStatusDetail string

const (
	PermissionsRequestChangeStatusDetail_Approved  PermissionsRequestChangeStatusDetail = "approved"
	PermissionsRequestChangeStatusDetail_Canceled  PermissionsRequestChangeStatusDetail = "canceled"
	PermissionsRequestChangeStatusDetail_Completed PermissionsRequestChangeStatusDetail = "completed"
	PermissionsRequestChangeStatusDetail_Rejected  PermissionsRequestChangeStatusDetail = "rejected"
	PermissionsRequestChangeStatusDetail_Submitted PermissionsRequestChangeStatusDetail = "submitted"
)

func PossibleValuesForPermissionsRequestChangeStatusDetail() []string {
	return []string{
		string(PermissionsRequestChangeStatusDetail_Approved),
		string(PermissionsRequestChangeStatusDetail_Canceled),
		string(PermissionsRequestChangeStatusDetail_Completed),
		string(PermissionsRequestChangeStatusDetail_Rejected),
		string(PermissionsRequestChangeStatusDetail_Submitted),
	}
}

func (s *PermissionsRequestChangeStatusDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsRequestChangeStatusDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsRequestChangeStatusDetail(input string) (*PermissionsRequestChangeStatusDetail, error) {
	vals := map[string]PermissionsRequestChangeStatusDetail{
		"approved":  PermissionsRequestChangeStatusDetail_Approved,
		"canceled":  PermissionsRequestChangeStatusDetail_Canceled,
		"completed": PermissionsRequestChangeStatusDetail_Completed,
		"rejected":  PermissionsRequestChangeStatusDetail_Rejected,
		"submitted": PermissionsRequestChangeStatusDetail_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsRequestChangeStatusDetail(input)
	return &out, nil
}
