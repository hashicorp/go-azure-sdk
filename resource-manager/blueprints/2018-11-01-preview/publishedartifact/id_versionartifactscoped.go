package publishedartifact

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VersionArtifactScopedId{}

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
	parser := resourceids.NewParserFromResourceIdType(VersionArtifactScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionArtifactScopedId{}

	if id.ResourceScope, ok = parsed.Parsed["resourceScope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", *parsed)
	}

	if id.BlueprintName, ok = parsed.Parsed["blueprintName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "blueprintName", *parsed)
	}

	if id.VersionId, ok = parsed.Parsed["versionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionId", *parsed)
	}

	if id.ArtifactName, ok = parsed.Parsed["artifactName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "artifactName", *parsed)
	}

	return &id, nil
}

// ParseVersionArtifactScopedIDInsensitively parses 'input' case-insensitively into a VersionArtifactScopedId
// note: this method should only be used for API response data and not user input
func ParseVersionArtifactScopedIDInsensitively(input string) (*VersionArtifactScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionArtifactScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionArtifactScopedId{}

	if id.ResourceScope, ok = parsed.Parsed["resourceScope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceScope", *parsed)
	}

	if id.BlueprintName, ok = parsed.Parsed["blueprintName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "blueprintName", *parsed)
	}

	if id.VersionId, ok = parsed.Parsed["versionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionId", *parsed)
	}

	if id.ArtifactName, ok = parsed.Parsed["artifactName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "artifactName", *parsed)
	}

	return &id, nil
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
