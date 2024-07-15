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
	recaser.RegisterResourceId(&DaprComponentResiliencyPolicyId{})
}

var _ resourceids.ResourceId = &DaprComponentResiliencyPolicyId{}

// DaprComponentResiliencyPolicyId is a struct representing the Resource ID for a Dapr Component Resiliency Policy
type DaprComponentResiliencyPolicyId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedEnvironmentName string
	DaprComponentName      string
	ResiliencyPolicyName   string
}

// NewDaprComponentResiliencyPolicyID returns a new DaprComponentResiliencyPolicyId struct
func NewDaprComponentResiliencyPolicyID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, daprComponentName string, resiliencyPolicyName string) DaprComponentResiliencyPolicyId {
	return DaprComponentResiliencyPolicyId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedEnvironmentName: managedEnvironmentName,
		DaprComponentName:      daprComponentName,
		ResiliencyPolicyName:   resiliencyPolicyName,
	}
}

// ParseDaprComponentResiliencyPolicyID parses 'input' into a DaprComponentResiliencyPolicyId
func ParseDaprComponentResiliencyPolicyID(input string) (*DaprComponentResiliencyPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DaprComponentResiliencyPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DaprComponentResiliencyPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDaprComponentResiliencyPolicyIDInsensitively parses 'input' case-insensitively into a DaprComponentResiliencyPolicyId
// note: this method should only be used for API response data and not user input
func ParseDaprComponentResiliencyPolicyIDInsensitively(input string) (*DaprComponentResiliencyPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DaprComponentResiliencyPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DaprComponentResiliencyPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DaprComponentResiliencyPolicyId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateDaprComponentResiliencyPolicyID checks that 'input' can be parsed as a Dapr Component Resiliency Policy ID
func ValidateDaprComponentResiliencyPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDaprComponentResiliencyPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dapr Component Resiliency Policy ID
func (id DaprComponentResiliencyPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/daprComponents/%s/resiliencyPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.DaprComponentName, id.ResiliencyPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dapr Component Resiliency Policy ID
func (id DaprComponentResiliencyPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentValue"),
		resourceids.StaticSegment("staticDaprComponents", "daprComponents", "daprComponents"),
		resourceids.UserSpecifiedSegment("daprComponentName", "daprComponentValue"),
		resourceids.StaticSegment("staticResiliencyPolicies", "resiliencyPolicies", "resiliencyPolicies"),
		resourceids.UserSpecifiedSegment("resiliencyPolicyName", "resiliencyPolicyValue"),
	}
}

// String returns a human-readable description of this Dapr Component Resiliency Policy ID
func (id DaprComponentResiliencyPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Dapr Component Name: %q", id.DaprComponentName),
		fmt.Sprintf("Resiliency Policy Name: %q", id.ResiliencyPolicyName),
	}
	return fmt.Sprintf("Dapr Component Resiliency Policy (%s)", strings.Join(components, "\n"))
}
