package addons

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = AddonId{}

// AddonId is a struct representing the Resource ID for a Addon
type AddonId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	RoleName              string
	AddonName             string
}

// NewAddonID returns a new AddonId struct
func NewAddonID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, roleName string, addonName string) AddonId {
	return AddonId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		RoleName:              roleName,
		AddonName:             addonName,
	}
}

// ParseAddonID parses 'input' into a AddonId
func ParseAddonID(input string) (*AddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(AddonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AddonId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.RoleName, ok = parsed.Parsed["roleName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleName' was not found in the resource id %q", input)
	}

	if id.AddonName, ok = parsed.Parsed["addonName"]; !ok {
		return nil, fmt.Errorf("the segment 'addonName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAddonIDInsensitively parses 'input' case-insensitively into a AddonId
// note: this method should only be used for API response data and not user input
func ParseAddonIDInsensitively(input string) (*AddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(AddonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AddonId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataBoxEdgeDeviceName' was not found in the resource id %q", input)
	}

	if id.RoleName, ok = parsed.Parsed["roleName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleName' was not found in the resource id %q", input)
	}

	if id.AddonName, ok = parsed.Parsed["addonName"]; !ok {
		return nil, fmt.Errorf("the segment 'addonName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAddonID checks that 'input' can be parsed as a Addon ID
func ValidateAddonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAddonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Addon ID
func (id AddonId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/roles/%s/addons/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.RoleName, id.AddonName)
}

// Segments returns a slice of Resource ID Segments which comprise this Addon ID
func (id AddonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticRoles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("roleName", "roleValue"),
		resourceids.StaticSegment("staticAddons", "addons", "addons"),
		resourceids.UserSpecifiedSegment("addonName", "addonValue"),
	}
}

// String returns a human-readable description of this Addon ID
func (id AddonId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Role Name: %q", id.RoleName),
		fmt.Sprintf("Addon Name: %q", id.AddonName),
	}
	return fmt.Sprintf("Addon (%s)", strings.Join(components, "\n"))
}
