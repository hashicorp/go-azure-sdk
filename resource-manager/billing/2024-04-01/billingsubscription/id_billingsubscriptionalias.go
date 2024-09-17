package billingsubscription

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingSubscriptionAliasId{})
}

var _ resourceids.ResourceId = &BillingSubscriptionAliasId{}

// BillingSubscriptionAliasId is a struct representing the Resource ID for a Billing Subscription Alias
type BillingSubscriptionAliasId struct {
	BillingAccountName           string
	BillingSubscriptionAliasName string
}

// NewBillingSubscriptionAliasID returns a new BillingSubscriptionAliasId struct
func NewBillingSubscriptionAliasID(billingAccountName string, billingSubscriptionAliasName string) BillingSubscriptionAliasId {
	return BillingSubscriptionAliasId{
		BillingAccountName:           billingAccountName,
		BillingSubscriptionAliasName: billingSubscriptionAliasName,
	}
}

// ParseBillingSubscriptionAliasID parses 'input' into a BillingSubscriptionAliasId
func ParseBillingSubscriptionAliasID(input string) (*BillingSubscriptionAliasId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionAliasId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionAliasId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingSubscriptionAliasIDInsensitively parses 'input' case-insensitively into a BillingSubscriptionAliasId
// note: this method should only be used for API response data and not user input
func ParseBillingSubscriptionAliasIDInsensitively(input string) (*BillingSubscriptionAliasId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionAliasId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionAliasId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingSubscriptionAliasId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingSubscriptionAliasName, ok = input.Parsed["billingSubscriptionAliasName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionAliasName", input)
	}

	return nil
}

// ValidateBillingSubscriptionAliasID checks that 'input' can be parsed as a Billing Subscription Alias ID
func ValidateBillingSubscriptionAliasID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingSubscriptionAliasID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Subscription Alias ID
func (id BillingSubscriptionAliasId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingSubscriptionAliases/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingSubscriptionAliasName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Subscription Alias ID
func (id BillingSubscriptionAliasId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingSubscriptionAliases", "billingSubscriptionAliases", "billingSubscriptionAliases"),
		resourceids.UserSpecifiedSegment("billingSubscriptionAliasName", "billingSubscriptionAliasValue"),
	}
}

// String returns a human-readable description of this Billing Subscription Alias ID
func (id BillingSubscriptionAliasId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Subscription Alias Name: %q", id.BillingSubscriptionAliasName),
	}
	return fmt.Sprintf("Billing Subscription Alias (%s)", strings.Join(components, "\n"))
}
