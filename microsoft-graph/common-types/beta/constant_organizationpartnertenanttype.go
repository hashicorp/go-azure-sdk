package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationPartnerTenantType string

const (
	OrganizationPartnerTenantType_BreadthPartner                          OrganizationPartnerTenantType = "breadthPartner"
	OrganizationPartnerTenantType_BreadthPartnerDelegatedAdmin            OrganizationPartnerTenantType = "breadthPartnerDelegatedAdmin"
	OrganizationPartnerTenantType_MicrosoftSupport                        OrganizationPartnerTenantType = "microsoftSupport"
	OrganizationPartnerTenantType_ResellerPartnerDelegatedAdmin           OrganizationPartnerTenantType = "resellerPartnerDelegatedAdmin"
	OrganizationPartnerTenantType_SyndicatePartner                        OrganizationPartnerTenantType = "syndicatePartner"
	OrganizationPartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin OrganizationPartnerTenantType = "valueAddedResellerPartnerDelegatedAdmin"
)

func PossibleValuesForOrganizationPartnerTenantType() []string {
	return []string{
		string(OrganizationPartnerTenantType_BreadthPartner),
		string(OrganizationPartnerTenantType_BreadthPartnerDelegatedAdmin),
		string(OrganizationPartnerTenantType_MicrosoftSupport),
		string(OrganizationPartnerTenantType_ResellerPartnerDelegatedAdmin),
		string(OrganizationPartnerTenantType_SyndicatePartner),
		string(OrganizationPartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin),
	}
}

func (s *OrganizationPartnerTenantType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrganizationPartnerTenantType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrganizationPartnerTenantType(input string) (*OrganizationPartnerTenantType, error) {
	vals := map[string]OrganizationPartnerTenantType{
		"breadthpartner":                          OrganizationPartnerTenantType_BreadthPartner,
		"breadthpartnerdelegatedadmin":            OrganizationPartnerTenantType_BreadthPartnerDelegatedAdmin,
		"microsoftsupport":                        OrganizationPartnerTenantType_MicrosoftSupport,
		"resellerpartnerdelegatedadmin":           OrganizationPartnerTenantType_ResellerPartnerDelegatedAdmin,
		"syndicatepartner":                        OrganizationPartnerTenantType_SyndicatePartner,
		"valueaddedresellerpartnerdelegatedadmin": OrganizationPartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationPartnerTenantType(input)
	return &out, nil
}
