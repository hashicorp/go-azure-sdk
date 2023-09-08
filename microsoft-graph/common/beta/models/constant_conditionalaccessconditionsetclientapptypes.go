package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetClientAppTypes string

const (
	ConditionalAccessConditionSetClientAppTypesall                         ConditionalAccessConditionSetClientAppTypes = "All"
	ConditionalAccessConditionSetClientAppTypesbrowser                     ConditionalAccessConditionSetClientAppTypes = "Browser"
	ConditionalAccessConditionSetClientAppTypeseasSupported                ConditionalAccessConditionSetClientAppTypes = "EasSupported"
	ConditionalAccessConditionSetClientAppTypesexchangeActiveSync          ConditionalAccessConditionSetClientAppTypes = "ExchangeActiveSync"
	ConditionalAccessConditionSetClientAppTypesmobileAppsAndDesktopClients ConditionalAccessConditionSetClientAppTypes = "MobileAppsAndDesktopClients"
	ConditionalAccessConditionSetClientAppTypesother                       ConditionalAccessConditionSetClientAppTypes = "Other"
)

func PossibleValuesForConditionalAccessConditionSetClientAppTypes() []string {
	return []string{
		string(ConditionalAccessConditionSetClientAppTypesall),
		string(ConditionalAccessConditionSetClientAppTypesbrowser),
		string(ConditionalAccessConditionSetClientAppTypeseasSupported),
		string(ConditionalAccessConditionSetClientAppTypesexchangeActiveSync),
		string(ConditionalAccessConditionSetClientAppTypesmobileAppsAndDesktopClients),
		string(ConditionalAccessConditionSetClientAppTypesother),
	}
}

func parseConditionalAccessConditionSetClientAppTypes(input string) (*ConditionalAccessConditionSetClientAppTypes, error) {
	vals := map[string]ConditionalAccessConditionSetClientAppTypes{
		"all":                         ConditionalAccessConditionSetClientAppTypesall,
		"browser":                     ConditionalAccessConditionSetClientAppTypesbrowser,
		"eassupported":                ConditionalAccessConditionSetClientAppTypeseasSupported,
		"exchangeactivesync":          ConditionalAccessConditionSetClientAppTypesexchangeActiveSync,
		"mobileappsanddesktopclients": ConditionalAccessConditionSetClientAppTypesmobileAppsAndDesktopClients,
		"other":                       ConditionalAccessConditionSetClientAppTypesother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetClientAppTypes(input)
	return &out, nil
}
