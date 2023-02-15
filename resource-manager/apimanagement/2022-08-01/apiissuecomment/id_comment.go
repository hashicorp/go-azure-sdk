package apiissuecomment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CommentId{}

// CommentId is a struct representing the Resource ID for a Comment
type CommentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ApiId             string
	IssueId           string
	CommentId         string
}

// NewCommentID returns a new CommentId struct
func NewCommentID(subscriptionId string, resourceGroupName string, serviceName string, apiId string, issueId string, commentId string) CommentId {
	return CommentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ApiId:             apiId,
		IssueId:           issueId,
		CommentId:         commentId,
	}
}

// ParseCommentID parses 'input' into a CommentId
func ParseCommentID(input string) (*CommentId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	if id.CommentId, ok = parsed.Parsed["commentId"]; !ok {
		return nil, fmt.Errorf("the segment 'commentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCommentIDInsensitively parses 'input' case-insensitively into a CommentId
// note: this method should only be used for API response data and not user input
func ParseCommentIDInsensitively(input string) (*CommentId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, fmt.Errorf("the segment 'apiId' was not found in the resource id %q", input)
	}

	if id.IssueId, ok = parsed.Parsed["issueId"]; !ok {
		return nil, fmt.Errorf("the segment 'issueId' was not found in the resource id %q", input)
	}

	if id.CommentId, ok = parsed.Parsed["commentId"]; !ok {
		return nil, fmt.Errorf("the segment 'commentId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCommentID checks that 'input' can be parsed as a Comment ID
func ValidateCommentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Comment ID
func (id CommentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/apis/%s/issues/%s/comments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ApiId, id.IssueId, id.CommentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Comment ID
func (id CommentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticApis", "apis", "apis"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
		resourceids.StaticSegment("staticIssues", "issues", "issues"),
		resourceids.UserSpecifiedSegment("issueId", "issueIdValue"),
		resourceids.StaticSegment("staticComments", "comments", "comments"),
		resourceids.UserSpecifiedSegment("commentId", "commentIdValue"),
	}
}

// String returns a human-readable description of this Comment ID
func (id CommentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Api: %q", id.ApiId),
		fmt.Sprintf("Issue: %q", id.IssueId),
		fmt.Sprintf("Comment: %q", id.CommentId),
	}
	return fmt.Sprintf("Comment (%s)", strings.Join(components, "\n"))
}
