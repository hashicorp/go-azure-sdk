package netappaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryStatus string

const (
	ActiveDirectoryStatusCreated  ActiveDirectoryStatus = "Created"
	ActiveDirectoryStatusDeleted  ActiveDirectoryStatus = "Deleted"
	ActiveDirectoryStatusError    ActiveDirectoryStatus = "Error"
	ActiveDirectoryStatusInUse    ActiveDirectoryStatus = "InUse"
	ActiveDirectoryStatusUpdating ActiveDirectoryStatus = "Updating"
)

func PossibleValuesForActiveDirectoryStatus() []string {
	return []string{
		string(ActiveDirectoryStatusCreated),
		string(ActiveDirectoryStatusDeleted),
		string(ActiveDirectoryStatusError),
		string(ActiveDirectoryStatusInUse),
		string(ActiveDirectoryStatusUpdating),
	}
}

type KeySource string

const (
	KeySourceMicrosoftPointKeyVault KeySource = "Microsoft.KeyVault"
	KeySourceMicrosoftPointNetApp   KeySource = "Microsoft.NetApp"
)

func PossibleValuesForKeySource() []string {
	return []string{
		string(KeySourceMicrosoftPointKeyVault),
		string(KeySourceMicrosoftPointNetApp),
	}
}

type KeyVaultStatus string

const (
	KeyVaultStatusCreated  KeyVaultStatus = "Created"
	KeyVaultStatusDeleted  KeyVaultStatus = "Deleted"
	KeyVaultStatusError    KeyVaultStatus = "Error"
	KeyVaultStatusInUse    KeyVaultStatus = "InUse"
	KeyVaultStatusUpdating KeyVaultStatus = "Updating"
)

func PossibleValuesForKeyVaultStatus() []string {
	return []string{
		string(KeyVaultStatusCreated),
		string(KeyVaultStatusDeleted),
		string(KeyVaultStatusError),
		string(KeyVaultStatusInUse),
		string(KeyVaultStatusUpdating),
	}
}
