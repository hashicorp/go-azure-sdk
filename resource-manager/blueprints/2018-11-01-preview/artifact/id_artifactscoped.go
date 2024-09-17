package artifact

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ArtifactScopedId{})
}

var _ resourceids.ResourceId = &ArtifactScopedId{}

// ArtifactScopedId is a struct representing the Resource ID for a Artifact Scoped
type ArtifactScopedId struct {
	ResourceScope string
	BlueprintName string
	ArtifactName  string
}

// NewArtifactScopedID returns a new ArtifactScopedId struct
func NewArtifactScopedID(resourceScope string, blueprintName string, artifactName string) ArtifactScopedId {
	return ArtifactScopedId{
		ResourceScope: resourceScope,
		BlueprintName: blueprintName,
		ArtifactName:  artifactName,
	}
}

// ParseArtifactScopedID parses 'input' into a ArtifactScopedId
func ParseArtifactScopedID(input string) (*ArtifactScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArtifactScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactScopedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseArtifactScopedIDInsensitively parses 'input' case-insensitively into a ArtifactScopedId
// note: this method should only be used for API response data and not user input
func ParseArtifactScopedIDInsensitively(input string) (*ArtifactScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArtifactScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactScopedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ArtifactScopedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceScope, ok = input.Parsed["resourceScope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", input)
	}

	if id.BlueprintName, ok = input.Parsed["blueprintName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "blueprintName", input)
	}

	if id.ArtifactName, ok = input.Parsed["artifactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "artifactName", input)
	}

	return nil
}

// ValidateArtifactScopedID checks that 'input' can be parsed as a Artifact Scoped ID
func ValidateArtifactScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseArtifactScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Artifact Scoped ID
func (id ArtifactScopedId) ID() string {
	fmtString := "/%s/providers/Microsoft.Blueprint/blueprints/%s/artifacts/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceScope, "/"), id.BlueprintName, id.ArtifactName)
}

// Segments returns a slice of Resource ID Segments which comprise this Artifact Scoped ID
func (id ArtifactScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceScope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBlueprint", "Microsoft.Blueprint", "Microsoft.Blueprint"),
		resourceids.StaticSegment("staticBlueprints", "blueprints", "blueprints"),
		resourceids.UserSpecifiedSegment("blueprintName", "blueprintValue"),
		resourceids.StaticSegment("staticArtifacts", "artifacts", "artifacts"),
		resourceids.UserSpecifiedSegment("artifactName", "artifactValue"),
	}
}

// String returns a human-readable description of this Artifact Scoped ID
func (id ArtifactScopedId) String() string {
	components := []string{
		fmt.Sprintf("Resource Scope: %q", id.ResourceScope),
		fmt.Sprintf("Blueprint Name: %q", id.BlueprintName),
		fmt.Sprintf("Artifact Name: %q", id.ArtifactName),
	}
	return fmt.Sprintf("Artifact Scoped (%s)", strings.Join(components, "\n"))
}
