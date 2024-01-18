package clusterversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnvironmentId{}

// EnvironmentId is a struct representing the Resource ID for a Environment
type EnvironmentId struct {
	SubscriptionId string
	LocationName   string
	Environment    ClusterVersionsEnvironment
}

// NewEnvironmentID returns a new EnvironmentId struct
func NewEnvironmentID(subscriptionId string, locationName string, environment ClusterVersionsEnvironment) EnvironmentId {
	return EnvironmentId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		Environment:    environment,
	}
}

// ParseEnvironmentID parses 'input' into a EnvironmentId
func ParseEnvironmentID(input string) (*EnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnvironmentIDInsensitively parses 'input' case-insensitively into a EnvironmentId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentIDInsensitively(input string) (*EnvironmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnvironmentId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateEnvironmentID checks that 'input' can be parsed as a Environment ID
func ValidateEnvironmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment ID
func (id EnvironmentId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/environments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, string(id.Environment))
}

// Segments returns a slice of Resource ID Segments which comprise this Environment ID
func (id EnvironmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.ConstantSegment("environment", PossibleValuesForClusterVersionsEnvironment(), "Linux"),
	}
}

// String returns a human-readable description of this Environment ID
func (id EnvironmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Environment: %q", string(id.Environment)),
	}
	return fmt.Sprintf("Environment (%s)", strings.Join(components, "\n"))
}
