package managementpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagementPolicyId{}

// ManagementPolicyId is a struct representing the Resource ID for a Management Policy
type ManagementPolicyId struct {
	SubscriptionId       string
	ResourceGroupName    string
	AccountName          string
	ManagementPolicyName ManagementPolicyName
}

// NewManagementPolicyID returns a new ManagementPolicyId struct
func NewManagementPolicyID(subscriptionId string, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName) ManagementPolicyId {
	return ManagementPolicyId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		AccountName:          accountName,
		ManagementPolicyName: managementPolicyName,
	}
}

// ParseManagementPolicyID parses 'input' into a ManagementPolicyId
func ParseManagementPolicyID(input string) (*ManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagementPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["managementPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'managementPolicyName' was not found in the resource id %q", input)
		}

		managementPolicyName, err := parseManagementPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ManagementPolicyName = *managementPolicyName
	}

	return &id, nil
}

// ParseManagementPolicyIDInsensitively parses 'input' case-insensitively into a ManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParseManagementPolicyIDInsensitively(input string) (*ManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagementPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["managementPolicyName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'managementPolicyName' was not found in the resource id %q", input)
		}

		managementPolicyName, err := parseManagementPolicyName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ManagementPolicyName = *managementPolicyName
	}

	return &id, nil
}

// ValidateManagementPolicyID checks that 'input' can be parsed as a Management Policy ID
func ValidateManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Management Policy ID
func (id ManagementPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/managementPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, string(id.ManagementPolicyName))
}

// Segments returns a slice of Resource ID Segments which comprise this Management Policy ID
func (id ManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorage", "Microsoft.Storage", "Microsoft.Storage"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticManagementPolicies", "managementPolicies", "managementPolicies"),
		resourceids.ConstantSegment("managementPolicyName", PossibleValuesForManagementPolicyName(), "default"),
	}
}

// String returns a human-readable description of this Management Policy ID
func (id ManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Management Policy Name: %q", string(id.ManagementPolicyName)),
	}
	return fmt.Sprintf("Management Policy (%s)", strings.Join(components, "\n"))
}
