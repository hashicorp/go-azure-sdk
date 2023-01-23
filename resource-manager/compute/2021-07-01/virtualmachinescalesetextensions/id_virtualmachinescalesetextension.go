package virtualmachinescalesetextensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineScaleSetExtensionId{}

// VirtualMachineScaleSetExtensionId is a struct representing the Resource ID for a Virtual Machine Scale Set Extension
type VirtualMachineScaleSetExtensionId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VirtualMachineScaleSetName string
	ExtensionName              string
}

// NewVirtualMachineScaleSetExtensionID returns a new VirtualMachineScaleSetExtensionId struct
func NewVirtualMachineScaleSetExtensionID(subscriptionId string, resourceGroupName string, virtualMachineScaleSetName string, extensionName string) VirtualMachineScaleSetExtensionId {
	return VirtualMachineScaleSetExtensionId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VirtualMachineScaleSetName: virtualMachineScaleSetName,
		ExtensionName:              extensionName,
	}
}

// ParseVirtualMachineScaleSetExtensionID parses 'input' into a VirtualMachineScaleSetExtensionId
func ParseVirtualMachineScaleSetExtensionID(input string) (*VirtualMachineScaleSetExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetExtensionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineScaleSetName' was not found in the resource id %q", input)
	}

	if id.ExtensionName, ok = parsed.Parsed["extensionName"]; !ok {
		return nil, fmt.Errorf("the segment 'extensionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineScaleSetExtensionIDInsensitively parses 'input' case-insensitively into a VirtualMachineScaleSetExtensionId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScaleSetExtensionIDInsensitively(input string) (*VirtualMachineScaleSetExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetExtensionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineScaleSetName' was not found in the resource id %q", input)
	}

	if id.ExtensionName, ok = parsed.Parsed["extensionName"]; !ok {
		return nil, fmt.Errorf("the segment 'extensionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVirtualMachineScaleSetExtensionID checks that 'input' can be parsed as a Virtual Machine Scale Set Extension ID
func ValidateVirtualMachineScaleSetExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScaleSetExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Scale Set Extension ID
func (id VirtualMachineScaleSetExtensionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineScaleSetName, id.ExtensionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Scale Set Extension ID
func (id VirtualMachineScaleSetExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionName", "extensionValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Scale Set Extension ID
func (id VirtualMachineScaleSetExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Scale Set Name: %q", id.VirtualMachineScaleSetName),
		fmt.Sprintf("Extension Name: %q", id.ExtensionName),
	}
	return fmt.Sprintf("Virtual Machine Scale Set Extension (%s)", strings.Join(components, "\n"))
}
