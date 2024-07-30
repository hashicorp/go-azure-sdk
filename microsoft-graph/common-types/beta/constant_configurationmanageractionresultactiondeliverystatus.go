package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionResultActionDeliveryStatus string

const (
	ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToConnectorService       ConfigurationManagerActionResultActionDeliveryStatus = "deliveredToConnectorService"
	ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToOnPremisesServer       ConfigurationManagerActionResultActionDeliveryStatus = "deliveredToOnPremisesServer"
	ConfigurationManagerActionResultActionDeliveryStatus_FailedToDeliverToConnectorService ConfigurationManagerActionResultActionDeliveryStatus = "failedToDeliverToConnectorService"
	ConfigurationManagerActionResultActionDeliveryStatus_PendingDelivery                   ConfigurationManagerActionResultActionDeliveryStatus = "pendingDelivery"
	ConfigurationManagerActionResultActionDeliveryStatus_Unknown                           ConfigurationManagerActionResultActionDeliveryStatus = "unknown"
)

func PossibleValuesForConfigurationManagerActionResultActionDeliveryStatus() []string {
	return []string{
		string(ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToConnectorService),
		string(ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToOnPremisesServer),
		string(ConfigurationManagerActionResultActionDeliveryStatus_FailedToDeliverToConnectorService),
		string(ConfigurationManagerActionResultActionDeliveryStatus_PendingDelivery),
		string(ConfigurationManagerActionResultActionDeliveryStatus_Unknown),
	}
}

func (s *ConfigurationManagerActionResultActionDeliveryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionResultActionDeliveryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionResultActionDeliveryStatus(input string) (*ConfigurationManagerActionResultActionDeliveryStatus, error) {
	vals := map[string]ConfigurationManagerActionResultActionDeliveryStatus{
		"deliveredtoconnectorservice":       ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToConnectorService,
		"deliveredtoonpremisesserver":       ConfigurationManagerActionResultActionDeliveryStatus_DeliveredToOnPremisesServer,
		"failedtodelivertoconnectorservice": ConfigurationManagerActionResultActionDeliveryStatus_FailedToDeliverToConnectorService,
		"pendingdelivery":                   ConfigurationManagerActionResultActionDeliveryStatus_PendingDelivery,
		"unknown":                           ConfigurationManagerActionResultActionDeliveryStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionResultActionDeliveryStatus(input)
	return &out, nil
}
