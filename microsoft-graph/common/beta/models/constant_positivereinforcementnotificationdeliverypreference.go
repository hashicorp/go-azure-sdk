package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositiveReinforcementNotificationDeliveryPreference string

const (
	PositiveReinforcementNotificationDeliveryPreferencedeliverAfterCampaignEnd PositiveReinforcementNotificationDeliveryPreference = "DeliverAfterCampaignEnd"
	PositiveReinforcementNotificationDeliveryPreferencedeliverImmedietly       PositiveReinforcementNotificationDeliveryPreference = "DeliverImmedietly"
	PositiveReinforcementNotificationDeliveryPreferenceunknown                 PositiveReinforcementNotificationDeliveryPreference = "Unknown"
)

func PossibleValuesForPositiveReinforcementNotificationDeliveryPreference() []string {
	return []string{
		string(PositiveReinforcementNotificationDeliveryPreferencedeliverAfterCampaignEnd),
		string(PositiveReinforcementNotificationDeliveryPreferencedeliverImmedietly),
		string(PositiveReinforcementNotificationDeliveryPreferenceunknown),
	}
}

func parsePositiveReinforcementNotificationDeliveryPreference(input string) (*PositiveReinforcementNotificationDeliveryPreference, error) {
	vals := map[string]PositiveReinforcementNotificationDeliveryPreference{
		"deliveraftercampaignend": PositiveReinforcementNotificationDeliveryPreferencedeliverAfterCampaignEnd,
		"deliverimmedietly":       PositiveReinforcementNotificationDeliveryPreferencedeliverImmedietly,
		"unknown":                 PositiveReinforcementNotificationDeliveryPreferenceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PositiveReinforcementNotificationDeliveryPreference(input)
	return &out, nil
}
