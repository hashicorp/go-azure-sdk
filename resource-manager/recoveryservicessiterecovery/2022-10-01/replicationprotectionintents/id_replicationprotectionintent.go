package replicationprotectionintents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationProtectionIntentId{}

// ReplicationProtectionIntentId is a struct representing the Resource ID for a Replication Protection Intent
type ReplicationProtectionIntentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	IntentObjectName  string
}

// NewReplicationProtectionIntentID returns a new ReplicationProtectionIntentId struct
func NewReplicationProtectionIntentID(subscriptionId string, resourceGroupName string, resourceName string, intentObjectName string) ReplicationProtectionIntentId {
	return ReplicationProtectionIntentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
		IntentObjectName:  intentObjectName,
	}
}

// ParseReplicationProtectionIntentID parses 'input' into a ReplicationProtectionIntentId
func ParseReplicationProtectionIntentID(input string) (*ReplicationProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationProtectionIntentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationProtectionIntentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.IntentObjectName, ok = parsed.Parsed["intentObjectName"]; !ok {
		return nil, fmt.Errorf("the segment 'intentObjectName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationProtectionIntentIDInsensitively parses 'input' case-insensitively into a ReplicationProtectionIntentId
// note: this method should only be used for API response data and not user input
func ParseReplicationProtectionIntentIDInsensitively(input string) (*ReplicationProtectionIntentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationProtectionIntentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationProtectionIntentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.IntentObjectName, ok = parsed.Parsed["intentObjectName"]; !ok {
		return nil, fmt.Errorf("the segment 'intentObjectName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationProtectionIntentID checks that 'input' can be parsed as a Replication Protection Intent ID
func ValidateReplicationProtectionIntentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationProtectionIntentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Protection Intent ID
func (id ReplicationProtectionIntentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationProtectionIntents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.IntentObjectName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Protection Intent ID
func (id ReplicationProtectionIntentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticReplicationProtectionIntents", "replicationProtectionIntents", "replicationProtectionIntents"),
		resourceids.UserSpecifiedSegment("intentObjectName", "intentObjectValue"),
	}
}

// String returns a human-readable description of this Replication Protection Intent ID
func (id ReplicationProtectionIntentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Intent Object Name: %q", id.IntentObjectName),
	}
	return fmt.Sprintf("Replication Protection Intent (%s)", strings.Join(components, "\n"))
}
