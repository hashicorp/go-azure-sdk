package marketplaceregistrationdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure MultiFactorAuthProvider = "Azure"
	MultiFactorAuthProviderNone  MultiFactorAuthProvider = "None"
)

func PossibleValuesForMultiFactorAuthProvider() []string {
	return []string{
		string(MultiFactorAuthProviderAzure),
		string(MultiFactorAuthProviderNone),
	}
}
