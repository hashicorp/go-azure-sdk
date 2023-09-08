package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessTemplateScenarios string

const (
	ConditionalAccessTemplateScenariosemergingThreats  ConditionalAccessTemplateScenarios = "EmergingThreats"
	ConditionalAccessTemplateScenariosnew              ConditionalAccessTemplateScenarios = "New"
	ConditionalAccessTemplateScenariosprotectAdmins    ConditionalAccessTemplateScenarios = "ProtectAdmins"
	ConditionalAccessTemplateScenariosremoteWork       ConditionalAccessTemplateScenarios = "RemoteWork"
	ConditionalAccessTemplateScenariossecureFoundation ConditionalAccessTemplateScenarios = "SecureFoundation"
	ConditionalAccessTemplateScenarioszeroTrust        ConditionalAccessTemplateScenarios = "ZeroTrust"
)

func PossibleValuesForConditionalAccessTemplateScenarios() []string {
	return []string{
		string(ConditionalAccessTemplateScenariosemergingThreats),
		string(ConditionalAccessTemplateScenariosnew),
		string(ConditionalAccessTemplateScenariosprotectAdmins),
		string(ConditionalAccessTemplateScenariosremoteWork),
		string(ConditionalAccessTemplateScenariossecureFoundation),
		string(ConditionalAccessTemplateScenarioszeroTrust),
	}
}

func parseConditionalAccessTemplateScenarios(input string) (*ConditionalAccessTemplateScenarios, error) {
	vals := map[string]ConditionalAccessTemplateScenarios{
		"emergingthreats":  ConditionalAccessTemplateScenariosemergingThreats,
		"new":              ConditionalAccessTemplateScenariosnew,
		"protectadmins":    ConditionalAccessTemplateScenariosprotectAdmins,
		"remotework":       ConditionalAccessTemplateScenariosremoteWork,
		"securefoundation": ConditionalAccessTemplateScenariossecureFoundation,
		"zerotrust":        ConditionalAccessTemplateScenarioszeroTrust,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessTemplateScenarios(input)
	return &out, nil
}
