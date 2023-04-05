package lots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotSource string

const (
	LotSourceConsumptionCommitment LotSource = "ConsumptionCommitment"
	LotSourcePromotionalCredit     LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit       LotSource = "PurchasedCredit"
)

func PossibleValuesForLotSource() []string {
	return []string{
		string(LotSourceConsumptionCommitment),
		string(LotSourcePromotionalCredit),
		string(LotSourcePurchasedCredit),
	}
}

type Status string

const (
	StatusActive   Status = "Active"
	StatusCanceled Status = "Canceled"
	StatusComplete Status = "Complete"
	StatusExpired  Status = "Expired"
	StatusInactive Status = "Inactive"
	StatusNone     Status = "None"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusActive),
		string(StatusCanceled),
		string(StatusComplete),
		string(StatusExpired),
		string(StatusInactive),
		string(StatusNone),
	}
}
