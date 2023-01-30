package bandwidthschedules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BandwidthScheduleId{}

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
	parser := resourceids.NewParserFromResourceIdType(BandwidthScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BandwidthScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.BandwidthScheduleName, ok = parsed.Parsed["bandwidthScheduleName"]; !ok {
		return nil, fmt.Errorf("the segment 'bandwidthScheduleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBandwidthScheduleIDInsensitively parses 'input' case-insensitively into a BandwidthScheduleId
// note: this method should only be used for API response data and not user input
func ParseBandwidthScheduleIDInsensitively(input string) (*BandwidthScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(BandwidthScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BandwidthScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.BandwidthScheduleName, ok = parsed.Parsed["bandwidthScheduleName"]; !ok {
		return nil, fmt.Errorf("the segment 'bandwidthScheduleName' was not found in the resource id %q", input)
	}

	return &id, nil
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
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticBandwidthSchedules", "bandwidthSchedules", "bandwidthSchedules"),
		resourceids.UserSpecifiedSegment("bandwidthScheduleName", "bandwidthScheduleValue"),
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
