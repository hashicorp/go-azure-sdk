package daprcomponentresiliencypolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ResiliencyPolicyId{})
}

var _ resourceids.ResourceId = &ResiliencyPolicyId{}

// ResiliencyPolicyId is a struct representing the Resource ID for a Resiliency Policy
type ResiliencyPolicyId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedEnvironmentName string
	DaprComponentName      string
	ResiliencyPolicyName   string
}

// NewResiliencyPolicyID returns a new ResiliencyPolicyId struct
func NewResiliencyPolicyID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, daprComponentName string, resiliencyPolicyName string) ResiliencyPolicyId {
	return ResiliencyPolicyId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedEnvironmentName: managedEnvironmentName,
		DaprComponentName:      daprComponentName,
		ResiliencyPolicyName:   resiliencyPolicyName,
	}
}

// ParseResiliencyPolicyID parses 'input' into a ResiliencyPolicyId
func ParseResiliencyPolicyID(input string) (*ResiliencyPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResiliencyPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResiliencyPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResiliencyPolicyIDInsensitively parses 'input' case-insensitively into a ResiliencyPolicyId
// note: this method should only be used for API response data and not user input
func ParseResiliencyPolicyIDInsensitively(input string) (*ResiliencyPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResiliencyPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResiliencyPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResiliencyPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedEnvironmentName, ok = input.Parsed["managedEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedEnvironmentName", input)
	}

	if id.DaprComponentName, ok = input.Parsed["daprComponentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "daprComponentName", input)
	}

	if id.ResiliencyPolicyName, ok = input.Parsed["resiliencyPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resiliencyPolicyName", input)
	}

	return nil
}

// ValidateResiliencyPolicyID checks that 'input' can be parsed as a Resiliency Policy ID
func ValidateResiliencyPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResiliencyPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resiliency Policy ID
func (id ResiliencyPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/daprComponents/%s/resiliencyPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.DaprComponentName, id.ResiliencyPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resiliency Policy ID
func (id ResiliencyPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentName"),
		resourceids.StaticSegment("staticDaprComponents", "daprComponents", "daprComponents"),
		resourceids.UserSpecifiedSegment("daprComponentName", "daprComponentName"),
		resourceids.StaticSegment("staticResiliencyPolicies", "resiliencyPolicies", "resiliencyPolicies"),
		resourceids.UserSpecifiedSegment("resiliencyPolicyName", "resiliencyPolicyName"),
	}
}

// String returns a human-readable description of this Resiliency Policy ID
func (id ResiliencyPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Dapr Component Name: %q", id.DaprComponentName),
		fmt.Sprintf("Resiliency Policy Name: %q", id.ResiliencyPolicyName),
	}
	return fmt.Sprintf("Resiliency Policy (%s)", strings.Join(components, "\n"))
}
