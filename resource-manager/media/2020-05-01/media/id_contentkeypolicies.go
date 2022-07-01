package media

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ContentKeyPoliciesId{}

// ContentKeyPoliciesId is a struct representing the Resource ID for a Content Key Policies
type ContentKeyPoliciesId struct {
	SubscriptionId       string
	ResourceGroupName    string
	AccountName          string
	ContentKeyPolicyName string
}

// NewContentKeyPoliciesID returns a new ContentKeyPoliciesId struct
func NewContentKeyPoliciesID(subscriptionId string, resourceGroupName string, accountName string, contentKeyPolicyName string) ContentKeyPoliciesId {
	return ContentKeyPoliciesId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		AccountName:          accountName,
		ContentKeyPolicyName: contentKeyPolicyName,
	}
}

// ParseContentKeyPoliciesID parses 'input' into a ContentKeyPoliciesId
func ParseContentKeyPoliciesID(input string) (*ContentKeyPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentKeyPoliciesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ContentKeyPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ContentKeyPolicyName, ok = parsed.Parsed["contentKeyPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'contentKeyPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseContentKeyPoliciesIDInsensitively parses 'input' case-insensitively into a ContentKeyPoliciesId
// note: this method should only be used for API response data and not user input
func ParseContentKeyPoliciesIDInsensitively(input string) (*ContentKeyPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(ContentKeyPoliciesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ContentKeyPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ContentKeyPolicyName, ok = parsed.Parsed["contentKeyPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'contentKeyPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateContentKeyPoliciesID checks that 'input' can be parsed as a Content Key Policies ID
func ValidateContentKeyPoliciesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentKeyPoliciesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Key Policies ID
func (id ContentKeyPoliciesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Media/mediaServices/%s/contentKeyPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ContentKeyPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Key Policies ID
func (id ContentKeyPoliciesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMedia", "Microsoft.Media", "Microsoft.Media"),
		resourceids.StaticSegment("staticMediaServices", "mediaServices", "mediaServices"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticContentKeyPolicies", "contentKeyPolicies", "contentKeyPolicies"),
		resourceids.UserSpecifiedSegment("contentKeyPolicyName", "contentKeyPolicyValue"),
	}
}

// String returns a human-readable description of this Content Key Policies ID
func (id ContentKeyPoliciesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Content Key Policy Name: %q", id.ContentKeyPolicyName),
	}
	return fmt.Sprintf("Content Key Policies (%s)", strings.Join(components, "\n"))
}
