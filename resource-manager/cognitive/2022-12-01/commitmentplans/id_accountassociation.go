package commitmentplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = AccountAssociationId{}

// AccountAssociationId is a struct representing the Resource ID for a Account Association
type AccountAssociationId struct {
	SubscriptionId                string
	ResourceGroupName             string
	CommitmentPlanName            string
	CommitmentPlanAssociationName string
}

// NewAccountAssociationID returns a new AccountAssociationId struct
func NewAccountAssociationID(subscriptionId string, resourceGroupName string, commitmentPlanName string, commitmentPlanAssociationName string) AccountAssociationId {
	return AccountAssociationId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		CommitmentPlanName:            commitmentPlanName,
		CommitmentPlanAssociationName: commitmentPlanAssociationName,
	}
}

// ParseAccountAssociationID parses 'input' into a AccountAssociationId
func ParseAccountAssociationID(input string) (*AccountAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountAssociationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountAssociationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanAssociationName, ok = parsed.Parsed["commitmentPlanAssociationName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanAssociationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAccountAssociationIDInsensitively parses 'input' case-insensitively into a AccountAssociationId
// note: this method should only be used for API response data and not user input
func ParseAccountAssociationIDInsensitively(input string) (*AccountAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountAssociationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountAssociationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanAssociationName, ok = parsed.Parsed["commitmentPlanAssociationName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanAssociationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAccountAssociationID checks that 'input' can be parsed as a Account Association ID
func ValidateAccountAssociationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccountAssociationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Account Association ID
func (id AccountAssociationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/commitmentPlans/%s/accountAssociations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CommitmentPlanName, id.CommitmentPlanAssociationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Account Association ID
func (id AccountAssociationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticCommitmentPlans", "commitmentPlans", "commitmentPlans"),
		resourceids.UserSpecifiedSegment("commitmentPlanName", "commitmentPlanValue"),
		resourceids.StaticSegment("staticAccountAssociations", "accountAssociations", "accountAssociations"),
		resourceids.UserSpecifiedSegment("commitmentPlanAssociationName", "commitmentPlanAssociationValue"),
	}
}

// String returns a human-readable description of this Account Association ID
func (id AccountAssociationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Commitment Plan Name: %q", id.CommitmentPlanName),
		fmt.Sprintf("Commitment Plan Association Name: %q", id.CommitmentPlanAssociationName),
	}
	return fmt.Sprintf("Account Association (%s)", strings.Join(components, "\n"))
}
