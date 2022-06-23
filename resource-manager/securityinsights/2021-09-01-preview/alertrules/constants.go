package alertrules

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRuleKind string

const (
	AlertRuleKindFusion                            AlertRuleKind = "Fusion"
	AlertRuleKindMLBehaviorAnalytics               AlertRuleKind = "MLBehaviorAnalytics"
	AlertRuleKindMicrosoftSecurityIncidentCreation AlertRuleKind = "MicrosoftSecurityIncidentCreation"
	AlertRuleKindNRT                               AlertRuleKind = "NRT"
	AlertRuleKindScheduled                         AlertRuleKind = "Scheduled"
	AlertRuleKindThreatIntelligence                AlertRuleKind = "ThreatIntelligence"
)

func PossibleValuesForAlertRuleKind() []string {
	return []string{
		string(AlertRuleKindFusion),
		string(AlertRuleKindMLBehaviorAnalytics),
		string(AlertRuleKindMicrosoftSecurityIncidentCreation),
		string(AlertRuleKindNRT),
		string(AlertRuleKindScheduled),
		string(AlertRuleKindThreatIntelligence),
	}
}

func parseAlertRuleKind(input string) (*AlertRuleKind, error) {
	vals := map[string]AlertRuleKind{
		"fusion":                            AlertRuleKindFusion,
		"mlbehavioranalytics":               AlertRuleKindMLBehaviorAnalytics,
		"microsoftsecurityincidentcreation": AlertRuleKindMicrosoftSecurityIncidentCreation,
		"nrt":                               AlertRuleKindNRT,
		"scheduled":                         AlertRuleKindScheduled,
		"threatintelligence":                AlertRuleKindThreatIntelligence,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertRuleKind(input)
	return &out, nil
}
