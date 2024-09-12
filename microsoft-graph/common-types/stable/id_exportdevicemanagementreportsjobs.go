package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportDeviceManagementReportsJobsId{}

// ExportDeviceManagementReportsJobsId is a struct representing the Resource ID for a Export Device Management Reports Jobs
type ExportDeviceManagementReportsJobsId struct {
	DeviceManagementExportJobId string
}

// NewExportDeviceManagementReportsJobsID returns a new ExportDeviceManagementReportsJobsId struct
func NewExportDeviceManagementReportsJobsID(deviceManagementExportJobId string) ExportDeviceManagementReportsJobsId {
	return ExportDeviceManagementReportsJobsId{
		DeviceManagementExportJobId: deviceManagementExportJobId,
	}
}

// ParseExportDeviceManagementReportsJobsID parses 'input' into a ExportDeviceManagementReportsJobsId
func ParseExportDeviceManagementReportsJobsID(input string) (*ExportDeviceManagementReportsJobsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportDeviceManagementReportsJobsId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportDeviceManagementReportsJobsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExportDeviceManagementReportsJobsIDInsensitively parses 'input' case-insensitively into a ExportDeviceManagementReportsJobsId
// note: this method should only be used for API response data and not user input
func ParseExportDeviceManagementReportsJobsIDInsensitively(input string) (*ExportDeviceManagementReportsJobsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportDeviceManagementReportsJobsId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportDeviceManagementReportsJobsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExportDeviceManagementReportsJobsId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementExportJobId, ok = input.Parsed["deviceManagementExportJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementExportJobId", input)
	}

	return nil
}

// ValidateExportDeviceManagementReportsJobsID checks that 'input' can be parsed as a Export Device Management Reports Jobs ID
func ValidateExportDeviceManagementReportsJobsID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExportDeviceManagementReportsJobsID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Export Device Management Reports Jobs ID
func (id ExportDeviceManagementReportsJobsId) ID() string {
	fmtString := "/deviceManagement/reports/exportJobs/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementExportJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Export Device Management Reports Jobs ID
func (id ExportDeviceManagementReportsJobsId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("exportJobs", "exportJobs", "exportJobs"),
		resourceids.UserSpecifiedSegment("deviceManagementExportJobId", "deviceManagementExportJobIdValue"),
	}
}

// String returns a human-readable description of this Export Device Management Reports Jobs ID
func (id ExportDeviceManagementReportsJobsId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Export Job: %q", id.DeviceManagementExportJobId),
	}
	return fmt.Sprintf("Export Device Management Reports Jobs (%s)", strings.Join(components, "\n"))
}
