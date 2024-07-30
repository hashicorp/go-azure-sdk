package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplateTemplateSubtype string

const (
	SecurityBaselineTemplateTemplateSubtype_AccountProtection        SecurityBaselineTemplateTemplateSubtype = "accountProtection"
	SecurityBaselineTemplateTemplateSubtype_Antivirus                SecurityBaselineTemplateTemplateSubtype = "antivirus"
	SecurityBaselineTemplateTemplateSubtype_AttackSurfaceReduction   SecurityBaselineTemplateTemplateSubtype = "attackSurfaceReduction"
	SecurityBaselineTemplateTemplateSubtype_DiskEncryption           SecurityBaselineTemplateTemplateSubtype = "diskEncryption"
	SecurityBaselineTemplateTemplateSubtype_EndpointDetectionReponse SecurityBaselineTemplateTemplateSubtype = "endpointDetectionReponse"
	SecurityBaselineTemplateTemplateSubtype_Firewall                 SecurityBaselineTemplateTemplateSubtype = "firewall"
	SecurityBaselineTemplateTemplateSubtype_FirewallSharedAppList    SecurityBaselineTemplateTemplateSubtype = "firewallSharedAppList"
	SecurityBaselineTemplateTemplateSubtype_FirewallSharedIpList     SecurityBaselineTemplateTemplateSubtype = "firewallSharedIpList"
	SecurityBaselineTemplateTemplateSubtype_FirewallSharedPortlist   SecurityBaselineTemplateTemplateSubtype = "firewallSharedPortlist"
	SecurityBaselineTemplateTemplateSubtype_None                     SecurityBaselineTemplateTemplateSubtype = "none"
)

func PossibleValuesForSecurityBaselineTemplateTemplateSubtype() []string {
	return []string{
		string(SecurityBaselineTemplateTemplateSubtype_AccountProtection),
		string(SecurityBaselineTemplateTemplateSubtype_Antivirus),
		string(SecurityBaselineTemplateTemplateSubtype_AttackSurfaceReduction),
		string(SecurityBaselineTemplateTemplateSubtype_DiskEncryption),
		string(SecurityBaselineTemplateTemplateSubtype_EndpointDetectionReponse),
		string(SecurityBaselineTemplateTemplateSubtype_Firewall),
		string(SecurityBaselineTemplateTemplateSubtype_FirewallSharedAppList),
		string(SecurityBaselineTemplateTemplateSubtype_FirewallSharedIpList),
		string(SecurityBaselineTemplateTemplateSubtype_FirewallSharedPortlist),
		string(SecurityBaselineTemplateTemplateSubtype_None),
	}
}

func (s *SecurityBaselineTemplateTemplateSubtype) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineTemplateTemplateSubtype(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineTemplateTemplateSubtype(input string) (*SecurityBaselineTemplateTemplateSubtype, error) {
	vals := map[string]SecurityBaselineTemplateTemplateSubtype{
		"accountprotection":        SecurityBaselineTemplateTemplateSubtype_AccountProtection,
		"antivirus":                SecurityBaselineTemplateTemplateSubtype_Antivirus,
		"attacksurfacereduction":   SecurityBaselineTemplateTemplateSubtype_AttackSurfaceReduction,
		"diskencryption":           SecurityBaselineTemplateTemplateSubtype_DiskEncryption,
		"endpointdetectionreponse": SecurityBaselineTemplateTemplateSubtype_EndpointDetectionReponse,
		"firewall":                 SecurityBaselineTemplateTemplateSubtype_Firewall,
		"firewallsharedapplist":    SecurityBaselineTemplateTemplateSubtype_FirewallSharedAppList,
		"firewallsharediplist":     SecurityBaselineTemplateTemplateSubtype_FirewallSharedIpList,
		"firewallsharedportlist":   SecurityBaselineTemplateTemplateSubtype_FirewallSharedPortlist,
		"none":                     SecurityBaselineTemplateTemplateSubtype_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplateTemplateSubtype(input)
	return &out, nil
}
