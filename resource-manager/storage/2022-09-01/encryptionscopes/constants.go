package encryptionscopes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionScopeSource string

const (
	EncryptionScopeSourceMicrosoftPointKeyVault EncryptionScopeSource = "Microsoft.KeyVault"
	EncryptionScopeSourceMicrosoftPointStorage  EncryptionScopeSource = "Microsoft.Storage"
)

func PossibleValuesForEncryptionScopeSource() []string {
	return []string{
		string(EncryptionScopeSourceMicrosoftPointKeyVault),
		string(EncryptionScopeSourceMicrosoftPointStorage),
	}
}

type EncryptionScopeState string

const (
	EncryptionScopeStateDisabled EncryptionScopeState = "Disabled"
	EncryptionScopeStateEnabled  EncryptionScopeState = "Enabled"
)

func PossibleValuesForEncryptionScopeState() []string {
	return []string{
		string(EncryptionScopeStateDisabled),
		string(EncryptionScopeStateEnabled),
	}
}

type ListEncryptionScopesInclude string

const (
	ListEncryptionScopesIncludeAll      ListEncryptionScopesInclude = "All"
	ListEncryptionScopesIncludeDisabled ListEncryptionScopesInclude = "Disabled"
	ListEncryptionScopesIncludeEnabled  ListEncryptionScopesInclude = "Enabled"
)

func PossibleValuesForListEncryptionScopesInclude() []string {
	return []string{
		string(ListEncryptionScopesIncludeAll),
		string(ListEncryptionScopesIncludeDisabled),
		string(ListEncryptionScopesIncludeEnabled),
	}
}
