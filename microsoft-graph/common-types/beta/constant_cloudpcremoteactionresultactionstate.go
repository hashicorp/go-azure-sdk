package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionResultActionState string

const (
	CloudPCRemoteActionResultActionState_Active       CloudPCRemoteActionResultActionState = "active"
	CloudPCRemoteActionResultActionState_Canceled     CloudPCRemoteActionResultActionState = "canceled"
	CloudPCRemoteActionResultActionState_Done         CloudPCRemoteActionResultActionState = "done"
	CloudPCRemoteActionResultActionState_Failed       CloudPCRemoteActionResultActionState = "failed"
	CloudPCRemoteActionResultActionState_None         CloudPCRemoteActionResultActionState = "none"
	CloudPCRemoteActionResultActionState_NotSupported CloudPCRemoteActionResultActionState = "notSupported"
	CloudPCRemoteActionResultActionState_Pending      CloudPCRemoteActionResultActionState = "pending"
)

func PossibleValuesForCloudPCRemoteActionResultActionState() []string {
	return []string{
		string(CloudPCRemoteActionResultActionState_Active),
		string(CloudPCRemoteActionResultActionState_Canceled),
		string(CloudPCRemoteActionResultActionState_Done),
		string(CloudPCRemoteActionResultActionState_Failed),
		string(CloudPCRemoteActionResultActionState_None),
		string(CloudPCRemoteActionResultActionState_NotSupported),
		string(CloudPCRemoteActionResultActionState_Pending),
	}
}

func (s *CloudPCRemoteActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRemoteActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRemoteActionResultActionState(input string) (*CloudPCRemoteActionResultActionState, error) {
	vals := map[string]CloudPCRemoteActionResultActionState{
		"active":       CloudPCRemoteActionResultActionState_Active,
		"canceled":     CloudPCRemoteActionResultActionState_Canceled,
		"done":         CloudPCRemoteActionResultActionState_Done,
		"failed":       CloudPCRemoteActionResultActionState_Failed,
		"none":         CloudPCRemoteActionResultActionState_None,
		"notsupported": CloudPCRemoteActionResultActionState_NotSupported,
		"pending":      CloudPCRemoteActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionResultActionState(input)
	return &out, nil
}
