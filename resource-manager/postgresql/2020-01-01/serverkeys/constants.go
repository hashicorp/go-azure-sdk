package serverkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerKeyType string

const (
	ServerKeyTypeAzureKeyVault ServerKeyType = "AzureKeyVault"
)

func PossibleValuesForServerKeyType() []string {
	return []string{
		string(ServerKeyTypeAzureKeyVault),
	}
}
