package ipfirewallrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = FirewallRuleId{}

// FirewallRuleId is a struct representing the Resource ID for a Firewall Rule
type FirewallRuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	RuleName          string
}

// NewFirewallRuleID returns a new FirewallRuleId struct
func NewFirewallRuleID(subscriptionId string, resourceGroupName string, workspaceName string, ruleName string) FirewallRuleId {
	return FirewallRuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		RuleName:          ruleName,
	}
}

// ParseFirewallRuleID parses 'input' into a FirewallRuleId
func ParseFirewallRuleID(input string) (*FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(FirewallRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.RuleName, ok = parsed.Parsed["ruleName"]; !ok {
		return nil, fmt.Errorf("the segment 'ruleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseFirewallRuleIDInsensitively parses 'input' case-insensitively into a FirewallRuleId
// note: this method should only be used for API response data and not user input
func ParseFirewallRuleIDInsensitively(input string) (*FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(FirewallRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.RuleName, ok = parsed.Parsed["ruleName"]; !ok {
		return nil, fmt.Errorf("the segment 'ruleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateFirewallRuleID checks that 'input' can be parsed as a Firewall Rule ID
func ValidateFirewallRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFirewallRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Firewall Rule ID
func (id FirewallRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/firewallRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.RuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Firewall Rule ID
func (id FirewallRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticFirewallRules", "firewallRules", "firewallRules"),
		resourceids.UserSpecifiedSegment("ruleName", "ruleValue"),
	}
}

// String returns a human-readable description of this Firewall Rule ID
func (id FirewallRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Rule Name: %q", id.RuleName),
	}
	return fmt.Sprintf("Firewall Rule (%s)", strings.Join(components, "\n"))
}
