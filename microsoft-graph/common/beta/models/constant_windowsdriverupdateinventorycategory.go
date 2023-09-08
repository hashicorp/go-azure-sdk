package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateInventoryCategory string

const (
	WindowsDriverUpdateInventoryCategoryother              WindowsDriverUpdateInventoryCategory = "Other"
	WindowsDriverUpdateInventoryCategorypreviouslyApproved WindowsDriverUpdateInventoryCategory = "PreviouslyApproved"
	WindowsDriverUpdateInventoryCategoryrecommended        WindowsDriverUpdateInventoryCategory = "Recommended"
)

func PossibleValuesForWindowsDriverUpdateInventoryCategory() []string {
	return []string{
		string(WindowsDriverUpdateInventoryCategoryother),
		string(WindowsDriverUpdateInventoryCategorypreviouslyApproved),
		string(WindowsDriverUpdateInventoryCategoryrecommended),
	}
}

func parseWindowsDriverUpdateInventoryCategory(input string) (*WindowsDriverUpdateInventoryCategory, error) {
	vals := map[string]WindowsDriverUpdateInventoryCategory{
		"other":              WindowsDriverUpdateInventoryCategoryother,
		"previouslyapproved": WindowsDriverUpdateInventoryCategorypreviouslyApproved,
		"recommended":        WindowsDriverUpdateInventoryCategoryrecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateInventoryCategory(input)
	return &out, nil
}
