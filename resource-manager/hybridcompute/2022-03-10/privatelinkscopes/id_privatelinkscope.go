package privatelinkscopes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PrivateLinkScopeId{}

// PrivateLinkScopeId is a struct representing the Resource ID for a Private Link Scope
type PrivateLinkScopeId struct {
	SubscriptionId     string
	Location           string
	PrivateLinkScopeId string
}

// NewPrivateLinkScopeID returns a new PrivateLinkScopeId struct
func NewPrivateLinkScopeID(subscriptionId string, location string, privateLinkScopeId string) PrivateLinkScopeId {
	return PrivateLinkScopeId{
		SubscriptionId:     subscriptionId,
		Location:           location,
		PrivateLinkScopeId: privateLinkScopeId,
	}
}

// ParsePrivateLinkScopeID parses 'input' into a PrivateLinkScopeId
func ParsePrivateLinkScopeID(input string) (*PrivateLinkScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkScopeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PrivateLinkScopeId, ok = parsed.Parsed["privateLinkScopeId"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkScopeId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePrivateLinkScopeIDInsensitively parses 'input' case-insensitively into a PrivateLinkScopeId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkScopeIDInsensitively(input string) (*PrivateLinkScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrivateLinkScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrivateLinkScopeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PrivateLinkScopeId, ok = parsed.Parsed["privateLinkScopeId"]; !ok {
		return nil, fmt.Errorf("the segment 'privateLinkScopeId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePrivateLinkScopeID checks that 'input' can be parsed as a Private Link Scope ID
func ValidatePrivateLinkScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link Scope ID
func (id PrivateLinkScopeId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.HybridCompute/locations/%s/privateLinkScopes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.PrivateLinkScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link Scope ID
func (id PrivateLinkScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticPrivateLinkScopes", "privateLinkScopes", "privateLinkScopes"),
		resourceids.UserSpecifiedSegment("privateLinkScopeId", "privateLinkScopeIdValue"),
	}
}

// String returns a human-readable description of this Private Link Scope ID
func (id PrivateLinkScopeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Private Link Scope: %q", id.PrivateLinkScopeId),
	}
	return fmt.Sprintf("Private Link Scope (%s)", strings.Join(components, "\n"))
}
