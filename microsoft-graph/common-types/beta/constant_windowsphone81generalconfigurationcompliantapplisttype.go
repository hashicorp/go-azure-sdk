package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81GeneralConfigurationCompliantAppListType string

const (
	WindowsPhone81GeneralConfigurationCompliantAppListType_AppsInListCompliant    WindowsPhone81GeneralConfigurationCompliantAppListType = "appsInListCompliant"
	WindowsPhone81GeneralConfigurationCompliantAppListType_AppsNotInListCompliant WindowsPhone81GeneralConfigurationCompliantAppListType = "appsNotInListCompliant"
	WindowsPhone81GeneralConfigurationCompliantAppListType_None                   WindowsPhone81GeneralConfigurationCompliantAppListType = "none"
)

func PossibleValuesForWindowsPhone81GeneralConfigurationCompliantAppListType() []string {
	return []string{
		string(WindowsPhone81GeneralConfigurationCompliantAppListType_AppsInListCompliant),
		string(WindowsPhone81GeneralConfigurationCompliantAppListType_AppsNotInListCompliant),
		string(WindowsPhone81GeneralConfigurationCompliantAppListType_None),
	}
}

func (s *WindowsPhone81GeneralConfigurationCompliantAppListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81GeneralConfigurationCompliantAppListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81GeneralConfigurationCompliantAppListType(input string) (*WindowsPhone81GeneralConfigurationCompliantAppListType, error) {
	vals := map[string]WindowsPhone81GeneralConfigurationCompliantAppListType{
		"appsinlistcompliant":    WindowsPhone81GeneralConfigurationCompliantAppListType_AppsInListCompliant,
		"appsnotinlistcompliant": WindowsPhone81GeneralConfigurationCompliantAppListType_AppsNotInListCompliant,
		"none":                   WindowsPhone81GeneralConfigurationCompliantAppListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81GeneralConfigurationCompliantAppListType(input)
	return &out, nil
}
