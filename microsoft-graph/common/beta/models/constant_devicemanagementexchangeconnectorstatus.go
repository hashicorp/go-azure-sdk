package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorStatus string

const (
	DeviceManagementExchangeConnectorStatusconnected         DeviceManagementExchangeConnectorStatus = "Connected"
	DeviceManagementExchangeConnectorStatusconnectionPending DeviceManagementExchangeConnectorStatus = "ConnectionPending"
	DeviceManagementExchangeConnectorStatusdisconnected      DeviceManagementExchangeConnectorStatus = "Disconnected"
	DeviceManagementExchangeConnectorStatusnone              DeviceManagementExchangeConnectorStatus = "None"
)

func PossibleValuesForDeviceManagementExchangeConnectorStatus() []string {
	return []string{
		string(DeviceManagementExchangeConnectorStatusconnected),
		string(DeviceManagementExchangeConnectorStatusconnectionPending),
		string(DeviceManagementExchangeConnectorStatusdisconnected),
		string(DeviceManagementExchangeConnectorStatusnone),
	}
}

func parseDeviceManagementExchangeConnectorStatus(input string) (*DeviceManagementExchangeConnectorStatus, error) {
	vals := map[string]DeviceManagementExchangeConnectorStatus{
		"connected":         DeviceManagementExchangeConnectorStatusconnected,
		"connectionpending": DeviceManagementExchangeConnectorStatusconnectionPending,
		"disconnected":      DeviceManagementExchangeConnectorStatusdisconnected,
		"none":              DeviceManagementExchangeConnectorStatusnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorStatus(input)
	return &out, nil
}
