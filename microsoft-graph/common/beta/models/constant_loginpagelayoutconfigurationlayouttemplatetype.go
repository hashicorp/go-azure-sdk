package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageLayoutConfigurationLayoutTemplateType string

const (
	LoginPageLayoutConfigurationLayoutTemplateTypedefault       LoginPageLayoutConfigurationLayoutTemplateType = "Default"
	LoginPageLayoutConfigurationLayoutTemplateTypeverticalSplit LoginPageLayoutConfigurationLayoutTemplateType = "VerticalSplit"
)

func PossibleValuesForLoginPageLayoutConfigurationLayoutTemplateType() []string {
	return []string{
		string(LoginPageLayoutConfigurationLayoutTemplateTypedefault),
		string(LoginPageLayoutConfigurationLayoutTemplateTypeverticalSplit),
	}
}

func parseLoginPageLayoutConfigurationLayoutTemplateType(input string) (*LoginPageLayoutConfigurationLayoutTemplateType, error) {
	vals := map[string]LoginPageLayoutConfigurationLayoutTemplateType{
		"default":       LoginPageLayoutConfigurationLayoutTemplateTypedefault,
		"verticalsplit": LoginPageLayoutConfigurationLayoutTemplateTypeverticalSplit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginPageLayoutConfigurationLayoutTemplateType(input)
	return &out, nil
}
