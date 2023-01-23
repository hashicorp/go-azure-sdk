package capabilities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CapabilityId{}

// CapabilityId is a struct representing the Resource ID for a Capability
type CapabilityId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ProviderName       string
	ParentResourceType string
	ParentResourceName string
	TargetName         string
	CapabilityName     string
}

// NewCapabilityID returns a new CapabilityId struct
func NewCapabilityID(subscriptionId string, resourceGroupName string, providerName string, parentResourceType string, parentResourceName string, targetName string, capabilityName string) CapabilityId {
	return CapabilityId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ProviderName:       providerName,
		ParentResourceType: parentResourceType,
		ParentResourceName: parentResourceName,
		TargetName:         targetName,
		CapabilityName:     capabilityName,
	}
}

// ParseCapabilityID parses 'input' into a CapabilityId
func ParseCapabilityID(input string) (*CapabilityId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ParentResourceType, ok = parsed.Parsed["parentResourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'parentResourceType' was not found in the resource id %q", input)
	}

	if id.ParentResourceName, ok = parsed.Parsed["parentResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'parentResourceName' was not found in the resource id %q", input)
	}

	if id.TargetName, ok = parsed.Parsed["targetName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetName' was not found in the resource id %q", input)
	}

	if id.CapabilityName, ok = parsed.Parsed["capabilityName"]; !ok {
		return nil, fmt.Errorf("the segment 'capabilityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCapabilityIDInsensitively parses 'input' case-insensitively into a CapabilityId
// note: this method should only be used for API response data and not user input
func ParseCapabilityIDInsensitively(input string) (*CapabilityId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ParentResourceType, ok = parsed.Parsed["parentResourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'parentResourceType' was not found in the resource id %q", input)
	}

	if id.ParentResourceName, ok = parsed.Parsed["parentResourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'parentResourceName' was not found in the resource id %q", input)
	}

	if id.TargetName, ok = parsed.Parsed["targetName"]; !ok {
		return nil, fmt.Errorf("the segment 'targetName' was not found in the resource id %q", input)
	}

	if id.CapabilityName, ok = parsed.Parsed["capabilityName"]; !ok {
		return nil, fmt.Errorf("the segment 'capabilityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCapabilityID checks that 'input' can be parsed as a Capability ID
func ValidateCapabilityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCapabilityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Capability ID
func (id CapabilityId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Chaos/targets/%s/capabilities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ParentResourceType, id.ParentResourceName, id.TargetName, id.CapabilityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Capability ID
func (id CapabilityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("parentResourceType", "parentResourceTypeValue"),
		resourceids.UserSpecifiedSegment("parentResourceName", "parentResourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticTargets", "targets", "targets"),
		resourceids.UserSpecifiedSegment("targetName", "targetValue"),
		resourceids.StaticSegment("staticCapabilities", "capabilities", "capabilities"),
		resourceids.UserSpecifiedSegment("capabilityName", "capabilityValue"),
	}
}

// String returns a human-readable description of this Capability ID
func (id CapabilityId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Parent Resource Type: %q", id.ParentResourceType),
		fmt.Sprintf("Parent Resource Name: %q", id.ParentResourceName),
		fmt.Sprintf("Target Name: %q", id.TargetName),
		fmt.Sprintf("Capability Name: %q", id.CapabilityName),
	}
	return fmt.Sprintf("Capability (%s)", strings.Join(components, "\n"))
}
