package agreements

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcceptanceMode string

const (
	AcceptanceModeClickToAccept AcceptanceMode = "ClickToAccept"
	AcceptanceModeESignEmbedded AcceptanceMode = "ESignEmbedded"
	AcceptanceModeESignOffline  AcceptanceMode = "ESignOffline"
)

func PossibleValuesForAcceptanceMode() []string {
	return []string{
		string(AcceptanceModeClickToAccept),
		string(AcceptanceModeESignEmbedded),
		string(AcceptanceModeESignOffline),
	}
}

func (s *AcceptanceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAcceptanceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAcceptanceMode(input string) (*AcceptanceMode, error) {
	vals := map[string]AcceptanceMode{
		"clicktoaccept": AcceptanceModeClickToAccept,
		"esignembedded": AcceptanceModeESignEmbedded,
		"esignoffline":  AcceptanceModeESignOffline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AcceptanceMode(input)
	return &out, nil
}

type Category string

const (
	CategoryAffiliatePurchaseTerms     Category = "AffiliatePurchaseTerms"
	CategoryMicrosoftCustomerAgreement Category = "MicrosoftCustomerAgreement"
	CategoryOther                      Category = "Other"
)

func PossibleValuesForCategory() []string {
	return []string{
		string(CategoryAffiliatePurchaseTerms),
		string(CategoryMicrosoftCustomerAgreement),
		string(CategoryOther),
	}
}

func (s *Category) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategory(input string) (*Category, error) {
	vals := map[string]Category{
		"affiliatepurchaseterms":     CategoryAffiliatePurchaseTerms,
		"microsoftcustomeragreement": CategoryMicrosoftCustomerAgreement,
		"other":                      CategoryOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Category(input)
	return &out, nil
}
