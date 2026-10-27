package hcpoperatoridentityrolesets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HcpOperatorIdentityRoleSetId{})
}

var _ resourceids.ResourceId = &HcpOperatorIdentityRoleSetId{}

// HcpOperatorIdentityRoleSetId is a struct representing the Resource ID for a Hcp Operator Identity Role Set
type HcpOperatorIdentityRoleSetId struct {
	SubscriptionId                 string
	LocationName                   string
	HcpOperatorIdentityRoleSetName string
}

// NewHcpOperatorIdentityRoleSetID returns a new HcpOperatorIdentityRoleSetId struct
func NewHcpOperatorIdentityRoleSetID(subscriptionId string, locationName string, hcpOperatorIdentityRoleSetName string) HcpOperatorIdentityRoleSetId {
	return HcpOperatorIdentityRoleSetId{
		SubscriptionId:                 subscriptionId,
		LocationName:                   locationName,
		HcpOperatorIdentityRoleSetName: hcpOperatorIdentityRoleSetName,
	}
}

// ParseHcpOperatorIdentityRoleSetID parses 'input' into a HcpOperatorIdentityRoleSetId
func ParseHcpOperatorIdentityRoleSetID(input string) (*HcpOperatorIdentityRoleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcpOperatorIdentityRoleSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcpOperatorIdentityRoleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHcpOperatorIdentityRoleSetIDInsensitively parses 'input' case-insensitively into a HcpOperatorIdentityRoleSetId
// note: this method should only be used for API response data and not user input
func ParseHcpOperatorIdentityRoleSetIDInsensitively(input string) (*HcpOperatorIdentityRoleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcpOperatorIdentityRoleSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcpOperatorIdentityRoleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HcpOperatorIdentityRoleSetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.HcpOperatorIdentityRoleSetName, ok = input.Parsed["hcpOperatorIdentityRoleSetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hcpOperatorIdentityRoleSetName", input)
	}

	return nil
}

// ValidateHcpOperatorIdentityRoleSetID checks that 'input' can be parsed as a Hcp Operator Identity Role Set ID
func ValidateHcpOperatorIdentityRoleSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHcpOperatorIdentityRoleSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hcp Operator Identity Role Set ID
func (id HcpOperatorIdentityRoleSetId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.RedHatOpenShift/locations/%s/hcpOperatorIdentityRoleSets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.HcpOperatorIdentityRoleSetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hcp Operator Identity Role Set ID
func (id HcpOperatorIdentityRoleSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticHcpOperatorIdentityRoleSets", "hcpOperatorIdentityRoleSets", "hcpOperatorIdentityRoleSets"),
		resourceids.UserSpecifiedSegment("hcpOperatorIdentityRoleSetName", "hcpOperatorIdentityRoleSetName"),
	}
}

// String returns a human-readable description of this Hcp Operator Identity Role Set ID
func (id HcpOperatorIdentityRoleSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Hcp Operator Identity Role Set Name: %q", id.HcpOperatorIdentityRoleSetName),
	}
	return fmt.Sprintf("Hcp Operator Identity Role Set (%s)", strings.Join(components, "\n"))
}
