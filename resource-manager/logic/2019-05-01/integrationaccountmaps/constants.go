package integrationaccountmaps

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

type MapType string

const (
	MapTypeLiquid        MapType = "Liquid"
	MapTypeNotSpecified  MapType = "NotSpecified"
	MapTypeXslt          MapType = "Xslt"
	MapTypeXsltThreeZero MapType = "Xslt30"
	MapTypeXsltTwoZero   MapType = "Xslt20"
)

func PossibleValuesForMapType() []string {
	return []string{
		string(MapTypeLiquid),
		string(MapTypeNotSpecified),
		string(MapTypeXslt),
		string(MapTypeXsltThreeZero),
		string(MapTypeXsltTwoZero),
	}
}
