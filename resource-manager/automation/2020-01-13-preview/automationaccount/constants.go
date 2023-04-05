package automationaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationAccountState string

const (
	AutomationAccountStateOk          AutomationAccountState = "Ok"
	AutomationAccountStateSuspended   AutomationAccountState = "Suspended"
	AutomationAccountStateUnavailable AutomationAccountState = "Unavailable"
)

func PossibleValuesForAutomationAccountState() []string {
	return []string{
		string(AutomationAccountStateOk),
		string(AutomationAccountStateSuspended),
		string(AutomationAccountStateUnavailable),
	}
}

type EncryptionKeySourceType string

const (
	EncryptionKeySourceTypeMicrosoftPointAutomation EncryptionKeySourceType = "Microsoft.Automation"
	EncryptionKeySourceTypeMicrosoftPointKeyvault   EncryptionKeySourceType = "Microsoft.Keyvault"
)

func PossibleValuesForEncryptionKeySourceType() []string {
	return []string{
		string(EncryptionKeySourceTypeMicrosoftPointAutomation),
		string(EncryptionKeySourceTypeMicrosoftPointKeyvault),
	}
}

type SkuNameEnum string

const (
	SkuNameEnumBasic SkuNameEnum = "Basic"
	SkuNameEnumFree  SkuNameEnum = "Free"
)

func PossibleValuesForSkuNameEnum() []string {
	return []string{
		string(SkuNameEnumBasic),
		string(SkuNameEnumFree),
	}
}
