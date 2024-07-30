package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationCompliantAppListType string

const (
	MacOSGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant    MacOSGeneralDeviceConfigurationCompliantAppListType = "appsInListCompliant"
	MacOSGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant MacOSGeneralDeviceConfigurationCompliantAppListType = "appsNotInListCompliant"
	MacOSGeneralDeviceConfigurationCompliantAppListType_None                   MacOSGeneralDeviceConfigurationCompliantAppListType = "none"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant),
		string(MacOSGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant),
		string(MacOSGeneralDeviceConfigurationCompliantAppListType_None),
	}
}

func (s *MacOSGeneralDeviceConfigurationCompliantAppListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSGeneralDeviceConfigurationCompliantAppListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSGeneralDeviceConfigurationCompliantAppListType(input string) (*MacOSGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    MacOSGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant,
		"appsnotinlistcompliant": MacOSGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant,
		"none":                   MacOSGeneralDeviceConfigurationCompliantAppListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
