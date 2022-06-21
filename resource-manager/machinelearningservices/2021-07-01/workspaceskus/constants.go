package workspaceskus

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReasonCode string

const (
	ReasonCodeNotAvailableForRegion       ReasonCode = "NotAvailableForRegion"
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeNotSpecified                ReasonCode = "NotSpecified"
)

func PossibleValuesForReasonCode() []string {
	return []string{
		string(ReasonCodeNotAvailableForRegion),
		string(ReasonCodeNotAvailableForSubscription),
		string(ReasonCodeNotSpecified),
	}
}

func parseReasonCode(input string) (*ReasonCode, error) {
	vals := map[string]ReasonCode{
		"notavailableforregion":       ReasonCodeNotAvailableForRegion,
		"notavailableforsubscription": ReasonCodeNotAvailableForSubscription,
		"notspecified":                ReasonCodeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReasonCode(input)
	return &out, nil
}
