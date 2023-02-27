// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EnvironmentManagedClusterVersionId{}

// EnvironmentManagedClusterVersionId is a struct representing the Resource ID for a Environment Managed Cluster Version
type EnvironmentManagedClusterVersionId struct {
	SubscriptionId            string
	LocationName              string
	ManagedClusterVersionName string
}

// NewEnvironmentManagedClusterVersionID returns a new EnvironmentManagedClusterVersionId struct
func NewEnvironmentManagedClusterVersionID(subscriptionId string, locationName string, managedClusterVersionName string) EnvironmentManagedClusterVersionId {
	return EnvironmentManagedClusterVersionId{
		SubscriptionId:            subscriptionId,
		LocationName:              locationName,
		ManagedClusterVersionName: managedClusterVersionName,
	}
}

// ParseEnvironmentManagedClusterVersionID parses 'input' into a EnvironmentManagedClusterVersionId
func ParseEnvironmentManagedClusterVersionID(input string) (*EnvironmentManagedClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentManagedClusterVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentManagedClusterVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ManagedClusterVersionName, ok = parsed.Parsed["managedClusterVersionName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedClusterVersionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEnvironmentManagedClusterVersionIDInsensitively parses 'input' case-insensitively into a EnvironmentManagedClusterVersionId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentManagedClusterVersionIDInsensitively(input string) (*EnvironmentManagedClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentManagedClusterVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EnvironmentManagedClusterVersionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ManagedClusterVersionName, ok = parsed.Parsed["managedClusterVersionName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedClusterVersionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEnvironmentManagedClusterVersionID checks that 'input' can be parsed as a Environment Managed Cluster Version ID
func ValidateEnvironmentManagedClusterVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentManagedClusterVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment Managed Cluster Version ID
func (id EnvironmentManagedClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/environments/Windows/managedClusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.ManagedClusterVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Environment Managed Cluster Version ID
func (id EnvironmentManagedClusterVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.StaticSegment("environment", "Windows", "Windows"),
		resourceids.StaticSegment("staticManagedClusterVersions", "managedClusterVersions", "managedClusterVersions"),
		resourceids.UserSpecifiedSegment("managedClusterVersionName", "managedClusterVersionValue"),
	}
}

// String returns a human-readable description of this Environment Managed Cluster Version ID
func (id EnvironmentManagedClusterVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Managed Cluster Version Name: %q", id.ManagedClusterVersionName),
	}
	return fmt.Sprintf("Environment Managed Cluster Version (%s)", strings.Join(components, "\n"))
}
