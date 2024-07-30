package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventEventResult string

const (
	CloudPCConnectivityEventEventResult_Failure CloudPCConnectivityEventEventResult = "failure"
	CloudPCConnectivityEventEventResult_Success CloudPCConnectivityEventEventResult = "success"
	CloudPCConnectivityEventEventResult_Unknown CloudPCConnectivityEventEventResult = "unknown"
)

func PossibleValuesForCloudPCConnectivityEventEventResult() []string {
	return []string{
		string(CloudPCConnectivityEventEventResult_Failure),
		string(CloudPCConnectivityEventEventResult_Success),
		string(CloudPCConnectivityEventEventResult_Unknown),
	}
}

func (s *CloudPCConnectivityEventEventResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityEventEventResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityEventEventResult(input string) (*CloudPCConnectivityEventEventResult, error) {
	vals := map[string]CloudPCConnectivityEventEventResult{
		"failure": CloudPCConnectivityEventEventResult_Failure,
		"success": CloudPCConnectivityEventEventResult_Success,
		"unknown": CloudPCConnectivityEventEventResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventEventResult(input)
	return &out, nil
}
