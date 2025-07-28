package groupquotalocationsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LocationSettingId{})
}

var _ resourceids.ResourceId = &LocationSettingId{}

// LocationSettingId is a struct representing the Resource ID for a Location Setting
type LocationSettingId struct {
	ManagementGroupId    string
	GroupQuotaName       string
	ResourceProviderName string
	LocationSettingName  string
}

// NewLocationSettingID returns a new LocationSettingId struct
func NewLocationSettingID(managementGroupId string, groupQuotaName string, resourceProviderName string, locationSettingName string) LocationSettingId {
	return LocationSettingId{
		ManagementGroupId:    managementGroupId,
		GroupQuotaName:       groupQuotaName,
		ResourceProviderName: resourceProviderName,
		LocationSettingName:  locationSettingName,
	}
}

// ParseLocationSettingID parses 'input' into a LocationSettingId
func ParseLocationSettingID(input string) (*LocationSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLocationSettingIDInsensitively parses 'input' case-insensitively into a LocationSettingId
// note: this method should only be used for API response data and not user input
func ParseLocationSettingIDInsensitively(input string) (*LocationSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LocationSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LocationSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LocationSettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LocationSettingName, ok = input.Parsed["locationSettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationSettingName", input)
	}

	return nil
}

// ValidateLocationSettingID checks that 'input' can be parsed as a Location Setting ID
func ValidateLocationSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location Setting ID
func (id LocationSettingId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Quota/groupQuotas/%s/resourceProviders/%s/locationSettings/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.GroupQuotaName, id.ResourceProviderName, id.LocationSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Location Setting ID
func (id LocationSettingId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticLocationSettings", "locationSettings", "locationSettings"),
		resourceids.UserSpecifiedSegment("locationSettingName", "locationSettingName"),
	}
}

// String returns a human-readable description of this Location Setting ID
func (id LocationSettingId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Group Quota Name: %q", id.GroupQuotaName),
		fmt.Sprintf("Resource Provider Name: %q", id.ResourceProviderName),
		fmt.Sprintf("Location Setting Name: %q", id.LocationSettingName),
	}
	return fmt.Sprintf("Location Setting (%s)", strings.Join(components, "\n"))
}
