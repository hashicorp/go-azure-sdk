package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RackSkuId{})
}

var _ resourceids.ResourceId = &RackSkuId{}

// RackSkuId is a struct representing the Resource ID for a Rack Sku
type RackSkuId struct {
	SubscriptionId string
	RackSkuName    string
}

// NewRackSkuID returns a new RackSkuId struct
func NewRackSkuID(subscriptionId string, rackSkuName string) RackSkuId {
	return RackSkuId{
		SubscriptionId: subscriptionId,
		RackSkuName:    rackSkuName,
	}
}

// ParseRackSkuID parses 'input' into a RackSkuId
func ParseRackSkuID(input string) (*RackSkuId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RackSkuId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RackSkuId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRackSkuIDInsensitively parses 'input' case-insensitively into a RackSkuId
// note: this method should only be used for API response data and not user input
func ParseRackSkuIDInsensitively(input string) (*RackSkuId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RackSkuId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RackSkuId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RackSkuId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.RackSkuName, ok = input.Parsed["rackSkuName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rackSkuName", input)
	}

	return nil
}

// ValidateRackSkuID checks that 'input' can be parsed as a Rack Sku ID
func ValidateRackSkuID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRackSkuID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rack Sku ID
func (id RackSkuId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.NetworkCloud/rackSkus/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RackSkuName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rack Sku ID
func (id RackSkuId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticRackSkus", "rackSkus", "rackSkus"),
		resourceids.UserSpecifiedSegment("rackSkuName", "rackSkuValue"),
	}
}

// String returns a human-readable description of this Rack Sku ID
func (id RackSkuId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Rack Sku Name: %q", id.RackSkuName),
	}
	return fmt.Sprintf("Rack Sku (%s)", strings.Join(components, "\n"))
}
