package backupprotectableitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureFileShareType string

const (
	AzureFileShareTypeInvalid AzureFileShareType = "Invalid"
	AzureFileShareTypeXSMB    AzureFileShareType = "XSMB"
	AzureFileShareTypeXSync   AzureFileShareType = "XSync"
)

func PossibleValuesForAzureFileShareType() []string {
	return []string{
		string(AzureFileShareTypeInvalid),
		string(AzureFileShareTypeXSMB),
		string(AzureFileShareTypeXSync),
	}
}

type InquiryStatus string

const (
	InquiryStatusFailed  InquiryStatus = "Failed"
	InquiryStatusInvalid InquiryStatus = "Invalid"
	InquiryStatusSuccess InquiryStatus = "Success"
)

func PossibleValuesForInquiryStatus() []string {
	return []string{
		string(InquiryStatusFailed),
		string(InquiryStatusInvalid),
		string(InquiryStatusSuccess),
	}
}

type ProtectionStatus string

const (
	ProtectionStatusInvalid          ProtectionStatus = "Invalid"
	ProtectionStatusNotProtected     ProtectionStatus = "NotProtected"
	ProtectionStatusProtected        ProtectionStatus = "Protected"
	ProtectionStatusProtecting       ProtectionStatus = "Protecting"
	ProtectionStatusProtectionFailed ProtectionStatus = "ProtectionFailed"
)

func PossibleValuesForProtectionStatus() []string {
	return []string{
		string(ProtectionStatusInvalid),
		string(ProtectionStatusNotProtected),
		string(ProtectionStatusProtected),
		string(ProtectionStatusProtecting),
		string(ProtectionStatusProtectionFailed),
	}
}
