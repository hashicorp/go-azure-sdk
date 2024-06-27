package associatedtenant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AssociatedTenantId{})
}

var _ resourceids.ResourceId = &AssociatedTenantId{}

// AssociatedTenantId is a struct representing the Resource ID for a Associated Tenant
type AssociatedTenantId struct {
	BillingAccountName   string
	AssociatedTenantName string
}

// NewAssociatedTenantID returns a new AssociatedTenantId struct
func NewAssociatedTenantID(billingAccountName string, associatedTenantName string) AssociatedTenantId {
	return AssociatedTenantId{
		BillingAccountName:   billingAccountName,
		AssociatedTenantName: associatedTenantName,
	}
}

// ParseAssociatedTenantID parses 'input' into a AssociatedTenantId
func ParseAssociatedTenantID(input string) (*AssociatedTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AssociatedTenantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AssociatedTenantId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAssociatedTenantIDInsensitively parses 'input' case-insensitively into a AssociatedTenantId
// note: this method should only be used for API response data and not user input
func ParseAssociatedTenantIDInsensitively(input string) (*AssociatedTenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AssociatedTenantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AssociatedTenantId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AssociatedTenantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.AssociatedTenantName, ok = input.Parsed["associatedTenantName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "associatedTenantName", input)
	}

	return nil
}

// ValidateAssociatedTenantID checks that 'input' can be parsed as a Associated Tenant ID
func ValidateAssociatedTenantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAssociatedTenantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Associated Tenant ID
func (id AssociatedTenantId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/associatedTenants/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.AssociatedTenantName)
}

// Segments returns a slice of Resource ID Segments which comprise this Associated Tenant ID
func (id AssociatedTenantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticAssociatedTenants", "associatedTenants", "associatedTenants"),
		resourceids.UserSpecifiedSegment("associatedTenantName", "associatedTenantValue"),
	}
}

// String returns a human-readable description of this Associated Tenant ID
func (id AssociatedTenantId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Associated Tenant Name: %q", id.AssociatedTenantName),
	}
	return fmt.Sprintf("Associated Tenant (%s)", strings.Join(components, "\n"))
}
