package jitrequests

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&JitRequestId{})
}

var _ resourceids.ResourceId = &JitRequestId{}

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
	parser := resourceids.NewParserFromResourceIdType(&JitRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := JitRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseJitRequestIDInsensitively parses 'input' case-insensitively into a JitRequestId
// note: this method should only be used for API response data and not user input
func ParseJitRequestIDInsensitively(input string) (*JitRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&JitRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := JitRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *JitRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.JitRequestName, ok = input.Parsed["jitRequestName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jitRequestName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("jitRequestName", "jitRequestName"),
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
