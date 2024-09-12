package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementAcceptanceId{}

// IdentityGovernanceTermsOfUseAgreementAcceptanceId is a struct representing the Resource ID for a Identity Governance Terms Of Use Agreement Acceptance
type IdentityGovernanceTermsOfUseAgreementAcceptanceId struct {
	AgreementAcceptanceId string
}

// NewIdentityGovernanceTermsOfUseAgreementAcceptanceID returns a new IdentityGovernanceTermsOfUseAgreementAcceptanceId struct
func NewIdentityGovernanceTermsOfUseAgreementAcceptanceID(agreementAcceptanceId string) IdentityGovernanceTermsOfUseAgreementAcceptanceId {
	return IdentityGovernanceTermsOfUseAgreementAcceptanceId{
		AgreementAcceptanceId: agreementAcceptanceId,
	}
}

// ParseIdentityGovernanceTermsOfUseAgreementAcceptanceID parses 'input' into a IdentityGovernanceTermsOfUseAgreementAcceptanceId
func ParseIdentityGovernanceTermsOfUseAgreementAcceptanceID(input string) (*IdentityGovernanceTermsOfUseAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceTermsOfUseAgreementAcceptanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceTermsOfUseAgreementAcceptanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceTermsOfUseAgreementAcceptanceIDInsensitively(input string) (*IdentityGovernanceTermsOfUseAgreementAcceptanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceTermsOfUseAgreementAcceptanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceTermsOfUseAgreementAcceptanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceTermsOfUseAgreementAcceptanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AgreementAcceptanceId, ok = input.Parsed["agreementAcceptanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementAcceptanceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceTermsOfUseAgreementAcceptanceID checks that 'input' can be parsed as a Identity Governance Terms Of Use Agreement Acceptance ID
func ValidateIdentityGovernanceTermsOfUseAgreementAcceptanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceTermsOfUseAgreementAcceptanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Terms Of Use Agreement Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementAcceptanceId) ID() string {
	fmtString := "/identityGovernance/termsOfUse/agreementAcceptances/%s"
	return fmt.Sprintf(fmtString, id.AgreementAcceptanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Terms Of Use Agreement Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementAcceptanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("termsOfUse", "termsOfUse", "termsOfUse"),
		resourceids.StaticSegment("agreementAcceptances", "agreementAcceptances", "agreementAcceptances"),
		resourceids.UserSpecifiedSegment("agreementAcceptanceId", "agreementAcceptanceIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Terms Of Use Agreement Acceptance ID
func (id IdentityGovernanceTermsOfUseAgreementAcceptanceId) String() string {
	components := []string{
		fmt.Sprintf("Agreement Acceptance: %q", id.AgreementAcceptanceId),
	}
	return fmt.Sprintf("Identity Governance Terms Of Use Agreement Acceptance (%s)", strings.Join(components, "\n"))
}
