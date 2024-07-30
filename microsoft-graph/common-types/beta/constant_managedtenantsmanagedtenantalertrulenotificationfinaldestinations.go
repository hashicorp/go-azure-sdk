package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations string

const (
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Api   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "api"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Email ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "email"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_None  ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "none"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Sms   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "sms"
)

func PossibleValuesForManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Api),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Email),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_None),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Sms),
	}
}

func (s *ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations(input string) (*ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations{
		"api":   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Api,
		"email": ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Email,
		"none":  ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_None,
		"sms":   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations_Sms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations(input)
	return &out, nil
}
