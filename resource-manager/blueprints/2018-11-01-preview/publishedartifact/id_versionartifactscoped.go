package publishedartifact

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VersionArtifactScopedId{})
}

var _ resourceids.ResourceId = &VersionArtifactScopedId{}

// VersionArtifactScopedId is a struct representing the Resource ID for a Version Artifact Scoped
type VersionArtifactScopedId struct {
	ResourceScope string
	BlueprintName string
	VersionId     string
	ArtifactName  string
}

// NewVersionArtifactScopedID returns a new VersionArtifactScopedId struct
func NewVersionArtifactScopedID(resourceScope string, blueprintName string, versionId string, artifactName string) VersionArtifactScopedId {
	return VersionArtifactScopedId{
		ResourceScope: resourceScope,
		BlueprintName: blueprintName,
		VersionId:     versionId,
		ArtifactName:  artifactName,
	}
}

// ParseVersionArtifactScopedID parses 'input' into a VersionArtifactScopedId
func ParseVersionArtifactScopedID(input string) (*VersionArtifactScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VersionArtifactScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VersionArtifactScopedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVersionArtifactScopedIDInsensitively parses 'input' case-insensitively into a VersionArtifactScopedId
// note: this method should only be used for API response data and not user input
func ParseVersionArtifactScopedIDInsensitively(input string) (*VersionArtifactScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VersionArtifactScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VersionArtifactScopedId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VersionArtifactScopedId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceScope, ok = input.Parsed["resourceScope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", input)
	}

	if id.BlueprintName, ok = input.Parsed["blueprintName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "blueprintName", input)
	}

	if id.VersionId, ok = input.Parsed["versionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionId", input)
	}

	if id.ArtifactName, ok = input.Parsed["artifactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "artifactName", input)
	}

	return nil
}

// ValidateVersionArtifactScopedID checks that 'input' can be parsed as a Version Artifact Scoped ID
func ValidateVersionArtifactScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVersionArtifactScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Version Artifact Scoped ID
func (id VersionArtifactScopedId) ID() string {
	fmtString := "/%s/providers/Microsoft.Blueprint/blueprints/%s/versions/%s/artifacts/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceScope, "/"), id.BlueprintName, id.VersionId, id.ArtifactName)
}

// Segments returns a slice of Resource ID Segments which comprise this Version Artifact Scoped ID
func (id VersionArtifactScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceScope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBlueprint", "Microsoft.Blueprint", "Microsoft.Blueprint"),
		resourceids.StaticSegment("staticBlueprints", "blueprints", "blueprints"),
		resourceids.UserSpecifiedSegment("blueprintName", "blueprintValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionId", "versionIdValue"),
		resourceids.StaticSegment("staticArtifacts", "artifacts", "artifacts"),
		resourceids.UserSpecifiedSegment("artifactName", "artifactValue"),
	}
}

// String returns a human-readable description of this Version Artifact Scoped ID
func (id VersionArtifactScopedId) String() string {
	components := []string{
		fmt.Sprintf("Resource Scope: %q", id.ResourceScope),
		fmt.Sprintf("Blueprint Name: %q", id.BlueprintName),
		fmt.Sprintf("Version: %q", id.VersionId),
		fmt.Sprintf("Artifact Name: %q", id.ArtifactName),
	}
	return fmt.Sprintf("Version Artifact Scoped (%s)", strings.Join(components, "\n"))
}
