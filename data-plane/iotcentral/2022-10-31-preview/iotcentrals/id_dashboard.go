package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DashboardId{}

// DashboardId is a struct representing the Resource ID for a Dashboard
type DashboardId struct {
	BaseURI     string
	DashboardId string
}

// NewDashboardID returns a new DashboardId struct
func NewDashboardID(baseURI string, dashboardId string) DashboardId {
	return DashboardId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		DashboardId: dashboardId,
	}
}

// ParseDashboardID parses 'input' into a DashboardId
func ParseDashboardID(input string) (*DashboardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DashboardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DashboardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDashboardIDInsensitively parses 'input' case-insensitively into a DashboardId
// note: this method should only be used for API response data and not user input
func ParseDashboardIDInsensitively(input string) (*DashboardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DashboardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DashboardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DashboardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DashboardId, ok = input.Parsed["dashboardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dashboardId", input)
	}

	return nil
}

// ValidateDashboardID checks that 'input' can be parsed as a Dashboard ID
func ValidateDashboardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDashboardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dashboard ID
func (id DashboardId) ID() string {
	fmtString := "%s/dashboards/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DashboardId)
}

// Path returns the formatted Dashboard ID without the BaseURI
func (id DashboardId) Path() string {
	fmtString := "/dashboards/%s"
	return fmt.Sprintf(fmtString, id.DashboardId)
}

// PathElements returns the values of Dashboard ID Segments without the BaseURI
func (id DashboardId) PathElements() []any {
	return []any{id.DashboardId}
}

// Segments returns a slice of Resource ID Segments which comprise this Dashboard ID
func (id DashboardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDashboards", "dashboards", "dashboards"),
		resourceids.UserSpecifiedSegment("dashboardId", "dashboardId"),
	}
}

// String returns a human-readable description of this Dashboard ID
func (id DashboardId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Dashboard: %q", id.DashboardId),
	}
	return fmt.Sprintf("Dashboard (%s)", strings.Join(components, "\n"))
}
