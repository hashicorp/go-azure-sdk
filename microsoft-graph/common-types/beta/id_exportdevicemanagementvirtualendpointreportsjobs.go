package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportDeviceManagementVirtualEndpointReportsJobsId{}

// ExportDeviceManagementVirtualEndpointReportsJobsId is a struct representing the Resource ID for a Export Device Management Virtual Endpoint Reports Jobs
type ExportDeviceManagementVirtualEndpointReportsJobsId struct {
	CloudPCExportJobId string
}

// NewExportDeviceManagementVirtualEndpointReportsJobsID returns a new ExportDeviceManagementVirtualEndpointReportsJobsId struct
func NewExportDeviceManagementVirtualEndpointReportsJobsID(cloudPCExportJobId string) ExportDeviceManagementVirtualEndpointReportsJobsId {
	return ExportDeviceManagementVirtualEndpointReportsJobsId{
		CloudPCExportJobId: cloudPCExportJobId,
	}
}

// ParseExportDeviceManagementVirtualEndpointReportsJobsID parses 'input' into a ExportDeviceManagementVirtualEndpointReportsJobsId
func ParseExportDeviceManagementVirtualEndpointReportsJobsID(input string) (*ExportDeviceManagementVirtualEndpointReportsJobsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportDeviceManagementVirtualEndpointReportsJobsId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportDeviceManagementVirtualEndpointReportsJobsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExportDeviceManagementVirtualEndpointReportsJobsIDInsensitively parses 'input' case-insensitively into a ExportDeviceManagementVirtualEndpointReportsJobsId
// note: this method should only be used for API response data and not user input
func ParseExportDeviceManagementVirtualEndpointReportsJobsIDInsensitively(input string) (*ExportDeviceManagementVirtualEndpointReportsJobsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportDeviceManagementVirtualEndpointReportsJobsId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportDeviceManagementVirtualEndpointReportsJobsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExportDeviceManagementVirtualEndpointReportsJobsId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCExportJobId, ok = input.Parsed["cloudPCExportJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCExportJobId", input)
	}

	return nil
}

// ValidateExportDeviceManagementVirtualEndpointReportsJobsID checks that 'input' can be parsed as a Export Device Management Virtual Endpoint Reports Jobs ID
func ValidateExportDeviceManagementVirtualEndpointReportsJobsID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExportDeviceManagementVirtualEndpointReportsJobsID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Export Device Management Virtual Endpoint Reports Jobs ID
func (id ExportDeviceManagementVirtualEndpointReportsJobsId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/reports/exportJobs/%s"
	return fmt.Sprintf(fmtString, id.CloudPCExportJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Export Device Management Virtual Endpoint Reports Jobs ID
func (id ExportDeviceManagementVirtualEndpointReportsJobsId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("reports", "reports", "reports"),
		resourceids.StaticSegment("exportJobs", "exportJobs", "exportJobs"),
		resourceids.UserSpecifiedSegment("cloudPCExportJobId", "cloudPCExportJobIdValue"),
	}
}

// String returns a human-readable description of this Export Device Management Virtual Endpoint Reports Jobs ID
func (id ExportDeviceManagementVirtualEndpointReportsJobsId) String() string {
	components := []string{
		fmt.Sprintf("Cloud P C Export Job: %q", id.CloudPCExportJobId),
	}
	return fmt.Sprintf("Export Device Management Virtual Endpoint Reports Jobs (%s)", strings.Join(components, "\n"))
}
