package virtualmachineruncommands

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineRunCommandId{}

// VirtualMachineRunCommandId is a struct representing the Resource ID for a Virtual Machine Run Command
type VirtualMachineRunCommandId struct {
	SubscriptionId    string
	ResourceGroupName string
	VmName            string
	RunCommandName    string
}

// NewVirtualMachineRunCommandID returns a new VirtualMachineRunCommandId struct
func NewVirtualMachineRunCommandID(subscriptionId string, resourceGroupName string, vmName string, runCommandName string) VirtualMachineRunCommandId {
	return VirtualMachineRunCommandId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VmName:            vmName,
		RunCommandName:    runCommandName,
	}
}

// ParseVirtualMachineRunCommandID parses 'input' into a VirtualMachineRunCommandId
func ParseVirtualMachineRunCommandID(input string) (*VirtualMachineRunCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineRunCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineRunCommandId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VmName, ok = parsed.Parsed["vmName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmName' was not found in the resource id %q", input)
	}

	if id.RunCommandName, ok = parsed.Parsed["runCommandName"]; !ok {
		return nil, fmt.Errorf("the segment 'runCommandName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineRunCommandIDInsensitively parses 'input' case-insensitively into a VirtualMachineRunCommandId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineRunCommandIDInsensitively(input string) (*VirtualMachineRunCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineRunCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineRunCommandId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VmName, ok = parsed.Parsed["vmName"]; !ok {
		return nil, fmt.Errorf("the segment 'vmName' was not found in the resource id %q", input)
	}

	if id.RunCommandName, ok = parsed.Parsed["runCommandName"]; !ok {
		return nil, fmt.Errorf("the segment 'runCommandName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVirtualMachineRunCommandID checks that 'input' can be parsed as a Virtual Machine Run Command ID
func ValidateVirtualMachineRunCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineRunCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Run Command ID
func (id VirtualMachineRunCommandId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s/runCommands/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VmName, id.RunCommandName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Run Command ID
func (id VirtualMachineRunCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("vmName", "vmValue"),
		resourceids.StaticSegment("staticRunCommands", "runCommands", "runCommands"),
		resourceids.UserSpecifiedSegment("runCommandName", "runCommandValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Run Command ID
func (id VirtualMachineRunCommandId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vm Name: %q", id.VmName),
		fmt.Sprintf("Run Command Name: %q", id.RunCommandName),
	}
	return fmt.Sprintf("Virtual Machine Run Command (%s)", strings.Join(components, "\n"))
}
