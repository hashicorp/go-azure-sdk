package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventEventType string

const (
	CloudPCConnectivityEventEventTypedeviceHealthCheck   CloudPCConnectivityEventEventType = "DeviceHealthCheck"
	CloudPCConnectivityEventEventTypeunknown             CloudPCConnectivityEventEventType = "Unknown"
	CloudPCConnectivityEventEventTypeuserConnection      CloudPCConnectivityEventEventType = "UserConnection"
	CloudPCConnectivityEventEventTypeuserTroubleshooting CloudPCConnectivityEventEventType = "UserTroubleshooting"
)

func PossibleValuesForCloudPCConnectivityEventEventType() []string {
	return []string{
		string(CloudPCConnectivityEventEventTypedeviceHealthCheck),
		string(CloudPCConnectivityEventEventTypeunknown),
		string(CloudPCConnectivityEventEventTypeuserConnection),
		string(CloudPCConnectivityEventEventTypeuserTroubleshooting),
	}
}

func parseCloudPCConnectivityEventEventType(input string) (*CloudPCConnectivityEventEventType, error) {
	vals := map[string]CloudPCConnectivityEventEventType{
		"devicehealthcheck":   CloudPCConnectivityEventEventTypedeviceHealthCheck,
		"unknown":             CloudPCConnectivityEventEventTypeunknown,
		"userconnection":      CloudPCConnectivityEventEventTypeuserConnection,
		"usertroubleshooting": CloudPCConnectivityEventEventTypeuserTroubleshooting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventEventType(input)
	return &out, nil
}
