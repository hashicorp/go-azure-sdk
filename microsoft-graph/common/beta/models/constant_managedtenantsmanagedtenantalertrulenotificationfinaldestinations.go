package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations string

const (
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsapi   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "Api"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsemail ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "Email"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsnone  ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "None"
	ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationssms   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations = "Sms"
)

func PossibleValuesForManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsapi),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsemail),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsnone),
		string(ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationssms),
	}
}

func parseManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations(input string) (*ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations{
		"api":   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsapi,
		"email": ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsemail,
		"none":  ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationsnone,
		"sms":   ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinationssms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertRuleNotificationFinalDestinations(input)
	return &out, nil
}
