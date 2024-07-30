package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationCompliantAppListType string

const (
	AndroidGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant    AndroidGeneralDeviceConfigurationCompliantAppListType = "appsInListCompliant"
	AndroidGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant AndroidGeneralDeviceConfigurationCompliantAppListType = "appsNotInListCompliant"
	AndroidGeneralDeviceConfigurationCompliantAppListType_None                   AndroidGeneralDeviceConfigurationCompliantAppListType = "none"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant),
		string(AndroidGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant),
		string(AndroidGeneralDeviceConfigurationCompliantAppListType_None),
	}
}

func (s *AndroidGeneralDeviceConfigurationCompliantAppListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidGeneralDeviceConfigurationCompliantAppListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidGeneralDeviceConfigurationCompliantAppListType(input string) (*AndroidGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    AndroidGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant,
		"appsnotinlistcompliant": AndroidGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant,
		"none":                   AndroidGeneralDeviceConfigurationCompliantAppListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
