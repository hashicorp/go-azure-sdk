package codeversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RegistryCodeVersionId{})
}

var _ resourceids.ResourceId = &RegistryCodeVersionId{}

// RegistryCodeVersionId is a struct representing the Resource ID for a Registry Code Version
type RegistryCodeVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	RegistryName      string
	CodeName          string
	VersionName       string
}

// NewRegistryCodeVersionID returns a new RegistryCodeVersionId struct
func NewRegistryCodeVersionID(subscriptionId string, resourceGroupName string, registryName string, codeName string, versionName string) RegistryCodeVersionId {
	return RegistryCodeVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RegistryName:      registryName,
		CodeName:          codeName,
		VersionName:       versionName,
	}
}

// ParseRegistryCodeVersionID parses 'input' into a RegistryCodeVersionId
func ParseRegistryCodeVersionID(input string) (*RegistryCodeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryCodeVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryCodeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRegistryCodeVersionIDInsensitively parses 'input' case-insensitively into a RegistryCodeVersionId
// note: this method should only be used for API response data and not user input
func ParseRegistryCodeVersionIDInsensitively(input string) (*RegistryCodeVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryCodeVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryCodeVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegistryCodeVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RegistryName, ok = input.Parsed["registryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryName", input)
	}

	if id.CodeName, ok = input.Parsed["codeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "codeName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateRegistryCodeVersionID checks that 'input' can be parsed as a Registry Code Version ID
func ValidateRegistryCodeVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegistryCodeVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registry Code Version ID
func (id RegistryCodeVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/registries/%s/codes/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RegistryName, id.CodeName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Registry Code Version ID
func (id RegistryCodeVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticRegistries", "registries", "registries"),
		resourceids.UserSpecifiedSegment("registryName", "registryName"),
		resourceids.StaticSegment("staticCodes", "codes", "codes"),
		resourceids.UserSpecifiedSegment("codeName", "codeName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "version"),
	}
}

// String returns a human-readable description of this Registry Code Version ID
func (id RegistryCodeVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Registry Name: %q", id.RegistryName),
		fmt.Sprintf("Code Name: %q", id.CodeName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Registry Code Version (%s)", strings.Join(components, "\n"))
}
