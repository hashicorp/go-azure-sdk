package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenLastSyncStatus string

const (
	VppTokenLastSyncStatus_Completed  VppTokenLastSyncStatus = "completed"
	VppTokenLastSyncStatus_Failed     VppTokenLastSyncStatus = "failed"
	VppTokenLastSyncStatus_InProgress VppTokenLastSyncStatus = "inProgress"
	VppTokenLastSyncStatus_None       VppTokenLastSyncStatus = "none"
)

func PossibleValuesForVppTokenLastSyncStatus() []string {
	return []string{
		string(VppTokenLastSyncStatus_Completed),
		string(VppTokenLastSyncStatus_Failed),
		string(VppTokenLastSyncStatus_InProgress),
		string(VppTokenLastSyncStatus_None),
	}
}

func (s *VppTokenLastSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenLastSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenLastSyncStatus(input string) (*VppTokenLastSyncStatus, error) {
	vals := map[string]VppTokenLastSyncStatus{
		"completed":  VppTokenLastSyncStatus_Completed,
		"failed":     VppTokenLastSyncStatus_Failed,
		"inprogress": VppTokenLastSyncStatus_InProgress,
		"none":       VppTokenLastSyncStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenLastSyncStatus(input)
	return &out, nil
}
