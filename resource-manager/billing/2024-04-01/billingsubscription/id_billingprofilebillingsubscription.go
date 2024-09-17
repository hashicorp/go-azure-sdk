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
	recaser.RegisterResourceId(&BillingProfileBillingSubscriptionId{})
}

var _ resourceids.ResourceId = &BillingProfileBillingSubscriptionId{}

// BillingProfileBillingSubscriptionId is a struct representing the Resource ID for a Billing Profile Billing Subscription
type BillingProfileBillingSubscriptionId struct {
	BillingAccountName      string
	BillingProfileName      string
	BillingSubscriptionName string
}

// NewBillingProfileBillingSubscriptionID returns a new BillingProfileBillingSubscriptionId struct
func NewBillingProfileBillingSubscriptionID(billingAccountName string, billingProfileName string, billingSubscriptionName string) BillingProfileBillingSubscriptionId {
	return BillingProfileBillingSubscriptionId{
		BillingAccountName:      billingAccountName,
		BillingProfileName:      billingProfileName,
		BillingSubscriptionName: billingSubscriptionName,
	}
}

// ParseBillingProfileBillingSubscriptionID parses 'input' into a BillingProfileBillingSubscriptionId
func ParseBillingProfileBillingSubscriptionID(input string) (*BillingProfileBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingProfileBillingSubscriptionIDInsensitively parses 'input' case-insensitively into a BillingProfileBillingSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseBillingProfileBillingSubscriptionIDInsensitively(input string) (*BillingProfileBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingProfileBillingSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.BillingSubscriptionName, ok = input.Parsed["billingSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionName", input)
	}

	return nil
}

// ValidateBillingProfileBillingSubscriptionID checks that 'input' can be parsed as a Billing Profile Billing Subscription ID
func ValidateBillingProfileBillingSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingProfileBillingSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Profile Billing Subscription ID
func (id BillingProfileBillingSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/billingSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.BillingSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Profile Billing Subscription ID
func (id BillingProfileBillingSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("billingSubscriptionName", "billingSubscriptionValue"),
	}
}

// String returns a human-readable description of this Billing Profile Billing Subscription ID
func (id BillingProfileBillingSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Billing Subscription Name: %q", id.BillingSubscriptionName),
	}
	return fmt.Sprintf("Billing Profile Billing Subscription (%s)", strings.Join(components, "\n"))
}
