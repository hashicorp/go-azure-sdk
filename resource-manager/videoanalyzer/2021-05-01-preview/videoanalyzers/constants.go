package videoanalyzers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountEncryptionKeyType string

const (
	AccountEncryptionKeyTypeCustomerKey AccountEncryptionKeyType = "CustomerKey"
	AccountEncryptionKeyTypeSystemKey   AccountEncryptionKeyType = "SystemKey"
)

func PossibleValuesForAccountEncryptionKeyType() []string {
	return []string{
		string(AccountEncryptionKeyTypeCustomerKey),
		string(AccountEncryptionKeyTypeSystemKey),
	}
}

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

type VideoAnalyzerEndpointType string

const (
	VideoAnalyzerEndpointTypeClientApi VideoAnalyzerEndpointType = "ClientApi"
)

func PossibleValuesForVideoAnalyzerEndpointType() []string {
	return []string{
		string(VideoAnalyzerEndpointTypeClientApi),
	}
}
