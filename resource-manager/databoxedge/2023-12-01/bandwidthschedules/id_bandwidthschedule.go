package bandwidthschedules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BandwidthScheduleId{})
}

var _ resourceids.ResourceId = &BandwidthScheduleId{}

// BandwidthScheduleId is a struct representing the Resource ID for a Bandwidth Schedule
type BandwidthScheduleId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	BandwidthScheduleName string
}

// NewBandwidthScheduleID returns a new BandwidthScheduleId struct
func NewBandwidthScheduleID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, bandwidthScheduleName string) BandwidthScheduleId {
	return BandwidthScheduleId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		BandwidthScheduleName: bandwidthScheduleName,
	}
}

// ParseBandwidthScheduleID parses 'input' into a BandwidthScheduleId
func ParseBandwidthScheduleID(input string) (*BandwidthScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BandwidthScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BandwidthScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBandwidthScheduleIDInsensitively parses 'input' case-insensitively into a BandwidthScheduleId
// note: this method should only be used for API response data and not user input
func ParseBandwidthScheduleIDInsensitively(input string) (*BandwidthScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BandwidthScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BandwidthScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BandwidthScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DataBoxEdgeDeviceName, ok = input.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", input)
	}

	if id.BandwidthScheduleName, ok = input.Parsed["bandwidthScheduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bandwidthScheduleName", input)
	}

	return nil
}

// ValidateBandwidthScheduleID checks that 'input' can be parsed as a Bandwidth Schedule ID
func ValidateBandwidthScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBandwidthScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bandwidth Schedule ID
func (id BandwidthScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/bandwidthSchedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.BandwidthScheduleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Bandwidth Schedule ID
func (id BandwidthScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceName"),
		resourceids.StaticSegment("staticBandwidthSchedules", "bandwidthSchedules", "bandwidthSchedules"),
		resourceids.UserSpecifiedSegment("bandwidthScheduleName", "bandwidthScheduleName"),
	}
}

// String returns a human-readable description of this Bandwidth Schedule ID
func (id BandwidthScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Bandwidth Schedule Name: %q", id.BandwidthScheduleName),
	}
	return fmt.Sprintf("Bandwidth Schedule (%s)", strings.Join(components, "\n"))
}
