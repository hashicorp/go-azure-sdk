package platformworkloadidentityrolesets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PlatformWorkloadIdentityRoleSetId{})
}

var _ resourceids.ResourceId = &PlatformWorkloadIdentityRoleSetId{}

// PlatformWorkloadIdentityRoleSetId is a struct representing the Resource ID for a Platform Workload Identity Role Set
type PlatformWorkloadIdentityRoleSetId struct {
	SubscriptionId                      string
	LocationName                        string
	PlatformWorkloadIdentityRoleSetName string
}

// NewPlatformWorkloadIdentityRoleSetID returns a new PlatformWorkloadIdentityRoleSetId struct
func NewPlatformWorkloadIdentityRoleSetID(subscriptionId string, locationName string, platformWorkloadIdentityRoleSetName string) PlatformWorkloadIdentityRoleSetId {
	return PlatformWorkloadIdentityRoleSetId{
		SubscriptionId:                      subscriptionId,
		LocationName:                        locationName,
		PlatformWorkloadIdentityRoleSetName: platformWorkloadIdentityRoleSetName,
	}
}

// ParsePlatformWorkloadIdentityRoleSetID parses 'input' into a PlatformWorkloadIdentityRoleSetId
func ParsePlatformWorkloadIdentityRoleSetID(input string) (*PlatformWorkloadIdentityRoleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PlatformWorkloadIdentityRoleSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PlatformWorkloadIdentityRoleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePlatformWorkloadIdentityRoleSetIDInsensitively parses 'input' case-insensitively into a PlatformWorkloadIdentityRoleSetId
// note: this method should only be used for API response data and not user input
func ParsePlatformWorkloadIdentityRoleSetIDInsensitively(input string) (*PlatformWorkloadIdentityRoleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PlatformWorkloadIdentityRoleSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PlatformWorkloadIdentityRoleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PlatformWorkloadIdentityRoleSetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.PlatformWorkloadIdentityRoleSetName, ok = input.Parsed["platformWorkloadIdentityRoleSetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "platformWorkloadIdentityRoleSetName", input)
	}

	return nil
}

// ValidatePlatformWorkloadIdentityRoleSetID checks that 'input' can be parsed as a Platform Workload Identity Role Set ID
func ValidatePlatformWorkloadIdentityRoleSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePlatformWorkloadIdentityRoleSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Platform Workload Identity Role Set ID
func (id PlatformWorkloadIdentityRoleSetId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.RedHatOpenShift/locations/%s/platformWorkloadIdentityRoleSets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.PlatformWorkloadIdentityRoleSetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Platform Workload Identity Role Set ID
func (id PlatformWorkloadIdentityRoleSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticPlatformWorkloadIdentityRoleSets", "platformWorkloadIdentityRoleSets", "platformWorkloadIdentityRoleSets"),
		resourceids.UserSpecifiedSegment("platformWorkloadIdentityRoleSetName", "platformWorkloadIdentityRoleSetName"),
	}
}

// String returns a human-readable description of this Platform Workload Identity Role Set ID
func (id PlatformWorkloadIdentityRoleSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Platform Workload Identity Role Set Name: %q", id.PlatformWorkloadIdentityRoleSetName),
	}
	return fmt.Sprintf("Platform Workload Identity Role Set (%s)", strings.Join(components, "\n"))
}
