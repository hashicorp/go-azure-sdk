package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositiveReinforcementNotificationDeliveryPreference string

const (
	PositiveReinforcementNotificationDeliveryPreference_DeliverAfterCampaignEnd PositiveReinforcementNotificationDeliveryPreference = "deliverAfterCampaignEnd"
	PositiveReinforcementNotificationDeliveryPreference_DeliverImmedietly       PositiveReinforcementNotificationDeliveryPreference = "deliverImmedietly"
	PositiveReinforcementNotificationDeliveryPreference_Unknown                 PositiveReinforcementNotificationDeliveryPreference = "unknown"
)

func PossibleValuesForPositiveReinforcementNotificationDeliveryPreference() []string {
	return []string{
		string(PositiveReinforcementNotificationDeliveryPreference_DeliverAfterCampaignEnd),
		string(PositiveReinforcementNotificationDeliveryPreference_DeliverImmedietly),
		string(PositiveReinforcementNotificationDeliveryPreference_Unknown),
	}
}

func (s *PositiveReinforcementNotificationDeliveryPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePositiveReinforcementNotificationDeliveryPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePositiveReinforcementNotificationDeliveryPreference(input string) (*PositiveReinforcementNotificationDeliveryPreference, error) {
	vals := map[string]PositiveReinforcementNotificationDeliveryPreference{
		"deliveraftercampaignend": PositiveReinforcementNotificationDeliveryPreference_DeliverAfterCampaignEnd,
		"deliverimmedietly":       PositiveReinforcementNotificationDeliveryPreference_DeliverImmedietly,
		"unknown":                 PositiveReinforcementNotificationDeliveryPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PositiveReinforcementNotificationDeliveryPreference(input)
	return &out, nil
}
