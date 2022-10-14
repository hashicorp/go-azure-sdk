package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SegmentId{}

// SegmentId is a struct representing the Resource ID for a Segment
type SegmentId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	SegmentId         string
}

// NewSegmentID returns a new SegmentId struct
func NewSegmentID(subscriptionId string, resourceGroupName string, privateCloudName string, segmentId string) SegmentId {
	return SegmentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		SegmentId:         segmentId,
	}
}

// ParseSegmentID parses 'input' into a SegmentId
func ParseSegmentID(input string) (*SegmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(SegmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SegmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.SegmentId, ok = parsed.Parsed["segmentId"]; !ok {
		return nil, fmt.Errorf("the segment 'segmentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSegmentIDInsensitively parses 'input' case-insensitively into a SegmentId
// note: this method should only be used for API response data and not user input
func ParseSegmentIDInsensitively(input string) (*SegmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(SegmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SegmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.SegmentId, ok = parsed.Parsed["segmentId"]; !ok {
		return nil, fmt.Errorf("the segment 'segmentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSegmentID checks that 'input' can be parsed as a Segment ID
func ValidateSegmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSegmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Segment ID
func (id SegmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/segments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.SegmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Segment ID
func (id SegmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticSegments", "segments", "segments"),
		resourceids.UserSpecifiedSegment("segmentId", "segmentIdValue"),
	}
}

// String returns a human-readable description of this Segment ID
func (id SegmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Segment: %q", id.SegmentId),
	}
	return fmt.Sprintf("Segment (%s)", strings.Join(components, "\n"))
}
