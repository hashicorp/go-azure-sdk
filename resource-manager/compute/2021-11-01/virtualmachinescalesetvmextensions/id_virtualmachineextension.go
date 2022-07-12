package virtualmachinescalesetvmextensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineExtensionId{}

// VirtualMachineExtensionId is a struct representing the Resource ID for a Virtual Machine Extension
type VirtualMachineExtensionId struct {
	SubscriptionId    string
	ResourceGroupName string
	VmScaleSetName    string
	InstanceId        string
	VmExtensionName   string
}

// NewVirtualMachineExtensionID returns a new VirtualMachineExtensionId struct
func NewVirtualMachineExtensionID(subscriptionId string, resourceGroupName string, vmScaleSetName string, instanceId string, vmExtensionName string) VirtualMachineExtensionId {
	return VirtualMachineExtensionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VmScaleSetName:    vmScaleSetName,
		InstanceId:        instanceId,
		VmExtensionName:   vmExtensionName,
	}
}

// ParseVirtualMachineExtensionID parses 'input' into a VirtualMachineExtensionId
func ParseVirtualMachineExtensionID(input string) (*VirtualMachineExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineExtensionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VmScaleSetName, ok = parsed.Parsed["vmScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmScaleSetName' was not found in the resource id %q", input)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, fmt.Errorf("the segment 'instanceId' was not found in the resource id %q", input)
	}

	if id.VmExtensionName, ok = parsed.Parsed["vmExtensionName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmExtensionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineExtensionIDInsensitively parses 'input' case-insensitively into a VirtualMachineExtensionId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineExtensionIDInsensitively(input string) (*VirtualMachineExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineExtensionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VmScaleSetName, ok = parsed.Parsed["vmScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmScaleSetName' was not found in the resource id %q", input)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, fmt.Errorf("the segment 'instanceId' was not found in the resource id %q", input)
	}

	if id.VmExtensionName, ok = parsed.Parsed["vmExtensionName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmExtensionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVirtualMachineExtensionID checks that 'input' can be parsed as a Virtual Machine Extension ID
func ValidateVirtualMachineExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Extension ID
func (id VirtualMachineExtensionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VmScaleSetName, id.InstanceId, id.VmExtensionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Extension ID
func (id VirtualMachineExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("vmScaleSetName", "vmScaleSetValue"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("instanceId", "instanceIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("vmExtensionName", "vmExtensionValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Extension ID
func (id VirtualMachineExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vm Scale Set Name: %q", id.VmScaleSetName),
		fmt.Sprintf("Instance: %q", id.InstanceId),
		fmt.Sprintf("Vm Extension Name: %q", id.VmExtensionName),
	}
	return fmt.Sprintf("Virtual Machine Extension (%s)", strings.Join(components, "\n"))
}
