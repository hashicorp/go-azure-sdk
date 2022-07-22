package capabilities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CapabilitiesId{}

// CapabilitiesId is a struct representing the Resource ID for a Capabilities
type CapabilitiesId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ParentProviderNamespace string
	ParentResourceType      string
	ParentResourceName      string
	TargetName              string
	CapabilityName          string
}

// NewCapabilitiesID returns a new CapabilitiesId struct
func NewCapabilitiesID(subscriptionId string, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string) CapabilitiesId {
	return CapabilitiesId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ParentProviderNamespace: parentProviderNamespace,
		ParentResourceType:      parentResourceType,
		ParentResourceName:      parentResourceName,
		TargetName:              targetName,
		CapabilityName:          capabilityName,
	}
}

// ParseCapabilitiesID parses 'input' into a CapabilitiesId
func ParseCapabilitiesID(input string) (*CapabilitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilitiesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilitiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ParentProviderNamespace, ok = parsed.Parsed["parentProviderNamespace"]; !ok {
		return nil, fmt.Errorf("the segment 'parentProviderNamespace' was not found in the resource id %q", input)
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

// ParseCapabilitiesIDInsensitively parses 'input' case-insensitively into a CapabilitiesId
// note: this method should only be used for API response data and not user input
func ParseCapabilitiesIDInsensitively(input string) (*CapabilitiesId, error) {
	parser := resourceids.NewParserFromResourceIdType(CapabilitiesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CapabilitiesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ParentProviderNamespace, ok = parsed.Parsed["parentProviderNamespace"]; !ok {
		return nil, fmt.Errorf("the segment 'parentProviderNamespace' was not found in the resource id %q", input)
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

// ValidateCapabilitiesID checks that 'input' can be parsed as a Capabilities ID
func ValidateCapabilitiesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCapabilitiesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Capabilities ID
func (id CapabilitiesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Chaos/targets/%s/capabilities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ParentProviderNamespace, id.ParentResourceType, id.ParentResourceName, id.TargetName, id.CapabilityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Capabilities ID
func (id CapabilitiesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("parentProviderNamespace", "parentProviderNamespaceValue"),
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

// String returns a human-readable description of this Capabilities ID
func (id CapabilitiesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Parent Provider Namespace: %q", id.ParentProviderNamespace),
		fmt.Sprintf("Parent Resource Type: %q", id.ParentResourceType),
		fmt.Sprintf("Parent Resource Name: %q", id.ParentResourceName),
		fmt.Sprintf("Target Name: %q", id.TargetName),
		fmt.Sprintf("Capability Name: %q", id.CapabilityName),
	}
	return fmt.Sprintf("Capabilities (%s)", strings.Join(components, "\n"))
}
