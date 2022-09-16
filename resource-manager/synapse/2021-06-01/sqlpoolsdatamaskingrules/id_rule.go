package sqlpoolsdatamaskingrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RuleId{}

// RuleId is a struct representing the Resource ID for a Rule
type RuleId struct {
	SubscriptionId      string
	ResourceGroupName   string
	WorkspaceName       string
	SqlPoolName         string
	DataMaskingRuleName string
}

// NewRuleID returns a new RuleId struct
func NewRuleID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string, dataMaskingRuleName string) RuleId {
	return RuleId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		WorkspaceName:       workspaceName,
		SqlPoolName:         sqlPoolName,
		DataMaskingRuleName: dataMaskingRuleName,
	}
}

// ParseRuleID parses 'input' into a RuleId
func ParseRuleID(input string) (*RuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SqlPoolName, ok = parsed.Parsed["sqlPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'sqlPoolName' was not found in the resource id %q", input)
	}

	if id.DataMaskingRuleName, ok = parsed.Parsed["dataMaskingRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataMaskingRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRuleIDInsensitively parses 'input' case-insensitively into a RuleId
// note: this method should only be used for API response data and not user input
func ParseRuleIDInsensitively(input string) (*RuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SqlPoolName, ok = parsed.Parsed["sqlPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'sqlPoolName' was not found in the resource id %q", input)
	}

	if id.DataMaskingRuleName, ok = parsed.Parsed["dataMaskingRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'dataMaskingRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRuleID checks that 'input' can be parsed as a Rule ID
func ValidateRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rule ID
func (id RuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/sqlPools/%s/dataMaskingPolicies/default/rules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName, id.DataMaskingRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rule ID
func (id RuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticSqlPools", "sqlPools", "sqlPools"),
		resourceids.UserSpecifiedSegment("sqlPoolName", "sqlPoolValue"),
		resourceids.StaticSegment("staticDataMaskingPolicies", "dataMaskingPolicies", "dataMaskingPolicies"),
		resourceids.StaticSegment("dataMaskingPolicyName", "default", "default"),
		resourceids.StaticSegment("staticRules", "rules", "rules"),
		resourceids.UserSpecifiedSegment("dataMaskingRuleName", "dataMaskingRuleValue"),
	}
}

// String returns a human-readable description of this Rule ID
func (id RuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
		fmt.Sprintf("Data Masking Rule Name: %q", id.DataMaskingRuleName),
	}
	return fmt.Sprintf("Rule (%s)", strings.Join(components, "\n"))
}
