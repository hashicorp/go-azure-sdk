package deletedservice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DeletedServiceId{}

// DeletedServiceId is a struct representing the Resource ID for a Deleted Service
type DeletedServiceId struct {
	SubscriptionId string
	Location       string
	ServiceName    string
}

// NewDeletedServiceID returns a new DeletedServiceId struct
func NewDeletedServiceID(subscriptionId string, location string, serviceName string) DeletedServiceId {
	return DeletedServiceId{
		SubscriptionId: subscriptionId,
		Location:       location,
		ServiceName:    serviceName,
	}
}

// ParseDeletedServiceID parses 'input' into a DeletedServiceId
func ParseDeletedServiceID(input string) (*DeletedServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDeletedServiceIDInsensitively parses 'input' case-insensitively into a DeletedServiceId
// note: this method should only be used for API response data and not user input
func ParseDeletedServiceIDInsensitively(input string) (*DeletedServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDeletedServiceID checks that 'input' can be parsed as a Deleted Service ID
func ValidateDeletedServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeletedServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deleted Service ID
func (id DeletedServiceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.ApiManagement/locations/%s/deletedServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.ServiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Deleted Service ID
func (id DeletedServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticDeletedServices", "deletedServices", "deletedServices"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
	}
}

// String returns a human-readable description of this Deleted Service ID
func (id DeletedServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
	}
	return fmt.Sprintf("Deleted Service (%s)", strings.Join(components, "\n"))
}
