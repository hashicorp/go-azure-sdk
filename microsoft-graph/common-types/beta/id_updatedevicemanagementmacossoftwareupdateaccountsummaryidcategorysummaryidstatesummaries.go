package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{}

// UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId is a struct representing the Resource ID for a Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries
type UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId struct {
	MacOSSoftwareUpdateAccountSummaryId  string
	MacOSSoftwareUpdateCategorySummaryId string
	MacOSSoftwareUpdateStateSummaryId    string
}

// NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID returns a new UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId struct
func NewUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(macOSSoftwareUpdateAccountSummaryId string, macOSSoftwareUpdateCategorySummaryId string, macOSSoftwareUpdateStateSummaryId string) UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId {
	return UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{
		MacOSSoftwareUpdateAccountSummaryId:  macOSSoftwareUpdateAccountSummaryId,
		MacOSSoftwareUpdateCategorySummaryId: macOSSoftwareUpdateCategorySummaryId,
		MacOSSoftwareUpdateStateSummaryId:    macOSSoftwareUpdateStateSummaryId,
	}
}

// ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID parses 'input' into a UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId
func ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(input string) (*UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesIDInsensitively parses 'input' case-insensitively into a UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId
// note: this method should only be used for API response data and not user input
func ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesIDInsensitively(input string) (*UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MacOSSoftwareUpdateAccountSummaryId, ok = input.Parsed["macOSSoftwareUpdateAccountSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateAccountSummaryId", input)
	}

	if id.MacOSSoftwareUpdateCategorySummaryId, ok = input.Parsed["macOSSoftwareUpdateCategorySummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateCategorySummaryId", input)
	}

	if id.MacOSSoftwareUpdateStateSummaryId, ok = input.Parsed["macOSSoftwareUpdateStateSummaryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "macOSSoftwareUpdateStateSummaryId", input)
	}

	return nil
}

// ValidateUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID checks that 'input' can be parsed as a Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries ID
func ValidateUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries ID
func (id UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId) ID() string {
	fmtString := "/deviceManagement/macOSSoftwareUpdateAccountSummaries/%s/categorySummaries/%s/updateStateSummaries/%s"
	return fmt.Sprintf(fmtString, id.MacOSSoftwareUpdateAccountSummaryId, id.MacOSSoftwareUpdateCategorySummaryId, id.MacOSSoftwareUpdateStateSummaryId)
}

// Segments returns a slice of Resource ID Segments which comprise this Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries ID
func (id UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries", "macOSSoftwareUpdateAccountSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateAccountSummaryId", "macOSSoftwareUpdateAccountSummaryIdValue"),
		resourceids.StaticSegment("categorySummaries", "categorySummaries", "categorySummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateCategorySummaryId", "macOSSoftwareUpdateCategorySummaryIdValue"),
		resourceids.StaticSegment("updateStateSummaries", "updateStateSummaries", "updateStateSummaries"),
		resourceids.UserSpecifiedSegment("macOSSoftwareUpdateStateSummaryId", "macOSSoftwareUpdateStateSummaryIdValue"),
	}
}

// String returns a human-readable description of this Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries ID
func (id UpdateDeviceManagementMacOSSoftwareUpdateAccountSummaryIdCategorySummaryIdStateSummariesId) String() string {
	components := []string{
		fmt.Sprintf("Mac O S Software Update Account Summary: %q", id.MacOSSoftwareUpdateAccountSummaryId),
		fmt.Sprintf("Mac O S Software Update Category Summary: %q", id.MacOSSoftwareUpdateCategorySummaryId),
		fmt.Sprintf("Mac O S Software Update State Summary: %q", id.MacOSSoftwareUpdateStateSummaryId),
	}
	return fmt.Sprintf("Update Device Management Mac O S Software Update Account Summary Id Category Summary Id State Summaries (%s)", strings.Join(components, "\n"))
}
