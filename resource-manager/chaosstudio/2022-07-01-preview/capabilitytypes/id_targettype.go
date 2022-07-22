package capabilitytypes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = TargetTypeId{}

// TargetTypeId is a struct representing the Resource ID for a Target Type
type TargetTypeId struct {
	SubscriptionId string
	LocationName   string
	TargetTypeName string
}

// NewTargetTypeID returns a new TargetTypeId struct
func NewTargetTypeID(subscriptionId string, locationName string, targetTypeName string) TargetTypeId {
	return TargetTypeId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		TargetTypeName: targetTypeName,
	}
}

// ParseTargetTypeID parses 'input' into a TargetTypeId
func ParseTargetTypeID(input string) (*TargetTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.TargetTypeName, ok = parsed.Parsed["targetTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseTargetTypeIDInsensitively parses 'input' case-insensitively into a TargetTypeId
// note: this method should only be used for API response data and not user input
func ParseTargetTypeIDInsensitively(input string) (*TargetTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.TargetTypeName, ok = parsed.Parsed["targetTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateTargetTypeID checks that 'input' can be parsed as a Target Type ID
func ValidateTargetTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTargetTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Target Type ID
func (id TargetTypeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Chaos/locations/%s/targetTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.TargetTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Target Type ID
func (id TargetTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticTargetTypes", "targetTypes", "targetTypes"),
		resourceids.UserSpecifiedSegment("targetTypeName", "targetTypeValue"),
	}
}

// String returns a human-readable description of this Target Type ID
func (id TargetTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Target Type Name: %q", id.TargetTypeName),
	}
	return fmt.Sprintf("Target Type (%s)", strings.Join(components, "\n"))
}
