package clusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = EnvironmentClusterVersionId{}

// EnvironmentClusterVersionId is a struct representing the Resource ID for a Environment Cluster Version
type EnvironmentClusterVersionId struct {
	SubscriptionId     string
	LocationName       string
	Environment        ClusterVersionsEnvironment
	ClusterVersionName string
}

// NewEnvironmentClusterVersionID returns a new EnvironmentClusterVersionId struct
func NewEnvironmentClusterVersionID(subscriptionId string, locationName string, environment ClusterVersionsEnvironment, clusterVersionName string) EnvironmentClusterVersionId {
	return EnvironmentClusterVersionId{
		SubscriptionId:     subscriptionId,
		LocationName:       locationName,
		Environment:        environment,
		ClusterVersionName: clusterVersionName,
	}
}

// ParseEnvironmentClusterVersionID parses 'input' into a EnvironmentClusterVersionId
func ParseEnvironmentClusterVersionID(input string) (*EnvironmentClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentClusterVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentClusterVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnvironmentClusterVersionIDInsensitively parses 'input' case-insensitively into a EnvironmentClusterVersionId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentClusterVersionIDInsensitively(input string) (*EnvironmentClusterVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(EnvironmentClusterVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentClusterVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnvironmentClusterVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if v, ok := input.Parsed["environment"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "environment", input)
		}

		environment, err := parseClusterVersionsEnvironment(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.Environment = *environment
	}

	if id.ClusterVersionName, ok = input.Parsed["clusterVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterVersionName", input)
	}

	return nil
}

// ValidateEnvironmentClusterVersionID checks that 'input' can be parsed as a Environment Cluster Version ID
func ValidateEnvironmentClusterVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentClusterVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment Cluster Version ID
func (id EnvironmentClusterVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/environments/%s/clusterVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, string(id.Environment), id.ClusterVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Environment Cluster Version ID
func (id EnvironmentClusterVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.ConstantSegment("environment", PossibleValuesForClusterVersionsEnvironment(), "Linux"),
		resourceids.StaticSegment("staticClusterVersions", "clusterVersions", "clusterVersions"),
		resourceids.UserSpecifiedSegment("clusterVersionName", "clusterVersionValue"),
	}
}

// String returns a human-readable description of this Environment Cluster Version ID
func (id EnvironmentClusterVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Environment: %q", string(id.Environment)),
		fmt.Sprintf("Cluster Version Name: %q", id.ClusterVersionName),
	}
	return fmt.Sprintf("Environment Cluster Version (%s)", strings.Join(components, "\n"))
}
