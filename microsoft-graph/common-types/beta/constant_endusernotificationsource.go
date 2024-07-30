package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSource string

const (
	EndUserNotificationSource_Global  EndUserNotificationSource = "global"
	EndUserNotificationSource_Tenant  EndUserNotificationSource = "tenant"
	EndUserNotificationSource_Unknown EndUserNotificationSource = "unknown"
)

func PossibleValuesForEndUserNotificationSource() []string {
	return []string{
		string(EndUserNotificationSource_Global),
		string(EndUserNotificationSource_Tenant),
		string(EndUserNotificationSource_Unknown),
	}
}

func (s *EndUserNotificationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationSource(input string) (*EndUserNotificationSource, error) {
	vals := map[string]EndUserNotificationSource{
		"global":  EndUserNotificationSource_Global,
		"tenant":  EndUserNotificationSource_Tenant,
		"unknown": EndUserNotificationSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSource(input)
	return &out, nil
}
