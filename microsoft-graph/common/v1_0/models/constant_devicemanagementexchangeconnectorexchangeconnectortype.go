package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorExchangeConnectorType string

const (
	DeviceManagementExchangeConnectorExchangeConnectorTypededicated        DeviceManagementExchangeConnectorExchangeConnectorType = "Dedicated"
	DeviceManagementExchangeConnectorExchangeConnectorTypehosted           DeviceManagementExchangeConnectorExchangeConnectorType = "Hosted"
	DeviceManagementExchangeConnectorExchangeConnectorTypeonPremises       DeviceManagementExchangeConnectorExchangeConnectorType = "OnPremises"
	DeviceManagementExchangeConnectorExchangeConnectorTypeserviceToService DeviceManagementExchangeConnectorExchangeConnectorType = "ServiceToService"
)

func PossibleValuesForDeviceManagementExchangeConnectorExchangeConnectorType() []string {
	return []string{
		string(DeviceManagementExchangeConnectorExchangeConnectorTypededicated),
		string(DeviceManagementExchangeConnectorExchangeConnectorTypehosted),
		string(DeviceManagementExchangeConnectorExchangeConnectorTypeonPremises),
		string(DeviceManagementExchangeConnectorExchangeConnectorTypeserviceToService),
	}
}

func parseDeviceManagementExchangeConnectorExchangeConnectorType(input string) (*DeviceManagementExchangeConnectorExchangeConnectorType, error) {
	vals := map[string]DeviceManagementExchangeConnectorExchangeConnectorType{
		"dedicated":        DeviceManagementExchangeConnectorExchangeConnectorTypededicated,
		"hosted":           DeviceManagementExchangeConnectorExchangeConnectorTypehosted,
		"onpremises":       DeviceManagementExchangeConnectorExchangeConnectorTypeonPremises,
		"servicetoservice": DeviceManagementExchangeConnectorExchangeConnectorTypeserviceToService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorExchangeConnectorType(input)
	return &out, nil
}
