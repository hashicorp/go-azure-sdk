package lots

import "strings"

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

func parseLotSource(input string) (*LotSource, error) {
	vals := map[string]LotSource{
		"consumptioncommitment": LotSourceConsumptionCommitment,
		"promotionalcredit":     LotSourcePromotionalCredit,
		"purchasedcredit":       LotSourcePurchasedCredit,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LotSource(input)
	return &out, nil
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

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"active":   StatusActive,
		"canceled": StatusCanceled,
		"complete": StatusComplete,
		"expired":  StatusExpired,
		"inactive": StatusInactive,
		"none":     StatusNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
