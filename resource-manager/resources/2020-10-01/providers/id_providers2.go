package providers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2Id{}

// Providers2Id is a struct representing the Resource ID for a Providers 2
type Providers2Id struct {
	GroupId                   string
	ResourceProviderNamespace string
}

// NewProviders2ID returns a new Providers2Id struct
func NewProviders2ID(groupId string, resourceProviderNamespace string) Providers2Id {
	return Providers2Id{
		GroupId:                   groupId,
		ResourceProviderNamespace: resourceProviderNamespace,
	}
}

// ParseProviders2ID parses 'input' into a Providers2Id
func ParseProviders2ID(input string) (*Providers2Id, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2Id{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2Id{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, fmt.Errorf("the segment 'groupId' was not found in the resource id %q", input)
	}

	if id.ResourceProviderNamespace, ok = parsed.Parsed["resourceProviderNamespace"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceProviderNamespace' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviders2IDInsensitively parses 'input' case-insensitively into a Providers2Id
// note: this method should only be used for API response data and not user input
func ParseProviders2IDInsensitively(input string) (*Providers2Id, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2Id{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2Id{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, fmt.Errorf("the segment 'groupId' was not found in the resource id %q", input)
	}

	if id.ResourceProviderNamespace, ok = parsed.Parsed["resourceProviderNamespace"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceProviderNamespace' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviders2ID checks that 'input' can be parsed as a Providers 2 ID
func ValidateProviders2ID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2ID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 ID
func (id Providers2Id) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ResourceProviderNamespace)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 ID
func (id Providers2Id) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.UserSpecifiedSegment("resourceProviderNamespace", "resourceProviderNamespaceValue"),
	}
}

// String returns a human-readable description of this Providers 2 ID
func (id Providers2Id) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Resource Provider Namespace: %q", id.ResourceProviderNamespace),
	}
	return fmt.Sprintf("Providers 2 (%s)", strings.Join(components, "\n"))
}
