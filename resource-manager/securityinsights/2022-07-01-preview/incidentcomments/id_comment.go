package incidentcomments

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
	WorkspaceName     string
	IncidentId        string
	IncidentCommentId string
}

// NewCommentID returns a new CommentId struct
func NewCommentID(subscriptionId string, resourceGroupName string, workspaceName string, incidentId string, incidentCommentId string) CommentId {
	return CommentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IncidentId:        incidentId,
		IncidentCommentId: incidentCommentId,
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

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentId' was not found in the resource id %q", input)
	}

	if id.IncidentCommentId, ok = parsed.Parsed["incidentCommentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentCommentId' was not found in the resource id %q", input)
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

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IncidentId, ok = parsed.Parsed["incidentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentId' was not found in the resource id %q", input)
	}

	if id.IncidentCommentId, ok = parsed.Parsed["incidentCommentId"]; !ok {
		return nil, fmt.Errorf("the segment 'incidentCommentId' was not found in the resource id %q", input)
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/incidents/%s/comments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IncidentId, id.IncidentCommentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Comment ID
func (id CommentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticIncidents", "incidents", "incidents"),
		resourceids.UserSpecifiedSegment("incidentId", "incidentIdValue"),
		resourceids.StaticSegment("staticComments", "comments", "comments"),
		resourceids.UserSpecifiedSegment("incidentCommentId", "incidentCommentIdValue"),
	}
}

// String returns a human-readable description of this Comment ID
func (id CommentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Incident: %q", id.IncidentId),
		fmt.Sprintf("Incident Comment: %q", id.IncidentCommentId),
	}
	return fmt.Sprintf("Comment (%s)", strings.Join(components, "\n"))
}
