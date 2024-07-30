package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventEventType string

const (
	CloudPCConnectivityEventEventType_DeviceHealthCheck   CloudPCConnectivityEventEventType = "deviceHealthCheck"
	CloudPCConnectivityEventEventType_Unknown             CloudPCConnectivityEventEventType = "unknown"
	CloudPCConnectivityEventEventType_UserConnection      CloudPCConnectivityEventEventType = "userConnection"
	CloudPCConnectivityEventEventType_UserTroubleshooting CloudPCConnectivityEventEventType = "userTroubleshooting"
)

func PossibleValuesForCloudPCConnectivityEventEventType() []string {
	return []string{
		string(CloudPCConnectivityEventEventType_DeviceHealthCheck),
		string(CloudPCConnectivityEventEventType_Unknown),
		string(CloudPCConnectivityEventEventType_UserConnection),
		string(CloudPCConnectivityEventEventType_UserTroubleshooting),
	}
}

func (s *CloudPCConnectivityEventEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityEventEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityEventEventType(input string) (*CloudPCConnectivityEventEventType, error) {
	vals := map[string]CloudPCConnectivityEventEventType{
		"devicehealthcheck":   CloudPCConnectivityEventEventType_DeviceHealthCheck,
		"unknown":             CloudPCConnectivityEventEventType_Unknown,
		"userconnection":      CloudPCConnectivityEventEventType_UserConnection,
		"usertroubleshooting": CloudPCConnectivityEventEventType_UserTroubleshooting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventEventType(input)
	return &out, nil
}
