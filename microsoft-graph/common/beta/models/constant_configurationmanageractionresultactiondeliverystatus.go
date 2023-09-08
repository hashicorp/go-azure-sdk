package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionResultActionDeliveryStatus string

const (
	ConfigurationManagerActionResultActionDeliveryStatusdeliveredToConnectorService       ConfigurationManagerActionResultActionDeliveryStatus = "DeliveredToConnectorService"
	ConfigurationManagerActionResultActionDeliveryStatusdeliveredToOnPremisesServer       ConfigurationManagerActionResultActionDeliveryStatus = "DeliveredToOnPremisesServer"
	ConfigurationManagerActionResultActionDeliveryStatusfailedToDeliverToConnectorService ConfigurationManagerActionResultActionDeliveryStatus = "FailedToDeliverToConnectorService"
	ConfigurationManagerActionResultActionDeliveryStatuspendingDelivery                   ConfigurationManagerActionResultActionDeliveryStatus = "PendingDelivery"
	ConfigurationManagerActionResultActionDeliveryStatusunknown                           ConfigurationManagerActionResultActionDeliveryStatus = "Unknown"
)

func PossibleValuesForConfigurationManagerActionResultActionDeliveryStatus() []string {
	return []string{
		string(ConfigurationManagerActionResultActionDeliveryStatusdeliveredToConnectorService),
		string(ConfigurationManagerActionResultActionDeliveryStatusdeliveredToOnPremisesServer),
		string(ConfigurationManagerActionResultActionDeliveryStatusfailedToDeliverToConnectorService),
		string(ConfigurationManagerActionResultActionDeliveryStatuspendingDelivery),
		string(ConfigurationManagerActionResultActionDeliveryStatusunknown),
	}
}

func parseConfigurationManagerActionResultActionDeliveryStatus(input string) (*ConfigurationManagerActionResultActionDeliveryStatus, error) {
	vals := map[string]ConfigurationManagerActionResultActionDeliveryStatus{
		"deliveredtoconnectorservice":       ConfigurationManagerActionResultActionDeliveryStatusdeliveredToConnectorService,
		"deliveredtoonpremisesserver":       ConfigurationManagerActionResultActionDeliveryStatusdeliveredToOnPremisesServer,
		"failedtodelivertoconnectorservice": ConfigurationManagerActionResultActionDeliveryStatusfailedToDeliverToConnectorService,
		"pendingdelivery":                   ConfigurationManagerActionResultActionDeliveryStatuspendingDelivery,
		"unknown":                           ConfigurationManagerActionResultActionDeliveryStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionResultActionDeliveryStatus(input)
	return &out, nil
}
