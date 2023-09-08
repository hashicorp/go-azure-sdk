package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSubscriptions string

const (
	DeviceManagementSubscriptionsintune        DeviceManagementSubscriptions = "Intune"
	DeviceManagementSubscriptionsintunePremium DeviceManagementSubscriptions = "IntunePremium"
	DeviceManagementSubscriptionsintune_EDU    DeviceManagementSubscriptions = "IntuneEDU"
	DeviceManagementSubscriptionsintune_SMB    DeviceManagementSubscriptions = "IntuneSMB"
	DeviceManagementSubscriptionsnone          DeviceManagementSubscriptions = "None"
	DeviceManagementSubscriptionsoffice365     DeviceManagementSubscriptions = "Office365"
)

func PossibleValuesForDeviceManagementSubscriptions() []string {
	return []string{
		string(DeviceManagementSubscriptionsintune),
		string(DeviceManagementSubscriptionsintunePremium),
		string(DeviceManagementSubscriptionsintune_EDU),
		string(DeviceManagementSubscriptionsintune_SMB),
		string(DeviceManagementSubscriptionsnone),
		string(DeviceManagementSubscriptionsoffice365),
	}
}

func parseDeviceManagementSubscriptions(input string) (*DeviceManagementSubscriptions, error) {
	vals := map[string]DeviceManagementSubscriptions{
		"intune":        DeviceManagementSubscriptionsintune,
		"intunepremium": DeviceManagementSubscriptionsintunePremium,
		"intuneedu":     DeviceManagementSubscriptionsintune_EDU,
		"intunesmb":     DeviceManagementSubscriptionsintune_SMB,
		"none":          DeviceManagementSubscriptionsnone,
		"office365":     DeviceManagementSubscriptionsoffice365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSubscriptions(input)
	return &out, nil
}
