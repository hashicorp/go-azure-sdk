package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDistributionMethod string

const (
	TeamsAppDistributionMethodorganization TeamsAppDistributionMethod = "Organization"
	TeamsAppDistributionMethodsideloaded   TeamsAppDistributionMethod = "Sideloaded"
	TeamsAppDistributionMethodstore        TeamsAppDistributionMethod = "Store"
)

func PossibleValuesForTeamsAppDistributionMethod() []string {
	return []string{
		string(TeamsAppDistributionMethodorganization),
		string(TeamsAppDistributionMethodsideloaded),
		string(TeamsAppDistributionMethodstore),
	}
}

func parseTeamsAppDistributionMethod(input string) (*TeamsAppDistributionMethod, error) {
	vals := map[string]TeamsAppDistributionMethod{
		"organization": TeamsAppDistributionMethodorganization,
		"sideloaded":   TeamsAppDistributionMethodsideloaded,
		"store":        TeamsAppDistributionMethodstore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDistributionMethod(input)
	return &out, nil
}
