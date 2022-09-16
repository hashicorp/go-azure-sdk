package libraries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = LibraryId{}

// LibraryId is a struct representing the Resource ID for a Library
type LibraryId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	LibraryName       string
}

// NewLibraryID returns a new LibraryId struct
func NewLibraryID(subscriptionId string, resourceGroupName string, workspaceName string, libraryName string) LibraryId {
	return LibraryId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		LibraryName:       libraryName,
	}
}

// ParseLibraryID parses 'input' into a LibraryId
func ParseLibraryID(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(LibraryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LibraryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.LibraryName, ok = parsed.Parsed["libraryName"]; !ok {
		return nil, fmt.Errorf("the segment 'libraryName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseLibraryIDInsensitively parses 'input' case-insensitively into a LibraryId
// note: this method should only be used for API response data and not user input
func ParseLibraryIDInsensitively(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(LibraryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LibraryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.LibraryName, ok = parsed.Parsed["libraryName"]; !ok {
		return nil, fmt.Errorf("the segment 'libraryName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateLibraryID checks that 'input' can be parsed as a Library ID
func ValidateLibraryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLibraryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Library ID
func (id LibraryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/libraries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.LibraryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Library ID
func (id LibraryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticLibraries", "libraries", "libraries"),
		resourceids.UserSpecifiedSegment("libraryName", "libraryValue"),
	}
}

// String returns a human-readable description of this Library ID
func (id LibraryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Library Name: %q", id.LibraryName),
	}
	return fmt.Sprintf("Library (%s)", strings.Join(components, "\n"))
}
