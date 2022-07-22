package capabilities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CapabilityTypeId{}

// CapabilityTypeId is a struct representing the Resource ID for a Capability Type
type CapabilityTypeId struct {
	SubscriptionId     string
	LocationName       string
	TargetTypeName     string
	CapabilityTypeName string
}

// NewCapabilityTypeID returns a new CapabilityTypeId struct
func NewCapabilityTypeID(subscriptionId string, locationName string, targetTypeName string, capabilityTypeName string) CapabilityTypeId {
	return CapabilityTypeId{
		SubscriptionId:     subscriptionId,
		LocationName:       locationName,
		TargetTypeName:     targetTypeName,
		CapabilityTypeName: capabilityTypeName,
	}
}

// ParseCapabilityTypeID parses 'input' into a CapabilityTypeId
func ParseCapabilityTypeID(input string) (*CapabilityTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilityTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilityTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.TargetTypeName, ok = parsed.Parsed["targetTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetTypeName' was not found in the resource id %q", input)
	}

	if id.CapabilityTypeName, ok = parsed.Parsed["capabilityTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'capabilityTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCapabilityTypeIDInsensitively parses 'input' case-insensitively into a CapabilityTypeId
// note: this method should only be used for API response data and not user input
func ParseCapabilityTypeIDInsensitively(input string) (*CapabilityTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilityTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilityTypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.TargetTypeName, ok = parsed.Parsed["targetTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetTypeName' was not found in the resource id %q", input)
	}

	if id.CapabilityTypeName, ok = parsed.Parsed["capabilityTypeName"]; !ok {
		return nil, fmt.Errorf("the segment 'capabilityTypeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCapabilityTypeID checks that 'input' can be parsed as a Capability Type ID
func ValidateCapabilityTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCapabilityTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Capability Type ID
func (id CapabilityTypeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Chaos/locations/%s/targetTypes/%s/capabilityTypes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.TargetTypeName, id.CapabilityTypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Capability Type ID
func (id CapabilityTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticTargetTypes", "targetTypes", "targetTypes"),
		resourceids.UserSpecifiedSegment("targetTypeName", "targetTypeValue"),
		resourceids.StaticSegment("staticCapabilityTypes", "capabilityTypes", "capabilityTypes"),
		resourceids.UserSpecifiedSegment("capabilityTypeName", "capabilityTypeValue"),
	}
}

// String returns a human-readable description of this Capability Type ID
func (id CapabilityTypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Target Type Name: %q", id.TargetTypeName),
		fmt.Sprintf("Capability Type Name: %q", id.CapabilityTypeName),
	}
	return fmt.Sprintf("Capability Type (%s)", strings.Join(components, "\n"))
}
