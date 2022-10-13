package resourcepools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ResourcePoolId{}

// ResourcePoolId is a struct representing the Resource ID for a Resource Pool
type ResourcePoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourcePoolName  string
}

// NewResourcePoolID returns a new ResourcePoolId struct
func NewResourcePoolID(subscriptionId string, resourceGroupName string, resourcePoolName string) ResourcePoolId {
	return ResourcePoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourcePoolName:  resourcePoolName,
	}
}

// ParseResourcePoolID parses 'input' into a ResourcePoolId
func ParseResourcePoolID(input string) (*ResourcePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourcePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResourcePoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourcePoolName, ok = parsed.Parsed["resourcePoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourcePoolName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseResourcePoolIDInsensitively parses 'input' case-insensitively into a ResourcePoolId
// note: this method should only be used for API response data and not user input
func ParseResourcePoolIDInsensitively(input string) (*ResourcePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourcePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResourcePoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourcePoolName, ok = parsed.Parsed["resourcePoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourcePoolName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateResourcePoolID checks that 'input' can be parsed as a Resource Pool ID
func ValidateResourcePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourcePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Pool ID
func (id ResourcePoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/resourcePools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourcePoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Pool ID
func (id ResourcePoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticResourcePools", "resourcePools", "resourcePools"),
		resourceids.UserSpecifiedSegment("resourcePoolName", "resourcePoolValue"),
	}
}

// String returns a human-readable description of this Resource Pool ID
func (id ResourcePoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Pool Name: %q", id.ResourcePoolName),
	}
	return fmt.Sprintf("Resource Pool (%s)", strings.Join(components, "\n"))
}
