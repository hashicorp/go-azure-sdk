package workspacemanagedsqlserverserverencryptionprotector

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerKeyType string

const (
	ServerKeyTypeAzureKeyVault  ServerKeyType = "AzureKeyVault"
	ServerKeyTypeServiceManaged ServerKeyType = "ServiceManaged"
)

func PossibleValuesForServerKeyType() []string {
	return []string{
		string(ServerKeyTypeAzureKeyVault),
		string(ServerKeyTypeServiceManaged),
	}
}

func parseServerKeyType(input string) (*ServerKeyType, error) {
	vals := map[string]ServerKeyType{
		"azurekeyvault":  ServerKeyTypeAzureKeyVault,
		"servicemanaged": ServerKeyTypeServiceManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerKeyType(input)
	return &out, nil
}
