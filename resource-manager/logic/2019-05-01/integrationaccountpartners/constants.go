package integrationaccountpartners

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyType string

const (
	KeyTypeNotSpecified KeyType = "NotSpecified"
	KeyTypePrimary      KeyType = "Primary"
	KeyTypeSecondary    KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeNotSpecified),
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

type PartnerType string

const (
	PartnerTypeBTwoB        PartnerType = "B2B"
	PartnerTypeNotSpecified PartnerType = "NotSpecified"
)

func PossibleValuesForPartnerType() []string {
	return []string{
		string(PartnerTypeBTwoB),
		string(PartnerTypeNotSpecified),
	}
}
