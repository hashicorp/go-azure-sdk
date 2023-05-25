package distributedavailabilitygroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DistributedAvailabilityGroupId{}

// DistributedAvailabilityGroupId is a struct representing the Resource ID for a Distributed Availability Group
type DistributedAvailabilityGroupId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	ManagedInstanceName              string
	DistributedAvailabilityGroupName string
}

// NewDistributedAvailabilityGroupID returns a new DistributedAvailabilityGroupId struct
func NewDistributedAvailabilityGroupID(subscriptionId string, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string) DistributedAvailabilityGroupId {
	return DistributedAvailabilityGroupId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		ManagedInstanceName:              managedInstanceName,
		DistributedAvailabilityGroupName: distributedAvailabilityGroupName,
	}
}

// ParseDistributedAvailabilityGroupID parses 'input' into a DistributedAvailabilityGroupId
func ParseDistributedAvailabilityGroupID(input string) (*DistributedAvailabilityGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(DistributedAvailabilityGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DistributedAvailabilityGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.DistributedAvailabilityGroupName, ok = parsed.Parsed["distributedAvailabilityGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "distributedAvailabilityGroupName", *parsed)
	}

	return &id, nil
}

// ParseDistributedAvailabilityGroupIDInsensitively parses 'input' case-insensitively into a DistributedAvailabilityGroupId
// note: this method should only be used for API response data and not user input
func ParseDistributedAvailabilityGroupIDInsensitively(input string) (*DistributedAvailabilityGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(DistributedAvailabilityGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DistributedAvailabilityGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.DistributedAvailabilityGroupName, ok = parsed.Parsed["distributedAvailabilityGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "distributedAvailabilityGroupName", *parsed)
	}

	return &id, nil
}

// ValidateDistributedAvailabilityGroupID checks that 'input' can be parsed as a Distributed Availability Group ID
func ValidateDistributedAvailabilityGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDistributedAvailabilityGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Distributed Availability Group ID
func (id DistributedAvailabilityGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/distributedAvailabilityGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DistributedAvailabilityGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Distributed Availability Group ID
func (id DistributedAvailabilityGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticDistributedAvailabilityGroups", "distributedAvailabilityGroups", "distributedAvailabilityGroups"),
		resourceids.UserSpecifiedSegment("distributedAvailabilityGroupName", "distributedAvailabilityGroupValue"),
	}
}

// String returns a human-readable description of this Distributed Availability Group ID
func (id DistributedAvailabilityGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Distributed Availability Group Name: %q", id.DistributedAvailabilityGroupName),
	}
	return fmt.Sprintf("Distributed Availability Group (%s)", strings.Join(components, "\n"))
}
