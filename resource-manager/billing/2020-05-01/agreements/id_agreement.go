package agreements

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AgreementId{}

// AgreementId is a struct representing the Resource ID for a Agreement
type AgreementId struct {
	BillingAccountName string
	AgreementName      string
}

// NewAgreementID returns a new AgreementId struct
func NewAgreementID(billingAccountName string, agreementName string) AgreementId {
	return AgreementId{
		BillingAccountName: billingAccountName,
		AgreementName:      agreementName,
	}
}

// ParseAgreementID parses 'input' into a AgreementId
func ParseAgreementID(input string) (*AgreementId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AgreementId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AgreementId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAgreementIDInsensitively parses 'input' case-insensitively into a AgreementId
// note: this method should only be used for API response data and not user input
func ParseAgreementIDInsensitively(input string) (*AgreementId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AgreementId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AgreementId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AgreementId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.AgreementName, ok = input.Parsed["agreementName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "agreementName", input)
	}

	return nil
}

// ValidateAgreementID checks that 'input' can be parsed as a Agreement ID
func ValidateAgreementID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAgreementID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Agreement ID
func (id AgreementId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/agreements/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.AgreementName)
}

// Segments returns a slice of Resource ID Segments which comprise this Agreement ID
func (id AgreementId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticAgreements", "agreements", "agreements"),
		resourceids.UserSpecifiedSegment("agreementName", "agreementValue"),
	}
}

// String returns a human-readable description of this Agreement ID
func (id AgreementId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Agreement Name: %q", id.AgreementName),
	}
	return fmt.Sprintf("Agreement (%s)", strings.Join(components, "\n"))
}
