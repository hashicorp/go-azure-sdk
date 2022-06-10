package blobcontainers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ImmutabilityPolicyId{}

// ImmutabilityPolicyId is a struct representing the Resource ID for a Immutability Policy
type ImmutabilityPolicyId struct {
	SubscriptionId         string
	ResourceGroupName      string
	AccountName            string
	ContainerName          string
	ImmutabilityPolicyName ImmutabilityPolicyName
}

// NewImmutabilityPolicyID returns a new ImmutabilityPolicyId struct
func NewImmutabilityPolicyID(subscriptionId string, resourceGroupName string, accountName string, containerName string, immutabilityPolicyName ImmutabilityPolicyName) ImmutabilityPolicyId {
	return ImmutabilityPolicyId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		AccountName:            accountName,
		ContainerName:          containerName,
		ImmutabilityPolicyName: immutabilityPolicyName,
	}
}

// ParseImmutabilityPolicyID parses 'input' into a ImmutabilityPolicyId
func ParseImmutabilityPolicyID(input string) (*ImmutabilityPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ImmutabilityPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ImmutabilityPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ContainerName, ok = parsed.Parsed["containerName"]; !ok {
		return nil, fmt.Errorf("the segment 'containerName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["immutabilityPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'immutabilityPolicyName' was not found in the resource id %q", input)
		}

		immutabilityPolicyName, err := parseImmutabilityPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ImmutabilityPolicyName = *immutabilityPolicyName
	}

	return &id, nil
}

// ParseImmutabilityPolicyIDInsensitively parses 'input' case-insensitively into a ImmutabilityPolicyId
// note: this method should only be used for API response data and not user input
func ParseImmutabilityPolicyIDInsensitively(input string) (*ImmutabilityPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ImmutabilityPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ImmutabilityPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ContainerName, ok = parsed.Parsed["containerName"]; !ok {
		return nil, fmt.Errorf("the segment 'containerName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["immutabilityPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'immutabilityPolicyName' was not found in the resource id %q", input)
		}

		immutabilityPolicyName, err := parseImmutabilityPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ImmutabilityPolicyName = *immutabilityPolicyName
	}

	return &id, nil
}

// ValidateImmutabilityPolicyID checks that 'input' can be parsed as a Immutability Policy ID
func ValidateImmutabilityPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseImmutabilityPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Immutability Policy ID
func (id ImmutabilityPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s/immutabilityPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ContainerName, string(id.ImmutabilityPolicyName))
}

// Segments returns a slice of Resource ID Segments which comprise this Immutability Policy ID
func (id ImmutabilityPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticBlobServices", "blobServices", "blobServices"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticContainers", "containers", "containers"),
		resourceids.UserSpecifiedSegment("containerName", "containerValue"),
		resourceids.StaticSegment("staticImmutabilityPolicies", "immutabilityPolicies", "immutabilityPolicies"),
		resourceids.ConstantSegment("immutabilityPolicyName", PossibleValuesForImmutabilityPolicyName(), "default"),
	}
}

// String returns a human-readable description of this Immutability Policy ID
func (id ImmutabilityPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Container Name: %q", id.ContainerName),
		fmt.Sprintf("Immutability Policy Name: %q", string(id.ImmutabilityPolicyName)),
	}
	return fmt.Sprintf("Immutability Policy (%s)", strings.Join(components, "\n"))
}
