package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaId                     ReasonCode = "QuotaId"
)

func PossibleValuesForReasonCode() []string {
	return []string{
		string(ReasonCodeNotAvailableForSubscription),
		string(ReasonCodeQuotaId),
	}
}
