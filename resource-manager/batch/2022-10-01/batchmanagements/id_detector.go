package batchmanagements

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DetectorId{}

// DetectorId is a struct representing the Resource ID for a Detector
type DetectorId struct {
	SubscriptionId    string
	ResourceGroupName string
	BatchAccountName  string
	DetectorId        string
}

// NewDetectorID returns a new DetectorId struct
func NewDetectorID(subscriptionId string, resourceGroupName string, batchAccountName string, detectorId string) DetectorId {
	return DetectorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		BatchAccountName:  batchAccountName,
		DetectorId:        detectorId,
	}
}

// ParseDetectorID parses 'input' into a DetectorId
func ParseDetectorID(input string) (*DetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DetectorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.BatchAccountName, ok = parsed.Parsed["batchAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'batchAccountName' was not found in the resource id %q", input)
	}

	if id.DetectorId, ok = parsed.Parsed["detectorId"]; !ok {
		return nil, fmt.Errorf("the segment 'detectorId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDetectorIDInsensitively parses 'input' case-insensitively into a DetectorId
// note: this method should only be used for API response data and not user input
func ParseDetectorIDInsensitively(input string) (*DetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DetectorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.BatchAccountName, ok = parsed.Parsed["batchAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'batchAccountName' was not found in the resource id %q", input)
	}

	if id.DetectorId, ok = parsed.Parsed["detectorId"]; !ok {
		return nil, fmt.Errorf("the segment 'detectorId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDetectorID checks that 'input' can be parsed as a Detector ID
func ValidateDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Detector ID
func (id DetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Batch/batchAccounts/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.BatchAccountName, id.DetectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Detector ID
func (id DetectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBatch", "Microsoft.Batch", "Microsoft.Batch"),
		resourceids.StaticSegment("staticBatchAccounts", "batchAccounts", "batchAccounts"),
		resourceids.UserSpecifiedSegment("batchAccountName", "batchAccountValue"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorId", "detectorIdValue"),
	}
}

// String returns a human-readable description of this Detector ID
func (id DetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Batch Account Name: %q", id.BatchAccountName),
		fmt.Sprintf("Detector: %q", id.DetectorId),
	}
	return fmt.Sprintf("Detector (%s)", strings.Join(components, "\n"))
}
