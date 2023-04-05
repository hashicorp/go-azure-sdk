package providers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AliasPathAttributes string

const (
	AliasPathAttributesModifiable AliasPathAttributes = "Modifiable"
	AliasPathAttributesNone       AliasPathAttributes = "None"
)

func PossibleValuesForAliasPathAttributes() []string {
	return []string{
		string(AliasPathAttributesModifiable),
		string(AliasPathAttributesNone),
	}
}

type AliasPathTokenType string

const (
	AliasPathTokenTypeAny          AliasPathTokenType = "Any"
	AliasPathTokenTypeArray        AliasPathTokenType = "Array"
	AliasPathTokenTypeBoolean      AliasPathTokenType = "Boolean"
	AliasPathTokenTypeInteger      AliasPathTokenType = "Integer"
	AliasPathTokenTypeNotSpecified AliasPathTokenType = "NotSpecified"
	AliasPathTokenTypeNumber       AliasPathTokenType = "Number"
	AliasPathTokenTypeObject       AliasPathTokenType = "Object"
	AliasPathTokenTypeString       AliasPathTokenType = "String"
)

func PossibleValuesForAliasPathTokenType() []string {
	return []string{
		string(AliasPathTokenTypeAny),
		string(AliasPathTokenTypeArray),
		string(AliasPathTokenTypeBoolean),
		string(AliasPathTokenTypeInteger),
		string(AliasPathTokenTypeNotSpecified),
		string(AliasPathTokenTypeNumber),
		string(AliasPathTokenTypeObject),
		string(AliasPathTokenTypeString),
	}
}

type AliasPatternType string

const (
	AliasPatternTypeExtract      AliasPatternType = "Extract"
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
)

func PossibleValuesForAliasPatternType() []string {
	return []string{
		string(AliasPatternTypeExtract),
		string(AliasPatternTypeNotSpecified),
	}
}

type AliasType string

const (
	AliasTypeMask         AliasType = "Mask"
	AliasTypeNotSpecified AliasType = "NotSpecified"
	AliasTypePlainText    AliasType = "PlainText"
)

func PossibleValuesForAliasType() []string {
	return []string{
		string(AliasTypeMask),
		string(AliasTypeNotSpecified),
		string(AliasTypePlainText),
	}
}

type ProviderAuthorizationConsentState string

const (
	ProviderAuthorizationConsentStateConsented    ProviderAuthorizationConsentState = "Consented"
	ProviderAuthorizationConsentStateNotRequired  ProviderAuthorizationConsentState = "NotRequired"
	ProviderAuthorizationConsentStateNotSpecified ProviderAuthorizationConsentState = "NotSpecified"
	ProviderAuthorizationConsentStateRequired     ProviderAuthorizationConsentState = "Required"
)

func PossibleValuesForProviderAuthorizationConsentState() []string {
	return []string{
		string(ProviderAuthorizationConsentStateConsented),
		string(ProviderAuthorizationConsentStateNotRequired),
		string(ProviderAuthorizationConsentStateNotSpecified),
		string(ProviderAuthorizationConsentStateRequired),
	}
}
