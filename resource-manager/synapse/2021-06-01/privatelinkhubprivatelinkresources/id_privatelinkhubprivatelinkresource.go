package privatelinkhubprivatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PrivateLinkHubPrivateLinkResourceId{}

// PrivateLinkHubPrivateLinkResourceId is a struct representing the Resource ID for a Private Link Hub Private Link Resource
type PrivateLinkHubPrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	PrivateLinkHubName      string
	PrivateLinkResourceName string
}

// NewPrivateLinkHubPrivateLinkResourceID returns a new PrivateLinkHubPrivateLinkResourceId struct
func NewPrivateLinkHubPrivateLinkResourceID(subscriptionId string, resourceGroupName string, privateLinkHubName string, privateLinkResourceName string) PrivateLinkHubPrivateLinkResourceId {
	return PrivateLinkHubPrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		PrivateLinkHubName:      privateLinkHubName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParsePrivateLinkHubPrivateLinkResourceID parses 'input' into a PrivateLinkHubPrivateLinkResourceId
func ParsePrivateLinkHubPrivateLinkResourceID(input string) (*PrivateLinkHubPrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkHubPrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkHubPrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkHubName, ok = parsed.Parsed["privateLinkHubName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkHubName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkResourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePrivateLinkHubPrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a PrivateLinkHubPrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkHubPrivateLinkResourceIDInsensitively(input string) (*PrivateLinkHubPrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkHubPrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkHubPrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkHubName, ok = parsed.Parsed["privateLinkHubName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkHubName' was not found in the resource id %q", input)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkResourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePrivateLinkHubPrivateLinkResourceID checks that 'input' can be parsed as a Private Link Hub Private Link Resource ID
func ValidatePrivateLinkHubPrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkHubPrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link Hub Private Link Resource ID
func (id PrivateLinkHubPrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/privateLinkHubs/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateLinkHubName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link Hub Private Link Resource ID
func (id PrivateLinkHubPrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticPrivateLinkHubs", "privateLinkHubs", "privateLinkHubs"),
		resourceids.UserSpecifiedSegment("privateLinkHubName", "privateLinkHubValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "privateLinkResourceValue"),
	}
}

// String returns a human-readable description of this Private Link Hub Private Link Resource ID
func (id PrivateLinkHubPrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Link Hub Name: %q", id.PrivateLinkHubName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Private Link Hub Private Link Resource (%s)", strings.Join(components, "\n"))
}
