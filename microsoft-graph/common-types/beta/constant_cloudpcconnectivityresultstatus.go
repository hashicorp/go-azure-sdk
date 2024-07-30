package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityResultStatus string

const (
	CloudPCConnectivityResultStatus_Available            CloudPCConnectivityResultStatus = "available"
	CloudPCConnectivityResultStatus_AvailableWithWarning CloudPCConnectivityResultStatus = "availableWithWarning"
	CloudPCConnectivityResultStatus_Unavailable          CloudPCConnectivityResultStatus = "unavailable"
	CloudPCConnectivityResultStatus_Unknown              CloudPCConnectivityResultStatus = "unknown"
)

func PossibleValuesForCloudPCConnectivityResultStatus() []string {
	return []string{
		string(CloudPCConnectivityResultStatus_Available),
		string(CloudPCConnectivityResultStatus_AvailableWithWarning),
		string(CloudPCConnectivityResultStatus_Unavailable),
		string(CloudPCConnectivityResultStatus_Unknown),
	}
}

func (s *CloudPCConnectivityResultStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityResultStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityResultStatus(input string) (*CloudPCConnectivityResultStatus, error) {
	vals := map[string]CloudPCConnectivityResultStatus{
		"available":            CloudPCConnectivityResultStatus_Available,
		"availablewithwarning": CloudPCConnectivityResultStatus_AvailableWithWarning,
		"unavailable":          CloudPCConnectivityResultStatus_Unavailable,
		"unknown":              CloudPCConnectivityResultStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityResultStatus(input)
	return &out, nil
}
