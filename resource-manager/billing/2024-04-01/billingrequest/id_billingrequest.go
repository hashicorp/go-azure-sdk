package billingrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingRequestId{})
}

var _ resourceids.ResourceId = &BillingRequestId{}

// BillingRequestId is a struct representing the Resource ID for a Billing Request
type BillingRequestId struct {
	BillingRequestName string
}

// NewBillingRequestID returns a new BillingRequestId struct
func NewBillingRequestID(billingRequestName string) BillingRequestId {
	return BillingRequestId{
		BillingRequestName: billingRequestName,
	}
}

// ParseBillingRequestID parses 'input' into a BillingRequestId
func ParseBillingRequestID(input string) (*BillingRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingRequestIDInsensitively parses 'input' case-insensitively into a BillingRequestId
// note: this method should only be used for API response data and not user input
func ParseBillingRequestIDInsensitively(input string) (*BillingRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingRequestName, ok = input.Parsed["billingRequestName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRequestName", input)
	}

	return nil
}

// ValidateBillingRequestID checks that 'input' can be parsed as a Billing Request ID
func ValidateBillingRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Request ID
func (id BillingRequestId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingRequests/%s"
	return fmt.Sprintf(fmtString, id.BillingRequestName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Request ID
func (id BillingRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingRequests", "billingRequests", "billingRequests"),
		resourceids.UserSpecifiedSegment("billingRequestName", "billingRequestValue"),
	}
}

// String returns a human-readable description of this Billing Request ID
func (id BillingRequestId) String() string {
	components := []string{
		fmt.Sprintf("Billing Request Name: %q", id.BillingRequestName),
	}
	return fmt.Sprintf("Billing Request (%s)", strings.Join(components, "\n"))
}
