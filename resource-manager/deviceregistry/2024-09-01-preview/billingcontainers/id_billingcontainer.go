package billingcontainers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingContainerId{})
}

var _ resourceids.ResourceId = &BillingContainerId{}

// BillingContainerId is a struct representing the Resource ID for a Billing Container
type BillingContainerId struct {
	SubscriptionId       string
	BillingContainerName string
}

// NewBillingContainerID returns a new BillingContainerId struct
func NewBillingContainerID(subscriptionId string, billingContainerName string) BillingContainerId {
	return BillingContainerId{
		SubscriptionId:       subscriptionId,
		BillingContainerName: billingContainerName,
	}
}

// ParseBillingContainerID parses 'input' into a BillingContainerId
func ParseBillingContainerID(input string) (*BillingContainerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingContainerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingContainerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingContainerIDInsensitively parses 'input' case-insensitively into a BillingContainerId
// note: this method should only be used for API response data and not user input
func ParseBillingContainerIDInsensitively(input string) (*BillingContainerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingContainerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingContainerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingContainerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.BillingContainerName, ok = input.Parsed["billingContainerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingContainerName", input)
	}

	return nil
}

// ValidateBillingContainerID checks that 'input' can be parsed as a Billing Container ID
func ValidateBillingContainerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingContainerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Container ID
func (id BillingContainerId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.DeviceRegistry/billingContainers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.BillingContainerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Container ID
func (id BillingContainerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticBillingContainers", "billingContainers", "billingContainers"),
		resourceids.UserSpecifiedSegment("billingContainerName", "billingContainerName"),
	}
}

// String returns a human-readable description of this Billing Container ID
func (id BillingContainerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Billing Container Name: %q", id.BillingContainerName),
	}
	return fmt.Sprintf("Billing Container (%s)", strings.Join(components, "\n"))
}
