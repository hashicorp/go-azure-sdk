package incrementalrestorepoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RestorePointId{}

// RestorePointId is a struct representing the Resource ID for a Restore Point
type RestorePointId struct {
	SubscriptionId             string
	ResourceGroupName          string
	RestorePointCollectionName string
	VmRestorePointName         string
}

// NewRestorePointID returns a new RestorePointId struct
func NewRestorePointID(subscriptionId string, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string) RestorePointId {
	return RestorePointId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		RestorePointCollectionName: restorePointCollectionName,
		VmRestorePointName:         vmRestorePointName,
	}
}

// ParseRestorePointID parses 'input' into a RestorePointId
func ParseRestorePointID(input string) (*RestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	if id.VmRestorePointName, ok = parsed.Parsed["vmRestorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmRestorePointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRestorePointIDInsensitively parses 'input' case-insensitively into a RestorePointId
// note: this method should only be used for API response data and not user input
func ParseRestorePointIDInsensitively(input string) (*RestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	if id.VmRestorePointName, ok = parsed.Parsed["vmRestorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmRestorePointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRestorePointID checks that 'input' can be parsed as a Restore Point ID
func ValidateRestorePointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorePointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restore Point ID
func (id RestorePointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/restorePointCollections/%s/restorePoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RestorePointCollectionName, id.VmRestorePointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Restore Point ID
func (id RestorePointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticRestorePointCollections", "restorePointCollections", "restorePointCollections"),
		resourceids.UserSpecifiedSegment("restorePointCollectionName", "restorePointCollectionValue"),
		resourceids.StaticSegment("staticRestorePoints", "restorePoints", "restorePoints"),
		resourceids.UserSpecifiedSegment("vmRestorePointName", "vmRestorePointValue"),
	}
}

// String returns a human-readable description of this Restore Point ID
func (id RestorePointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Restore Point Collection Name: %q", id.RestorePointCollectionName),
		fmt.Sprintf("Vm Restore Point Name: %q", id.VmRestorePointName),
	}
	return fmt.Sprintf("Restore Point (%s)", strings.Join(components, "\n"))
}
