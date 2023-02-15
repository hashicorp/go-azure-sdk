package policyfragment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PolicyFragmentId{}

// PolicyFragmentId is a struct representing the Resource ID for a Policy Fragment
type PolicyFragmentId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ServiceName        string
	PolicyFragmentName string
}

// NewPolicyFragmentID returns a new PolicyFragmentId struct
func NewPolicyFragmentID(subscriptionId string, resourceGroupName string, serviceName string, policyFragmentName string) PolicyFragmentId {
	return PolicyFragmentId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ServiceName:        serviceName,
		PolicyFragmentName: policyFragmentName,
	}
}

// ParsePolicyFragmentID parses 'input' into a PolicyFragmentId
func ParsePolicyFragmentID(input string) (*PolicyFragmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyFragmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyFragmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.PolicyFragmentName, ok = parsed.Parsed["policyFragmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'policyFragmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePolicyFragmentIDInsensitively parses 'input' case-insensitively into a PolicyFragmentId
// note: this method should only be used for API response data and not user input
func ParsePolicyFragmentIDInsensitively(input string) (*PolicyFragmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyFragmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyFragmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.PolicyFragmentName, ok = parsed.Parsed["policyFragmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'policyFragmentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePolicyFragmentID checks that 'input' can be parsed as a Policy Fragment ID
func ValidatePolicyFragmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyFragmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Fragment ID
func (id PolicyFragmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/policyFragments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.PolicyFragmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Fragment ID
func (id PolicyFragmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticPolicyFragments", "policyFragments", "policyFragments"),
		resourceids.UserSpecifiedSegment("policyFragmentName", "policyFragmentValue"),
	}
}

// String returns a human-readable description of this Policy Fragment ID
func (id PolicyFragmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Policy Fragment Name: %q", id.PolicyFragmentName),
	}
	return fmt.Sprintf("Policy Fragment (%s)", strings.Join(components, "\n"))
}
