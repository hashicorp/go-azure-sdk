package hcpopenshiftversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HcpOpenShiftVersionId{})
}

var _ resourceids.ResourceId = &HcpOpenShiftVersionId{}

// HcpOpenShiftVersionId is a struct representing the Resource ID for a Hcp Open Shift Version
type HcpOpenShiftVersionId struct {
	SubscriptionId          string
	LocationName            string
	HcpOpenShiftVersionName string
}

// NewHcpOpenShiftVersionID returns a new HcpOpenShiftVersionId struct
func NewHcpOpenShiftVersionID(subscriptionId string, locationName string, hcpOpenShiftVersionName string) HcpOpenShiftVersionId {
	return HcpOpenShiftVersionId{
		SubscriptionId:          subscriptionId,
		LocationName:            locationName,
		HcpOpenShiftVersionName: hcpOpenShiftVersionName,
	}
}

// ParseHcpOpenShiftVersionID parses 'input' into a HcpOpenShiftVersionId
func ParseHcpOpenShiftVersionID(input string) (*HcpOpenShiftVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcpOpenShiftVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcpOpenShiftVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHcpOpenShiftVersionIDInsensitively parses 'input' case-insensitively into a HcpOpenShiftVersionId
// note: this method should only be used for API response data and not user input
func ParseHcpOpenShiftVersionIDInsensitively(input string) (*HcpOpenShiftVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcpOpenShiftVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcpOpenShiftVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HcpOpenShiftVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.HcpOpenShiftVersionName, ok = input.Parsed["hcpOpenShiftVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hcpOpenShiftVersionName", input)
	}

	return nil
}

// ValidateHcpOpenShiftVersionID checks that 'input' can be parsed as a Hcp Open Shift Version ID
func ValidateHcpOpenShiftVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHcpOpenShiftVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hcp Open Shift Version ID
func (id HcpOpenShiftVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.RedHatOpenShift/locations/%s/hcpOpenShiftVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.HcpOpenShiftVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hcp Open Shift Version ID
func (id HcpOpenShiftVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticHcpOpenShiftVersions", "hcpOpenShiftVersions", "hcpOpenShiftVersions"),
		resourceids.UserSpecifiedSegment("hcpOpenShiftVersionName", "hcpOpenShiftVersionName"),
	}
}

// String returns a human-readable description of this Hcp Open Shift Version ID
func (id HcpOpenShiftVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Hcp Open Shift Version Name: %q", id.HcpOpenShiftVersionName),
	}
	return fmt.Sprintf("Hcp Open Shift Version (%s)", strings.Join(components, "\n"))
}
