package customlocations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CustomLocationId{}

// CustomLocationId is a struct representing the Resource ID for a Custom Location
type CustomLocationId struct {
	SubscriptionId     string
	ResourceGroupName  string
	CustomLocationName string
}

// NewCustomLocationID returns a new CustomLocationId struct
func NewCustomLocationID(subscriptionId string, resourceGroupName string, customLocationName string) CustomLocationId {
	return CustomLocationId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		CustomLocationName: customLocationName,
	}
}

// ParseCustomLocationID parses 'input' into a CustomLocationId
func ParseCustomLocationID(input string) (*CustomLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomLocationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomLocationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CustomLocationName, ok = parsed.Parsed["customLocationName"]; !ok {
		return nil, fmt.Errorf("the segment 'customLocationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCustomLocationIDInsensitively parses 'input' case-insensitively into a CustomLocationId
// note: this method should only be used for API response data and not user input
func ParseCustomLocationIDInsensitively(input string) (*CustomLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomLocationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomLocationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CustomLocationName, ok = parsed.Parsed["customLocationName"]; !ok {
		return nil, fmt.Errorf("the segment 'customLocationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCustomLocationID checks that 'input' can be parsed as a Custom Location ID
func ValidateCustomLocationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomLocationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Custom Location ID
func (id CustomLocationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ExtendedLocation/customLocations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CustomLocationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Custom Location ID
func (id CustomLocationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftExtendedLocation", "Microsoft.ExtendedLocation", "Microsoft.ExtendedLocation"),
		resourceids.StaticSegment("staticCustomLocations", "customLocations", "customLocations"),
		resourceids.UserSpecifiedSegment("customLocationName", "customLocationValue"),
	}
}

// String returns a human-readable description of this Custom Location ID
func (id CustomLocationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Custom Location Name: %q", id.CustomLocationName),
	}
	return fmt.Sprintf("Custom Location (%s)", strings.Join(components, "\n"))
}
