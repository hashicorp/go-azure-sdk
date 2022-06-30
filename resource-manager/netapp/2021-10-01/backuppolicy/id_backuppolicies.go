package backuppolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BackupPoliciesId{}

// BackupPoliciesId is a struct representing the Resource ID for a Backup Policies
type BackupPoliciesId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	BackupPolicyName  string
}

// NewBackupPoliciesID returns a new BackupPoliciesId struct
func NewBackupPoliciesID(subscriptionId string, resourceGroupName string, accountName string, backupPolicyName string) BackupPoliciesId {
	return BackupPoliciesId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		BackupPolicyName:  backupPolicyName,
	}
}

// ParseBackupPoliciesID parses 'input' into a BackupPoliciesId
func ParseBackupPoliciesID(input string) (*BackupPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupPoliciesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.BackupPolicyName, ok = parsed.Parsed["backupPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseBackupPoliciesIDInsensitively parses 'input' case-insensitively into a BackupPoliciesId
// note: this method should only be used for API response data and not user input
func ParseBackupPoliciesIDInsensitively(input string) (*BackupPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupPoliciesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BackupPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.BackupPolicyName, ok = parsed.Parsed["backupPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'backupPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateBackupPoliciesID checks that 'input' can be parsed as a Backup Policies ID
func ValidateBackupPoliciesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupPoliciesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup Policies ID
func (id BackupPoliciesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/backupPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.BackupPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup Policies ID
func (id BackupPoliciesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticBackupPolicies", "backupPolicies", "backupPolicies"),
		resourceids.UserSpecifiedSegment("backupPolicyName", "backupPolicyValue"),
	}
}

// String returns a human-readable description of this Backup Policies ID
func (id BackupPoliciesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Backup Policy Name: %q", id.BackupPolicyName),
	}
	return fmt.Sprintf("Backup Policies (%s)", strings.Join(components, "\n"))
}
