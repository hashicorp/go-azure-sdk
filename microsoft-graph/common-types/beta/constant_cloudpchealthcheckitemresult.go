package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCHealthCheckItemResult string

const (
	CloudPCHealthCheckItemResult_Failure CloudPCHealthCheckItemResult = "failure"
	CloudPCHealthCheckItemResult_Success CloudPCHealthCheckItemResult = "success"
	CloudPCHealthCheckItemResult_Unknown CloudPCHealthCheckItemResult = "unknown"
)

func PossibleValuesForCloudPCHealthCheckItemResult() []string {
	return []string{
		string(CloudPCHealthCheckItemResult_Failure),
		string(CloudPCHealthCheckItemResult_Success),
		string(CloudPCHealthCheckItemResult_Unknown),
	}
}

func (s *CloudPCHealthCheckItemResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCHealthCheckItemResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCHealthCheckItemResult(input string) (*CloudPCHealthCheckItemResult, error) {
	vals := map[string]CloudPCHealthCheckItemResult{
		"failure": CloudPCHealthCheckItemResult_Failure,
		"success": CloudPCHealthCheckItemResult_Success,
		"unknown": CloudPCHealthCheckItemResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCHealthCheckItemResult(input)
	return &out, nil
}
