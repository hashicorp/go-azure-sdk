package privatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PrivateLinkResourceId{}

// PrivateLinkResourceId is a struct representing the Resource ID for a Private Link Resource
type PrivateLinkResourceId struct {
	SubscriptionId    string
	ResourceGroupName string
	ScopeName         string
	GroupName         string
}

// NewPrivateLinkResourceID returns a new PrivateLinkResourceId struct
func NewPrivateLinkResourceID(subscriptionId string, resourceGroupName string, scopeName string, groupName string) PrivateLinkResourceId {
	return PrivateLinkResourceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ScopeName:         scopeName,
		GroupName:         groupName,
	}
}

// ParsePrivateLinkResourceID parses 'input' into a PrivateLinkResourceId
func ParsePrivateLinkResourceID(input string) (*PrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ScopeName, ok = parsed.Parsed["scopeName"]; !ok {
		return nil, fmt.Errorf("the segment 'scopeName' was not found in the resource id %q", input)
	}

	if id.GroupName, ok = parsed.Parsed["groupName"]; !ok {
		return nil, fmt.Errorf("the segment 'groupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a PrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkResourceIDInsensitively(input string) (*PrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ScopeName, ok = parsed.Parsed["scopeName"]; !ok {
		return nil, fmt.Errorf("the segment 'scopeName' was not found in the resource id %q", input)
	}

	if id.GroupName, ok = parsed.Parsed["groupName"]; !ok {
		return nil, fmt.Errorf("the segment 'groupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePrivateLinkResourceID checks that 'input' can be parsed as a Private Link Resource ID
func ValidatePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link Resource ID
func (id PrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/privateLinkScopes/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ScopeName, id.GroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link Resource ID
func (id PrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticPrivateLinkScopes", "privateLinkScopes", "privateLinkScopes"),
		resourceids.UserSpecifiedSegment("scopeName", "scopeValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("groupName", "groupValue"),
	}
}

// String returns a human-readable description of this Private Link Resource ID
func (id PrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Scope Name: %q", id.ScopeName),
		fmt.Sprintf("Group Name: %q", id.GroupName),
	}
	return fmt.Sprintf("Private Link Resource (%s)", strings.Join(components, "\n"))
}
