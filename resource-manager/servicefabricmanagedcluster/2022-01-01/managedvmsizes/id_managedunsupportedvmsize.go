package managedvmsizes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagedUnsupportedVMSizeId{}

// ManagedUnsupportedVMSizeId is a struct representing the Resource ID for a Managed Unsupported V M Size
type ManagedUnsupportedVMSizeId struct {
	SubscriptionId               string
	LocationName                 string
	ManagedUnsupportedVMSizeName string
}

// NewManagedUnsupportedVMSizeID returns a new ManagedUnsupportedVMSizeId struct
func NewManagedUnsupportedVMSizeID(subscriptionId string, locationName string, managedUnsupportedVMSizeName string) ManagedUnsupportedVMSizeId {
	return ManagedUnsupportedVMSizeId{
		SubscriptionId:               subscriptionId,
		LocationName:                 locationName,
		ManagedUnsupportedVMSizeName: managedUnsupportedVMSizeName,
	}
}

// ParseManagedUnsupportedVMSizeID parses 'input' into a ManagedUnsupportedVMSizeId
func ParseManagedUnsupportedVMSizeID(input string) (*ManagedUnsupportedVMSizeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedUnsupportedVMSizeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedUnsupportedVMSizeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ManagedUnsupportedVMSizeName, ok = parsed.Parsed["managedUnsupportedVMSizeName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedUnsupportedVMSizeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseManagedUnsupportedVMSizeIDInsensitively parses 'input' case-insensitively into a ManagedUnsupportedVMSizeId
// note: this method should only be used for API response data and not user input
func ParseManagedUnsupportedVMSizeIDInsensitively(input string) (*ManagedUnsupportedVMSizeId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedUnsupportedVMSizeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedUnsupportedVMSizeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.ManagedUnsupportedVMSizeName, ok = parsed.Parsed["managedUnsupportedVMSizeName"]; !ok {
		return nil, fmt.Errorf("the segment 'managedUnsupportedVMSizeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateManagedUnsupportedVMSizeID checks that 'input' can be parsed as a Managed Unsupported V M Size ID
func ValidateManagedUnsupportedVMSizeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedUnsupportedVMSizeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Unsupported V M Size ID
func (id ManagedUnsupportedVMSizeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ServiceFabric/locations/%s/managedUnsupportedVMSizes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.ManagedUnsupportedVMSizeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Unsupported V M Size ID
func (id ManagedUnsupportedVMSizeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticManagedUnsupportedVMSizes", "managedUnsupportedVMSizes", "managedUnsupportedVMSizes"),
		resourceids.UserSpecifiedSegment("managedUnsupportedVMSizeName", "managedUnsupportedVMSizeValue"),
	}
}

// String returns a human-readable description of this Managed Unsupported V M Size ID
func (id ManagedUnsupportedVMSizeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Managed Unsupported V M Size Name: %q", id.ManagedUnsupportedVMSizeName),
	}
	return fmt.Sprintf("Managed Unsupported V M Size (%s)", strings.Join(components, "\n"))
}
