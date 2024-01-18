package managedclusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ManagedClusterVersionId{}

// ManagedClusterVersionId is a struct representing the Resource ID for a Managed Cluster Version
type ManagedClusterVersionId struct {
	SubscriptionId            string
	LocationName              string
	ManagedClusterVersionName string
}

// NewManagedClusterVersionID returns a new ManagedClusterVersionId struct
func NewManagedClusterVersionID(subscriptionId string, locationName string, managedClusterVersionName string) ManagedClusterVersionId {
	return ManagedClusterVersionId{
		SubscriptionId:            subscriptionId,
		LocationName:              locationName,
		ManagedClusterVersionName: managedClusterVersionName,
	}
}

// ParseManagedClusterVersionID parses 'input' into a ManagedClusterVersionId
func ParseManagedClusterVersionID(input string) (*ManagedClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedClusterVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedClusterVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedClusterVersionIDInsensitively parses 'input' case-insensitively into a ManagedClusterVersionId
// note: this method should only be used for API response data and not user input
func ParseManagedClusterVersionIDInsensitively(input string) (*ManagedClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedClusterVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedClusterVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedClusterVersionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateManagedClusterVersionID checks that 'input' can be parsed as a Managed Cluster Version ID
func ValidateManagedClusterVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedClusterVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Cluster Version ID
func (id ManagedClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/managedClusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.ManagedClusterVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Cluster Version ID
func (id ManagedClusterVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticManagedClusterVersions", "managedClusterVersions", "managedClusterVersions"),
		resourceids.UserSpecifiedSegment("managedClusterVersionName", "managedClusterVersionValue"),
	}
}

// String returns a human-readable description of this Managed Cluster Version ID
func (id ManagedClusterVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Managed Cluster Version Name: %q", id.ManagedClusterVersionName),
	}
	return fmt.Sprintf("Managed Cluster Version (%s)", strings.Join(components, "\n"))
}
