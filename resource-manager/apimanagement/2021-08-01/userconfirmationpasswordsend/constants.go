package userconfirmationpasswordsend

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppType string

const (
	AppTypeDeveloperPortal AppType = "developerPortal"
	AppTypePortal          AppType = "portal"
)

func PossibleValuesForAppType() []string {
	return []string{
		string(AppTypeDeveloperPortal),
		string(AppTypePortal),
	}
}

func parseAppType(input string) (*AppType, error) {
	vals := map[string]AppType{
		"developerportal": AppTypeDeveloperPortal,
		"portal":          AppTypePortal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppType(input)
	return &out, nil
}
