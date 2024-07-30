package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceChassisType string

const (
	ManagedDeviceChassisType_Desktop          ManagedDeviceChassisType = "desktop"
	ManagedDeviceChassisType_EnterpriseServer ManagedDeviceChassisType = "enterpriseServer"
	ManagedDeviceChassisType_Laptop           ManagedDeviceChassisType = "laptop"
	ManagedDeviceChassisType_MobileOther      ManagedDeviceChassisType = "mobileOther"
	ManagedDeviceChassisType_MobileUnknown    ManagedDeviceChassisType = "mobileUnknown"
	ManagedDeviceChassisType_Phone            ManagedDeviceChassisType = "phone"
	ManagedDeviceChassisType_Tablet           ManagedDeviceChassisType = "tablet"
	ManagedDeviceChassisType_Unknown          ManagedDeviceChassisType = "unknown"
	ManagedDeviceChassisType_WorksWorkstation ManagedDeviceChassisType = "worksWorkstation"
)

func PossibleValuesForManagedDeviceChassisType() []string {
	return []string{
		string(ManagedDeviceChassisType_Desktop),
		string(ManagedDeviceChassisType_EnterpriseServer),
		string(ManagedDeviceChassisType_Laptop),
		string(ManagedDeviceChassisType_MobileOther),
		string(ManagedDeviceChassisType_MobileUnknown),
		string(ManagedDeviceChassisType_Phone),
		string(ManagedDeviceChassisType_Tablet),
		string(ManagedDeviceChassisType_Unknown),
		string(ManagedDeviceChassisType_WorksWorkstation),
	}
}

func (s *ManagedDeviceChassisType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceChassisType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceChassisType(input string) (*ManagedDeviceChassisType, error) {
	vals := map[string]ManagedDeviceChassisType{
		"desktop":          ManagedDeviceChassisType_Desktop,
		"enterpriseserver": ManagedDeviceChassisType_EnterpriseServer,
		"laptop":           ManagedDeviceChassisType_Laptop,
		"mobileother":      ManagedDeviceChassisType_MobileOther,
		"mobileunknown":    ManagedDeviceChassisType_MobileUnknown,
		"phone":            ManagedDeviceChassisType_Phone,
		"tablet":           ManagedDeviceChassisType_Tablet,
		"unknown":          ManagedDeviceChassisType_Unknown,
		"worksworkstation": ManagedDeviceChassisType_WorksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceChassisType(input)
	return &out, nil
}
