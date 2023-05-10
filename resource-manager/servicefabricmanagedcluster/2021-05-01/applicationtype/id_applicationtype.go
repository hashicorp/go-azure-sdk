package applicationtype

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationTypeId{}

// ApplicationTypeId is a struct representing the Resource ID for a Application Type
type ApplicationTypeId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedClusterName  string
	ApplicationTypeName string
}

// NewApplicationTypeID returns a new ApplicationTypeId struct
func NewApplicationTypeID(subscriptionId string, resourceGroupName string, managedClusterName string, applicationTypeName string) ApplicationTypeId {
	return ApplicationTypeId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedClusterName:  managedClusterName,
		ApplicationTypeName: applicationTypeName,
	}
}

// ParseApplicationTypeID parses 'input' into a ApplicationTypeId
func ParseApplicationTypeID(input string) (*ApplicationTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedClusterName, ok = parsed.Parsed["managedClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", *parsed)
	}

	if id.ApplicationTypeName, ok = parsed.Parsed["applicationTypeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationTypeName", *parsed)
	}

	return &id, nil
}

// ParseApplicationTypeIDInsensitively parses 'input' case-insensitively into a ApplicationTypeId
// note: this method should only be used for API response data and not user input
func ParseApplicationTypeIDInsensitively(input string) (*ApplicationTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedClusterName, ok = parsed.Parsed["managedClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", *parsed)
	}

	if id.ApplicationTypeName, ok = parsed.Parsed["applicationTypeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationTypeName", *parsed)
	}

	return &id, nil
}

// ValidateApplicationTypeID checks that 'input' can be parsed as a Application Type ID
func ValidateApplicationTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Type ID
func (id ApplicationTypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applicationTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedClusterName, id.ApplicationTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Type ID
func (id ApplicationTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticManagedClusters", "managedClusters", "managedClusters"),
		resourceids.UserSpecifiedSegment("managedClusterName", "managedClusterValue"),
		resourceids.StaticSegment("staticApplicationTypes", "applicationTypes", "applicationTypes"),
		resourceids.UserSpecifiedSegment("applicationTypeName", "applicationTypeValue"),
	}
}

// String returns a human-readable description of this Application Type ID
func (id ApplicationTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Cluster Name: %q", id.ManagedClusterName),
		fmt.Sprintf("Application Type Name: %q", id.ApplicationTypeName),
	}
	return fmt.Sprintf("Application Type (%s)", strings.Join(components, "\n"))
}
