package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateTemplateSubtype string

const (
	DeviceManagementTemplateTemplateSubtype_AccountProtection        DeviceManagementTemplateTemplateSubtype = "accountProtection"
	DeviceManagementTemplateTemplateSubtype_Antivirus                DeviceManagementTemplateTemplateSubtype = "antivirus"
	DeviceManagementTemplateTemplateSubtype_AttackSurfaceReduction   DeviceManagementTemplateTemplateSubtype = "attackSurfaceReduction"
	DeviceManagementTemplateTemplateSubtype_DiskEncryption           DeviceManagementTemplateTemplateSubtype = "diskEncryption"
	DeviceManagementTemplateTemplateSubtype_EndpointDetectionReponse DeviceManagementTemplateTemplateSubtype = "endpointDetectionReponse"
	DeviceManagementTemplateTemplateSubtype_Firewall                 DeviceManagementTemplateTemplateSubtype = "firewall"
	DeviceManagementTemplateTemplateSubtype_FirewallSharedAppList    DeviceManagementTemplateTemplateSubtype = "firewallSharedAppList"
	DeviceManagementTemplateTemplateSubtype_FirewallSharedIpList     DeviceManagementTemplateTemplateSubtype = "firewallSharedIpList"
	DeviceManagementTemplateTemplateSubtype_FirewallSharedPortlist   DeviceManagementTemplateTemplateSubtype = "firewallSharedPortlist"
	DeviceManagementTemplateTemplateSubtype_None                     DeviceManagementTemplateTemplateSubtype = "none"
)

func PossibleValuesForDeviceManagementTemplateTemplateSubtype() []string {
	return []string{
		string(DeviceManagementTemplateTemplateSubtype_AccountProtection),
		string(DeviceManagementTemplateTemplateSubtype_Antivirus),
		string(DeviceManagementTemplateTemplateSubtype_AttackSurfaceReduction),
		string(DeviceManagementTemplateTemplateSubtype_DiskEncryption),
		string(DeviceManagementTemplateTemplateSubtype_EndpointDetectionReponse),
		string(DeviceManagementTemplateTemplateSubtype_Firewall),
		string(DeviceManagementTemplateTemplateSubtype_FirewallSharedAppList),
		string(DeviceManagementTemplateTemplateSubtype_FirewallSharedIpList),
		string(DeviceManagementTemplateTemplateSubtype_FirewallSharedPortlist),
		string(DeviceManagementTemplateTemplateSubtype_None),
	}
}

func (s *DeviceManagementTemplateTemplateSubtype) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplateTemplateSubtype(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplateTemplateSubtype(input string) (*DeviceManagementTemplateTemplateSubtype, error) {
	vals := map[string]DeviceManagementTemplateTemplateSubtype{
		"accountprotection":        DeviceManagementTemplateTemplateSubtype_AccountProtection,
		"antivirus":                DeviceManagementTemplateTemplateSubtype_Antivirus,
		"attacksurfacereduction":   DeviceManagementTemplateTemplateSubtype_AttackSurfaceReduction,
		"diskencryption":           DeviceManagementTemplateTemplateSubtype_DiskEncryption,
		"endpointdetectionreponse": DeviceManagementTemplateTemplateSubtype_EndpointDetectionReponse,
		"firewall":                 DeviceManagementTemplateTemplateSubtype_Firewall,
		"firewallsharedapplist":    DeviceManagementTemplateTemplateSubtype_FirewallSharedAppList,
		"firewallsharediplist":     DeviceManagementTemplateTemplateSubtype_FirewallSharedIpList,
		"firewallsharedportlist":   DeviceManagementTemplateTemplateSubtype_FirewallSharedPortlist,
		"none":                     DeviceManagementTemplateTemplateSubtype_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateTemplateSubtype(input)
	return &out, nil
}
