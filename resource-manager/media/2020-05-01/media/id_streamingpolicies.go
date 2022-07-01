package media

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = StreamingPoliciesId{}

// StreamingPoliciesId is a struct representing the Resource ID for a Streaming Policies
type StreamingPoliciesId struct {
	SubscriptionId      string
	ResourceGroupName   string
	AccountName         string
	StreamingPolicyName string
}

// NewStreamingPoliciesID returns a new StreamingPoliciesId struct
func NewStreamingPoliciesID(subscriptionId string, resourceGroupName string, accountName string, streamingPolicyName string) StreamingPoliciesId {
	return StreamingPoliciesId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		AccountName:         accountName,
		StreamingPolicyName: streamingPolicyName,
	}
}

// ParseStreamingPoliciesID parses 'input' into a StreamingPoliciesId
func ParseStreamingPoliciesID(input string) (*StreamingPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(StreamingPoliciesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StreamingPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.StreamingPolicyName, ok = parsed.Parsed["streamingPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'streamingPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseStreamingPoliciesIDInsensitively parses 'input' case-insensitively into a StreamingPoliciesId
// note: this method should only be used for API response data and not user input
func ParseStreamingPoliciesIDInsensitively(input string) (*StreamingPoliciesId, error) {
	parser := resourceids.NewParserFromResourceIdType(StreamingPoliciesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StreamingPoliciesId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.StreamingPolicyName, ok = parsed.Parsed["streamingPolicyName"]; !ok {
		return nil, fmt.Errorf("the segment 'streamingPolicyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateStreamingPoliciesID checks that 'input' can be parsed as a Streaming Policies ID
func ValidateStreamingPoliciesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStreamingPoliciesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Streaming Policies ID
func (id StreamingPoliciesId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Media/mediaServices/%s/streamingPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.StreamingPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Streaming Policies ID
func (id StreamingPoliciesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMedia", "Microsoft.Media", "Microsoft.Media"),
		resourceids.StaticSegment("staticMediaServices", "mediaServices", "mediaServices"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticStreamingPolicies", "streamingPolicies", "streamingPolicies"),
		resourceids.UserSpecifiedSegment("streamingPolicyName", "streamingPolicyValue"),
	}
}

// String returns a human-readable description of this Streaming Policies ID
func (id StreamingPoliciesId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Streaming Policy Name: %q", id.StreamingPolicyName),
	}
	return fmt.Sprintf("Streaming Policies (%s)", strings.Join(components, "\n"))
}
