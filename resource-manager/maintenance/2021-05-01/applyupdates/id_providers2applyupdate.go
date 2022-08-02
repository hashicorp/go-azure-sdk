package applyupdates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2ApplyUpdateId{}

// Providers2ApplyUpdateId is a struct representing the Resource ID for a Providers 2 Apply Update
type Providers2ApplyUpdateId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ProviderName       string
	ResourceParentType string
	ResourceParentName string
	ResourceType       string
	ResourceName       string
	ApplyUpdateName    string
}

// NewProviders2ApplyUpdateID returns a new Providers2ApplyUpdateId struct
func NewProviders2ApplyUpdateID(subscriptionId string, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, applyUpdateName string) Providers2ApplyUpdateId {
	return Providers2ApplyUpdateId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ProviderName:       providerName,
		ResourceParentType: resourceParentType,
		ResourceParentName: resourceParentName,
		ResourceType:       resourceType,
		ResourceName:       resourceName,
		ApplyUpdateName:    applyUpdateName,
	}
}

// ParseProviders2ApplyUpdateID parses 'input' into a Providers2ApplyUpdateId
func ParseProviders2ApplyUpdateID(input string) (*Providers2ApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ApplyUpdateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ApplyUpdateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceParentType, ok = parsed.Parsed["resourceParentType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceParentType' was not found in the resource id %q", input)
	}

	if id.ResourceParentName, ok = parsed.Parsed["resourceParentName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceParentName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.ApplyUpdateName, ok = parsed.Parsed["applyUpdateName"]; !ok {
		return nil, fmt.Errorf("the segment 'applyUpdateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviders2ApplyUpdateIDInsensitively parses 'input' case-insensitively into a Providers2ApplyUpdateId
// note: this method should only be used for API response data and not user input
func ParseProviders2ApplyUpdateIDInsensitively(input string) (*Providers2ApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ApplyUpdateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ApplyUpdateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceParentType, ok = parsed.Parsed["resourceParentType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceParentType' was not found in the resource id %q", input)
	}

	if id.ResourceParentName, ok = parsed.Parsed["resourceParentName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceParentName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.ApplyUpdateName, ok = parsed.Parsed["applyUpdateName"]; !ok {
		return nil, fmt.Errorf("the segment 'applyUpdateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviders2ApplyUpdateID checks that 'input' can be parsed as a Providers 2 Apply Update ID
func ValidateProviders2ApplyUpdateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2ApplyUpdateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Apply Update ID
func (id Providers2ApplyUpdateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/%s/%s/providers/Microsoft.Maintenance/applyUpdates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceParentType, id.ResourceParentName, id.ResourceType, id.ResourceName, id.ApplyUpdateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Apply Update ID
func (id Providers2ApplyUpdateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceParentType", "resourceParentTypeValue"),
		resourceids.UserSpecifiedSegment("resourceParentName", "resourceParentValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMaintenance", "Microsoft.Maintenance", "Microsoft.Maintenance"),
		resourceids.StaticSegment("staticApplyUpdates", "applyUpdates", "applyUpdates"),
		resourceids.UserSpecifiedSegment("applyUpdateName", "applyUpdateValue"),
	}
}

// String returns a human-readable description of this Providers 2 Apply Update ID
func (id Providers2ApplyUpdateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Parent Type: %q", id.ResourceParentType),
		fmt.Sprintf("Resource Parent Name: %q", id.ResourceParentName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Apply Update Name: %q", id.ApplyUpdateName),
	}
	return fmt.Sprintf("Providers 2 Apply Update (%s)", strings.Join(components, "\n"))
}
