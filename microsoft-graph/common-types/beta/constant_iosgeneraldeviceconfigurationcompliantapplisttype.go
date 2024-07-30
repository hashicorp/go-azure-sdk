package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationCompliantAppListType string

const (
	IosGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant    IosGeneralDeviceConfigurationCompliantAppListType = "appsInListCompliant"
	IosGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant IosGeneralDeviceConfigurationCompliantAppListType = "appsNotInListCompliant"
	IosGeneralDeviceConfigurationCompliantAppListType_None                   IosGeneralDeviceConfigurationCompliantAppListType = "none"
)

func PossibleValuesForIosGeneralDeviceConfigurationCompliantAppListType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant),
		string(IosGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant),
		string(IosGeneralDeviceConfigurationCompliantAppListType_None),
	}
}

func (s *IosGeneralDeviceConfigurationCompliantAppListType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationCompliantAppListType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationCompliantAppListType(input string) (*IosGeneralDeviceConfigurationCompliantAppListType, error) {
	vals := map[string]IosGeneralDeviceConfigurationCompliantAppListType{
		"appsinlistcompliant":    IosGeneralDeviceConfigurationCompliantAppListType_AppsInListCompliant,
		"appsnotinlistcompliant": IosGeneralDeviceConfigurationCompliantAppListType_AppsNotInListCompliant,
		"none":                   IosGeneralDeviceConfigurationCompliantAppListType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationCompliantAppListType(input)
	return &out, nil
}
