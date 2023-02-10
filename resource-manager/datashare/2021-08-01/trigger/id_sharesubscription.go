package trigger

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ShareSubscriptionId{}

// ShareSubscriptionId is a struct representing the Resource ID for a Share Subscription
type ShareSubscriptionId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AccountName           string
	ShareSubscriptionName string
}

// NewShareSubscriptionID returns a new ShareSubscriptionId struct
func NewShareSubscriptionID(subscriptionId string, resourceGroupName string, accountName string, shareSubscriptionName string) ShareSubscriptionId {
	return ShareSubscriptionId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AccountName:           accountName,
		ShareSubscriptionName: shareSubscriptionName,
	}
}

// ParseShareSubscriptionID parses 'input' into a ShareSubscriptionId
func ParseShareSubscriptionID(input string) (*ShareSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ShareSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ShareSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareSubscriptionName, ok = parsed.Parsed["shareSubscriptionName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareSubscriptionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseShareSubscriptionIDInsensitively parses 'input' case-insensitively into a ShareSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseShareSubscriptionIDInsensitively(input string) (*ShareSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ShareSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ShareSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareSubscriptionName, ok = parsed.Parsed["shareSubscriptionName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareSubscriptionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateShareSubscriptionID checks that 'input' can be parsed as a Share Subscription ID
func ValidateShareSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseShareSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Share Subscription ID
func (id ShareSubscriptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataShare/accounts/%s/shareSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ShareSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Share Subscription ID
func (id ShareSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticShareSubscriptions", "shareSubscriptions", "shareSubscriptions"),
		resourceids.UserSpecifiedSegment("shareSubscriptionName", "shareSubscriptionValue"),
	}
}

// String returns a human-readable description of this Share Subscription ID
func (id ShareSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Share Subscription Name: %q", id.ShareSubscriptionName),
	}
	return fmt.Sprintf("Share Subscription (%s)", strings.Join(components, "\n"))
}
