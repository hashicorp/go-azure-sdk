package restorables

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RestorableDatabaseAccountId{}

// RestorableDatabaseAccountId is a struct representing the Resource ID for a Restorable Database Account
type RestorableDatabaseAccountId struct {
	SubscriptionId string
	Location       string
	InstanceId     string
}

// NewRestorableDatabaseAccountID returns a new RestorableDatabaseAccountId struct
func NewRestorableDatabaseAccountID(subscriptionId string, location string, instanceId string) RestorableDatabaseAccountId {
	return RestorableDatabaseAccountId{
		SubscriptionId: subscriptionId,
		Location:       location,
		InstanceId:     instanceId,
	}
}

// ParseRestorableDatabaseAccountID parses 'input' into a RestorableDatabaseAccountId
func ParseRestorableDatabaseAccountID(input string) (*RestorableDatabaseAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorableDatabaseAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorableDatabaseAccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, fmt.Errorf("the segment 'instanceId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRestorableDatabaseAccountIDInsensitively parses 'input' case-insensitively into a RestorableDatabaseAccountId
// note: this method should only be used for API response data and not user input
func ParseRestorableDatabaseAccountIDInsensitively(input string) (*RestorableDatabaseAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorableDatabaseAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorableDatabaseAccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, fmt.Errorf("the segment 'instanceId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRestorableDatabaseAccountID checks that 'input' can be parsed as a Restorable Database Account ID
func ValidateRestorableDatabaseAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorableDatabaseAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restorable Database Account ID
func (id RestorableDatabaseAccountId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.DocumentDB/locations/%s/restorableDatabaseAccounts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.InstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Restorable Database Account ID
func (id RestorableDatabaseAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticRestorableDatabaseAccounts", "restorableDatabaseAccounts", "restorableDatabaseAccounts"),
		resourceids.UserSpecifiedSegment("instanceId", "instanceIdValue"),
	}
}

// String returns a human-readable description of this Restorable Database Account ID
func (id RestorableDatabaseAccountId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Instance: %q", id.InstanceId),
	}
	return fmt.Sprintf("Restorable Database Account (%s)", strings.Join(components, "\n"))
}
