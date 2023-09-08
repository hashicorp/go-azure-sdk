package meinformationprotectionsensitivitylabelsublabel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeInformationProtectionSensitivityLabelSublabelId{}

// MeInformationProtectionSensitivityLabelSublabelId is a struct representing the Resource ID for a Me Information Protection Sensitivity Label Sublabel
type MeInformationProtectionSensitivityLabelSublabelId struct {
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewMeInformationProtectionSensitivityLabelSublabelID returns a new MeInformationProtectionSensitivityLabelSublabelId struct
func NewMeInformationProtectionSensitivityLabelSublabelID(sensitivityLabelId string, sensitivityLabelId1 string) MeInformationProtectionSensitivityLabelSublabelId {
	return MeInformationProtectionSensitivityLabelSublabelId{
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseMeInformationProtectionSensitivityLabelSublabelID parses 'input' into a MeInformationProtectionSensitivityLabelSublabelId
func ParseMeInformationProtectionSensitivityLabelSublabelID(input string) (*MeInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeInformationProtectionSensitivityLabelSublabelId{}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ParseMeInformationProtectionSensitivityLabelSublabelIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionSensitivityLabelSublabelId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionSensitivityLabelSublabelIDInsensitively(input string) (*MeInformationProtectionSensitivityLabelSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeInformationProtectionSensitivityLabelSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeInformationProtectionSensitivityLabelSublabelId{}

	if id.SensitivityLabelId, ok = parsed.Parsed["sensitivityLabelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", *parsed)
	}

	if id.SensitivityLabelId1, ok = parsed.Parsed["sensitivityLabelId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", *parsed)
	}

	return &id, nil
}

// ValidateMeInformationProtectionSensitivityLabelSublabelID checks that 'input' can be parsed as a Me Information Protection Sensitivity Label Sublabel ID
func ValidateMeInformationProtectionSensitivityLabelSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionSensitivityLabelSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Sensitivity Label Sublabel ID
func (id MeInformationProtectionSensitivityLabelSublabelId) ID() string {
	fmtString := "/me/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Sensitivity Label Sublabel ID
func (id MeInformationProtectionSensitivityLabelSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticInformationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelIdValue"),
		resourceids.StaticSegment("staticSublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1Value"),
	}
}

// String returns a human-readable description of this Me Information Protection Sensitivity Label Sublabel ID
func (id MeInformationProtectionSensitivityLabelSublabelId) String() string {
	components := []string{
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("Me Information Protection Sensitivity Label Sublabel (%s)", strings.Join(components, "\n"))
}
