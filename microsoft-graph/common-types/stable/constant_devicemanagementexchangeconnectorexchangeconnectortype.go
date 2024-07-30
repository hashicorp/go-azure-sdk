package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorExchangeConnectorType string

const (
	DeviceManagementExchangeConnectorExchangeConnectorType_Dedicated        DeviceManagementExchangeConnectorExchangeConnectorType = "dedicated"
	DeviceManagementExchangeConnectorExchangeConnectorType_Hosted           DeviceManagementExchangeConnectorExchangeConnectorType = "hosted"
	DeviceManagementExchangeConnectorExchangeConnectorType_OnPremises       DeviceManagementExchangeConnectorExchangeConnectorType = "onPremises"
	DeviceManagementExchangeConnectorExchangeConnectorType_ServiceToService DeviceManagementExchangeConnectorExchangeConnectorType = "serviceToService"
)

func PossibleValuesForDeviceManagementExchangeConnectorExchangeConnectorType() []string {
	return []string{
		string(DeviceManagementExchangeConnectorExchangeConnectorType_Dedicated),
		string(DeviceManagementExchangeConnectorExchangeConnectorType_Hosted),
		string(DeviceManagementExchangeConnectorExchangeConnectorType_OnPremises),
		string(DeviceManagementExchangeConnectorExchangeConnectorType_ServiceToService),
	}
}

func (s *DeviceManagementExchangeConnectorExchangeConnectorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeConnectorExchangeConnectorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeConnectorExchangeConnectorType(input string) (*DeviceManagementExchangeConnectorExchangeConnectorType, error) {
	vals := map[string]DeviceManagementExchangeConnectorExchangeConnectorType{
		"dedicated":        DeviceManagementExchangeConnectorExchangeConnectorType_Dedicated,
		"hosted":           DeviceManagementExchangeConnectorExchangeConnectorType_Hosted,
		"onpremises":       DeviceManagementExchangeConnectorExchangeConnectorType_OnPremises,
		"servicetoservice": DeviceManagementExchangeConnectorExchangeConnectorType_ServiceToService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorExchangeConnectorType(input)
	return &out, nil
}
