package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInformationCompanyType string

const (
	PartnerInformationCompanyType_BreadthPartner                          PartnerInformationCompanyType = "breadthPartner"
	PartnerInformationCompanyType_BreadthPartnerDelegatedAdmin            PartnerInformationCompanyType = "breadthPartnerDelegatedAdmin"
	PartnerInformationCompanyType_MicrosoftSupport                        PartnerInformationCompanyType = "microsoftSupport"
	PartnerInformationCompanyType_ResellerPartnerDelegatedAdmin           PartnerInformationCompanyType = "resellerPartnerDelegatedAdmin"
	PartnerInformationCompanyType_SyndicatePartner                        PartnerInformationCompanyType = "syndicatePartner"
	PartnerInformationCompanyType_ValueAddedResellerPartnerDelegatedAdmin PartnerInformationCompanyType = "valueAddedResellerPartnerDelegatedAdmin"
)

func PossibleValuesForPartnerInformationCompanyType() []string {
	return []string{
		string(PartnerInformationCompanyType_BreadthPartner),
		string(PartnerInformationCompanyType_BreadthPartnerDelegatedAdmin),
		string(PartnerInformationCompanyType_MicrosoftSupport),
		string(PartnerInformationCompanyType_ResellerPartnerDelegatedAdmin),
		string(PartnerInformationCompanyType_SyndicatePartner),
		string(PartnerInformationCompanyType_ValueAddedResellerPartnerDelegatedAdmin),
	}
}

func (s *PartnerInformationCompanyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerInformationCompanyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerInformationCompanyType(input string) (*PartnerInformationCompanyType, error) {
	vals := map[string]PartnerInformationCompanyType{
		"breadthpartner":                          PartnerInformationCompanyType_BreadthPartner,
		"breadthpartnerdelegatedadmin":            PartnerInformationCompanyType_BreadthPartnerDelegatedAdmin,
		"microsoftsupport":                        PartnerInformationCompanyType_MicrosoftSupport,
		"resellerpartnerdelegatedadmin":           PartnerInformationCompanyType_ResellerPartnerDelegatedAdmin,
		"syndicatepartner":                        PartnerInformationCompanyType_SyndicatePartner,
		"valueaddedresellerpartnerdelegatedadmin": PartnerInformationCompanyType_ValueAddedResellerPartnerDelegatedAdmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerInformationCompanyType(input)
	return &out, nil
}
