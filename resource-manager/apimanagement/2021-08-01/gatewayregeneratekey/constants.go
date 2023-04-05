package gatewayregeneratekey

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyType string

const (
	KeyTypePrimary   KeyType = "primary"
	KeyTypeSecondary KeyType = "secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}
