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
	recaser.RegisterResourceId(&LocationUsageId{})
}

var _ resourceids.ResourceId = &LocationUsageId{}

// LocationUsageId is a struct representing the Resource ID for a Location Usage
type LocationUsageId struct {
	ManagementGroupId    string
	GroupQuotaName       string
	ResourceProviderName string
	LocationUsageName    string
}

// NewLocationUsageID returns a new LocationUsageId struct
func NewLocationUsageID(managementGroupId string, groupQuotaName string, resourceProviderName string, locationUsageName string) LocationUsageId {
	return LocationUsageId{
		ManagementGroupId:    managementGroupId,
		GroupQuotaName:       groupQuotaName,
		ResourceProviderName: resourceProviderName,
		LocationUsageName:    locationUsageName,
	}
}

// ParseLocationUsageID parses 'input' into a LocationUsageId
func ParseLocationUsageID(input string) (*LocationUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationUsageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationUsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLocationUsageIDInsensitively parses 'input' case-insensitively into a LocationUsageId
// note: this method should only be used for API response data and not user input
func ParseLocationUsageIDInsensitively(input string) (*LocationUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationUsageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationUsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LocationUsageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LocationUsageName, ok = input.Parsed["locationUsageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationUsageName", input)
	}

	return nil
}

// ValidateLocationUsageID checks that 'input' can be parsed as a Location Usage ID
func ValidateLocationUsageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationUsageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location Usage ID
func (id LocationUsageId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Quota/groupQuotas/%s/resourceProviders/%s/locationUsages/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.GroupQuotaName, id.ResourceProviderName, id.LocationUsageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Location Usage ID
func (id LocationUsageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticLocationUsages", "locationUsages", "locationUsages"),
		resourceids.UserSpecifiedSegment("locationUsageName", "locationUsageName"),
	}
}

// String returns a human-readable description of this Location Usage ID
func (id LocationUsageId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Resource Provider Name: %q", id.ResourceProviderName),
		fmt.Sprintf("Location Usage Name: %q", id.LocationUsageName),
	}
	return fmt.Sprintf("Location Usage (%s)", strings.Join(components, "\n"))
}
