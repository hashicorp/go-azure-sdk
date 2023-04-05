package dataexport

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Type string

const (
	TypeEventHub       Type = "EventHub"
	TypeStorageAccount Type = "StorageAccount"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeEventHub),
		string(TypeStorageAccount),
	}
}
