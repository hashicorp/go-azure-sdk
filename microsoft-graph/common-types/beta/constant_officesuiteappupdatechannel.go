package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppUpdateChannel string

const (
	OfficeSuiteAppUpdateChannel_Current              OfficeSuiteAppUpdateChannel = "current"
	OfficeSuiteAppUpdateChannel_Deferred             OfficeSuiteAppUpdateChannel = "deferred"
	OfficeSuiteAppUpdateChannel_FirstReleaseCurrent  OfficeSuiteAppUpdateChannel = "firstReleaseCurrent"
	OfficeSuiteAppUpdateChannel_FirstReleaseDeferred OfficeSuiteAppUpdateChannel = "firstReleaseDeferred"
	OfficeSuiteAppUpdateChannel_MonthlyEnterprise    OfficeSuiteAppUpdateChannel = "monthlyEnterprise"
	OfficeSuiteAppUpdateChannel_None                 OfficeSuiteAppUpdateChannel = "none"
)

func PossibleValuesForOfficeSuiteAppUpdateChannel() []string {
	return []string{
		string(OfficeSuiteAppUpdateChannel_Current),
		string(OfficeSuiteAppUpdateChannel_Deferred),
		string(OfficeSuiteAppUpdateChannel_FirstReleaseCurrent),
		string(OfficeSuiteAppUpdateChannel_FirstReleaseDeferred),
		string(OfficeSuiteAppUpdateChannel_MonthlyEnterprise),
		string(OfficeSuiteAppUpdateChannel_None),
	}
}

func (s *OfficeSuiteAppUpdateChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppUpdateChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppUpdateChannel(input string) (*OfficeSuiteAppUpdateChannel, error) {
	vals := map[string]OfficeSuiteAppUpdateChannel{
		"current":              OfficeSuiteAppUpdateChannel_Current,
		"deferred":             OfficeSuiteAppUpdateChannel_Deferred,
		"firstreleasecurrent":  OfficeSuiteAppUpdateChannel_FirstReleaseCurrent,
		"firstreleasedeferred": OfficeSuiteAppUpdateChannel_FirstReleaseDeferred,
		"monthlyenterprise":    OfficeSuiteAppUpdateChannel_MonthlyEnterprise,
		"none":                 OfficeSuiteAppUpdateChannel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppUpdateChannel(input)
	return &out, nil
}
