package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetClientAppTypes string

const (
	ConditionalAccessConditionSetClientAppTypes_All                         ConditionalAccessConditionSetClientAppTypes = "all"
	ConditionalAccessConditionSetClientAppTypes_Browser                     ConditionalAccessConditionSetClientAppTypes = "browser"
	ConditionalAccessConditionSetClientAppTypes_EasSupported                ConditionalAccessConditionSetClientAppTypes = "easSupported"
	ConditionalAccessConditionSetClientAppTypes_ExchangeActiveSync          ConditionalAccessConditionSetClientAppTypes = "exchangeActiveSync"
	ConditionalAccessConditionSetClientAppTypes_MobileAppsAndDesktopClients ConditionalAccessConditionSetClientAppTypes = "mobileAppsAndDesktopClients"
	ConditionalAccessConditionSetClientAppTypes_Other                       ConditionalAccessConditionSetClientAppTypes = "other"
)

func PossibleValuesForConditionalAccessConditionSetClientAppTypes() []string {
	return []string{
		string(ConditionalAccessConditionSetClientAppTypes_All),
		string(ConditionalAccessConditionSetClientAppTypes_Browser),
		string(ConditionalAccessConditionSetClientAppTypes_EasSupported),
		string(ConditionalAccessConditionSetClientAppTypes_ExchangeActiveSync),
		string(ConditionalAccessConditionSetClientAppTypes_MobileAppsAndDesktopClients),
		string(ConditionalAccessConditionSetClientAppTypes_Other),
	}
}

func (s *ConditionalAccessConditionSetClientAppTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessConditionSetClientAppTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessConditionSetClientAppTypes(input string) (*ConditionalAccessConditionSetClientAppTypes, error) {
	vals := map[string]ConditionalAccessConditionSetClientAppTypes{
		"all":                         ConditionalAccessConditionSetClientAppTypes_All,
		"browser":                     ConditionalAccessConditionSetClientAppTypes_Browser,
		"eassupported":                ConditionalAccessConditionSetClientAppTypes_EasSupported,
		"exchangeactivesync":          ConditionalAccessConditionSetClientAppTypes_ExchangeActiveSync,
		"mobileappsanddesktopclients": ConditionalAccessConditionSetClientAppTypes_MobileAppsAndDesktopClients,
		"other":                       ConditionalAccessConditionSetClientAppTypes_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetClientAppTypes(input)
	return &out, nil
}
