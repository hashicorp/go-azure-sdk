package skus

import "strings"

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

func parseReasonCode(input string) (*ReasonCode, error) {
	vals := map[string]ReasonCode{
		"notavailableforsubscription": ReasonCodeNotAvailableForSubscription,
		"quotaid":                     ReasonCodeQuotaId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReasonCode(input)
	return &out, nil
}
