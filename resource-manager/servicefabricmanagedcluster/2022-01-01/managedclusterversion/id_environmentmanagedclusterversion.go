package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EnvironmentManagedClusterVersionId{})
}

var _ resourceids.ResourceId = &EnvironmentManagedClusterVersionId{}

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
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentManagedClusterVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentManagedClusterVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnvironmentManagedClusterVersionIDInsensitively parses 'input' case-insensitively into a EnvironmentManagedClusterVersionId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentManagedClusterVersionIDInsensitively(input string) (*EnvironmentManagedClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentManagedClusterVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentManagedClusterVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnvironmentManagedClusterVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.ManagedClusterVersionName, ok = input.Parsed["managedClusterVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedClusterVersionName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("locationName", "location"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.StaticSegment("environment", "Windows", "Windows"),
		resourceids.StaticSegment("staticManagedClusterVersions", "managedClusterVersions", "managedClusterVersions"),
		resourceids.UserSpecifiedSegment("managedClusterVersionName", "clusterVersion"),
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
