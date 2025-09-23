package groupquotasentities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ResourceProviderId{})
}

var _ resourceids.ResourceId = &ResourceProviderId{}

// ResourceProviderId is a struct representing the Resource ID for a Resource Provider
type ResourceProviderId struct {
	ManagementGroupId    string
	GroupQuotaName       string
	ResourceProviderName string
}

// NewResourceProviderID returns a new ResourceProviderId struct
func NewResourceProviderID(managementGroupId string, groupQuotaName string, resourceProviderName string) ResourceProviderId {
	return ResourceProviderId{
		ManagementGroupId:    managementGroupId,
		GroupQuotaName:       groupQuotaName,
		ResourceProviderName: resourceProviderName,
	}
}

// ParseResourceProviderID parses 'input' into a ResourceProviderId
func ParseResourceProviderID(input string) (*ResourceProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResourceProviderIDInsensitively parses 'input' case-insensitively into a ResourceProviderId
// note: this method should only be used for API response data and not user input
func ParseResourceProviderIDInsensitively(input string) (*ResourceProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResourceProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupId, ok = input.Parsed["managementGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", input)
	}

	if id.GroupQuotaName, ok = input.Parsed["groupQuotaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupQuotaName", input)
	}

	if id.ResourceProviderName, ok = input.Parsed["resourceProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceProviderName", input)
	}

	return nil
}

// ValidateResourceProviderID checks that 'input' can be parsed as a Resource Provider ID
func ValidateResourceProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourceProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Provider ID
func (id ResourceProviderId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Quota/groupQuotas/%s/resourceProviders/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.GroupQuotaName, id.ResourceProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Provider ID
func (id ResourceProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupId"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticGroupQuotas", "groupQuotas", "groupQuotas"),
		resourceids.UserSpecifiedSegment("groupQuotaName", "groupQuotaName"),
		resourceids.StaticSegment("staticResourceProviders", "resourceProviders", "resourceProviders"),
		resourceids.UserSpecifiedSegment("resourceProviderName", "resourceProviderName"),
	}
}

// String returns a human-readable description of this Resource Provider ID
func (id ResourceProviderId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Resource Provider Name: %q", id.ResourceProviderName),
	}
	return fmt.Sprintf("Resource Provider (%s)", strings.Join(components, "\n"))
}
