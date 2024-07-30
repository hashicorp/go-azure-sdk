package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobExportJobStatus string

const (
	CloudPCExportJobExportJobStatus_Completed  CloudPCExportJobExportJobStatus = "completed"
	CloudPCExportJobExportJobStatus_Failed     CloudPCExportJobExportJobStatus = "failed"
	CloudPCExportJobExportJobStatus_InProgress CloudPCExportJobExportJobStatus = "inProgress"
	CloudPCExportJobExportJobStatus_NotStarted CloudPCExportJobExportJobStatus = "notStarted"
)

func PossibleValuesForCloudPCExportJobExportJobStatus() []string {
	return []string{
		string(CloudPCExportJobExportJobStatus_Completed),
		string(CloudPCExportJobExportJobStatus_Failed),
		string(CloudPCExportJobExportJobStatus_InProgress),
		string(CloudPCExportJobExportJobStatus_NotStarted),
	}
}

func (s *CloudPCExportJobExportJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCExportJobExportJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCExportJobExportJobStatus(input string) (*CloudPCExportJobExportJobStatus, error) {
	vals := map[string]CloudPCExportJobExportJobStatus{
		"completed":  CloudPCExportJobExportJobStatus_Completed,
		"failed":     CloudPCExportJobExportJobStatus_Failed,
		"inprogress": CloudPCExportJobExportJobStatus_InProgress,
		"notstarted": CloudPCExportJobExportJobStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExportJobExportJobStatus(input)
	return &out, nil
}
