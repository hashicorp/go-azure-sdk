package hybrididentitymetadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = HybridIdentityMetadataId{}

// HybridIdentityMetadataId is a struct representing the Resource ID for a Hybrid Identity Metadata
type HybridIdentityMetadataId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VirtualMachineName string
	MetadataName       string
}

// NewHybridIdentityMetadataID returns a new HybridIdentityMetadataId struct
func NewHybridIdentityMetadataID(subscriptionId string, resourceGroupName string, virtualMachineName string, metadataName string) HybridIdentityMetadataId {
	return HybridIdentityMetadataId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VirtualMachineName: virtualMachineName,
		MetadataName:       metadataName,
	}
}

// ParseHybridIdentityMetadataID parses 'input' into a HybridIdentityMetadataId
func ParseHybridIdentityMetadataID(input string) (*HybridIdentityMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridIdentityMetadataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridIdentityMetadataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.MetadataName, ok = parsed.Parsed["metadataName"]; !ok {
		return nil, fmt.Errorf("the segment 'metadataName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseHybridIdentityMetadataIDInsensitively parses 'input' case-insensitively into a HybridIdentityMetadataId
// note: this method should only be used for API response data and not user input
func ParseHybridIdentityMetadataIDInsensitively(input string) (*HybridIdentityMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridIdentityMetadataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridIdentityMetadataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.MetadataName, ok = parsed.Parsed["metadataName"]; !ok {
		return nil, fmt.Errorf("the segment 'metadataName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateHybridIdentityMetadataID checks that 'input' can be parsed as a Hybrid Identity Metadata ID
func ValidateHybridIdentityMetadataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHybridIdentityMetadataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s/hybridIdentityMetadata/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.MetadataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticHybridIdentityMetadata", "hybridIdentityMetadata", "hybridIdentityMetadata"),
		resourceids.UserSpecifiedSegment("metadataName", "metadataValue"),
	}
}

// String returns a human-readable description of this Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Metadata Name: %q", id.MetadataName),
	}
	return fmt.Sprintf("Hybrid Identity Metadata (%s)", strings.Join(components, "\n"))
}
