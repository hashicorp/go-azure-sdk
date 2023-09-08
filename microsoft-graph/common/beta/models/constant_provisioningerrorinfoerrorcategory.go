package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningErrorInfoErrorCategory string

const (
	ProvisioningErrorInfoErrorCategoryfailure           ProvisioningErrorInfoErrorCategory = "Failure"
	ProvisioningErrorInfoErrorCategorynonServiceFailure ProvisioningErrorInfoErrorCategory = "NonServiceFailure"
	ProvisioningErrorInfoErrorCategorysuccess           ProvisioningErrorInfoErrorCategory = "Success"
)

func PossibleValuesForProvisioningErrorInfoErrorCategory() []string {
	return []string{
		string(ProvisioningErrorInfoErrorCategoryfailure),
		string(ProvisioningErrorInfoErrorCategorynonServiceFailure),
		string(ProvisioningErrorInfoErrorCategorysuccess),
	}
}

func parseProvisioningErrorInfoErrorCategory(input string) (*ProvisioningErrorInfoErrorCategory, error) {
	vals := map[string]ProvisioningErrorInfoErrorCategory{
		"failure":           ProvisioningErrorInfoErrorCategoryfailure,
		"nonservicefailure": ProvisioningErrorInfoErrorCategorynonServiceFailure,
		"success":           ProvisioningErrorInfoErrorCategorysuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningErrorInfoErrorCategory(input)
	return &out, nil
}
