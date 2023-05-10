package jitrequests

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = JitRequestId{}

// JitRequestId is a struct representing the Resource ID for a Jit Request
type JitRequestId struct {
	SubscriptionId    string
	ResourceGroupName string
	JitRequestName    string
}

// NewJitRequestID returns a new JitRequestId struct
func NewJitRequestID(subscriptionId string, resourceGroupName string, jitRequestName string) JitRequestId {
	return JitRequestId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		JitRequestName:    jitRequestName,
	}
}

// ParseJitRequestID parses 'input' into a JitRequestId
func ParseJitRequestID(input string) (*JitRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(JitRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JitRequestId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.JitRequestName, ok = parsed.Parsed["jitRequestName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jitRequestName", *parsed)
	}

	return &id, nil
}

// ParseJitRequestIDInsensitively parses 'input' case-insensitively into a JitRequestId
// note: this method should only be used for API response data and not user input
func ParseJitRequestIDInsensitively(input string) (*JitRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(JitRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JitRequestId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.JitRequestName, ok = parsed.Parsed["jitRequestName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jitRequestName", *parsed)
	}

	return &id, nil
}

// ValidateJitRequestID checks that 'input' can be parsed as a Jit Request ID
func ValidateJitRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseJitRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Jit Request ID
func (id JitRequestId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Solutions/jitRequests/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.JitRequestName)
}

// Segments returns a slice of Resource ID Segments which comprise this Jit Request ID
func (id JitRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSolutions", "Microsoft.Solutions", "Microsoft.Solutions"),
		resourceids.StaticSegment("staticJitRequests", "jitRequests", "jitRequests"),
		resourceids.UserSpecifiedSegment("jitRequestName", "jitRequestValue"),
	}
}

// String returns a human-readable description of this Jit Request ID
func (id JitRequestId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Jit Request Name: %q", id.JitRequestName),
	}
	return fmt.Sprintf("Jit Request (%s)", strings.Join(components, "\n"))
}
