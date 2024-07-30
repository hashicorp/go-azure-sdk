package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceChassisType string

const (
	WindowsManagedDeviceChassisType_Desktop          WindowsManagedDeviceChassisType = "desktop"
	WindowsManagedDeviceChassisType_EnterpriseServer WindowsManagedDeviceChassisType = "enterpriseServer"
	WindowsManagedDeviceChassisType_Laptop           WindowsManagedDeviceChassisType = "laptop"
	WindowsManagedDeviceChassisType_MobileOther      WindowsManagedDeviceChassisType = "mobileOther"
	WindowsManagedDeviceChassisType_MobileUnknown    WindowsManagedDeviceChassisType = "mobileUnknown"
	WindowsManagedDeviceChassisType_Phone            WindowsManagedDeviceChassisType = "phone"
	WindowsManagedDeviceChassisType_Tablet           WindowsManagedDeviceChassisType = "tablet"
	WindowsManagedDeviceChassisType_Unknown          WindowsManagedDeviceChassisType = "unknown"
	WindowsManagedDeviceChassisType_WorksWorkstation WindowsManagedDeviceChassisType = "worksWorkstation"
)

func PossibleValuesForWindowsManagedDeviceChassisType() []string {
	return []string{
		string(WindowsManagedDeviceChassisType_Desktop),
		string(WindowsManagedDeviceChassisType_EnterpriseServer),
		string(WindowsManagedDeviceChassisType_Laptop),
		string(WindowsManagedDeviceChassisType_MobileOther),
		string(WindowsManagedDeviceChassisType_MobileUnknown),
		string(WindowsManagedDeviceChassisType_Phone),
		string(WindowsManagedDeviceChassisType_Tablet),
		string(WindowsManagedDeviceChassisType_Unknown),
		string(WindowsManagedDeviceChassisType_WorksWorkstation),
	}
}

func (s *WindowsManagedDeviceChassisType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceChassisType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceChassisType(input string) (*WindowsManagedDeviceChassisType, error) {
	vals := map[string]WindowsManagedDeviceChassisType{
		"desktop":          WindowsManagedDeviceChassisType_Desktop,
		"enterpriseserver": WindowsManagedDeviceChassisType_EnterpriseServer,
		"laptop":           WindowsManagedDeviceChassisType_Laptop,
		"mobileother":      WindowsManagedDeviceChassisType_MobileOther,
		"mobileunknown":    WindowsManagedDeviceChassisType_MobileUnknown,
		"phone":            WindowsManagedDeviceChassisType_Phone,
		"tablet":           WindowsManagedDeviceChassisType_Tablet,
		"unknown":          WindowsManagedDeviceChassisType_Unknown,
		"worksworkstation": WindowsManagedDeviceChassisType_WorksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceChassisType(input)
	return &out, nil
}
