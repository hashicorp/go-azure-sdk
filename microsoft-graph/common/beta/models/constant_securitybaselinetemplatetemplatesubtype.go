package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplateTemplateSubtype string

const (
	SecurityBaselineTemplateTemplateSubtypeaccountProtection        SecurityBaselineTemplateTemplateSubtype = "AccountProtection"
	SecurityBaselineTemplateTemplateSubtypeantivirus                SecurityBaselineTemplateTemplateSubtype = "Antivirus"
	SecurityBaselineTemplateTemplateSubtypeattackSurfaceReduction   SecurityBaselineTemplateTemplateSubtype = "AttackSurfaceReduction"
	SecurityBaselineTemplateTemplateSubtypediskEncryption           SecurityBaselineTemplateTemplateSubtype = "DiskEncryption"
	SecurityBaselineTemplateTemplateSubtypeendpointDetectionReponse SecurityBaselineTemplateTemplateSubtype = "EndpointDetectionReponse"
	SecurityBaselineTemplateTemplateSubtypefirewall                 SecurityBaselineTemplateTemplateSubtype = "Firewall"
	SecurityBaselineTemplateTemplateSubtypefirewallSharedAppList    SecurityBaselineTemplateTemplateSubtype = "FirewallSharedAppList"
	SecurityBaselineTemplateTemplateSubtypefirewallSharedIpList     SecurityBaselineTemplateTemplateSubtype = "FirewallSharedIpList"
	SecurityBaselineTemplateTemplateSubtypefirewallSharedPortlist   SecurityBaselineTemplateTemplateSubtype = "FirewallSharedPortlist"
	SecurityBaselineTemplateTemplateSubtypenone                     SecurityBaselineTemplateTemplateSubtype = "None"
)

func PossibleValuesForSecurityBaselineTemplateTemplateSubtype() []string {
	return []string{
		string(SecurityBaselineTemplateTemplateSubtypeaccountProtection),
		string(SecurityBaselineTemplateTemplateSubtypeantivirus),
		string(SecurityBaselineTemplateTemplateSubtypeattackSurfaceReduction),
		string(SecurityBaselineTemplateTemplateSubtypediskEncryption),
		string(SecurityBaselineTemplateTemplateSubtypeendpointDetectionReponse),
		string(SecurityBaselineTemplateTemplateSubtypefirewall),
		string(SecurityBaselineTemplateTemplateSubtypefirewallSharedAppList),
		string(SecurityBaselineTemplateTemplateSubtypefirewallSharedIpList),
		string(SecurityBaselineTemplateTemplateSubtypefirewallSharedPortlist),
		string(SecurityBaselineTemplateTemplateSubtypenone),
	}
}

func parseSecurityBaselineTemplateTemplateSubtype(input string) (*SecurityBaselineTemplateTemplateSubtype, error) {
	vals := map[string]SecurityBaselineTemplateTemplateSubtype{
		"accountprotection":        SecurityBaselineTemplateTemplateSubtypeaccountProtection,
		"antivirus":                SecurityBaselineTemplateTemplateSubtypeantivirus,
		"attacksurfacereduction":   SecurityBaselineTemplateTemplateSubtypeattackSurfaceReduction,
		"diskencryption":           SecurityBaselineTemplateTemplateSubtypediskEncryption,
		"endpointdetectionreponse": SecurityBaselineTemplateTemplateSubtypeendpointDetectionReponse,
		"firewall":                 SecurityBaselineTemplateTemplateSubtypefirewall,
		"firewallsharedapplist":    SecurityBaselineTemplateTemplateSubtypefirewallSharedAppList,
		"firewallsharediplist":     SecurityBaselineTemplateTemplateSubtypefirewallSharedIpList,
		"firewallsharedportlist":   SecurityBaselineTemplateTemplateSubtypefirewallSharedPortlist,
		"none":                     SecurityBaselineTemplateTemplateSubtypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplateTemplateSubtype(input)
	return &out, nil
}
