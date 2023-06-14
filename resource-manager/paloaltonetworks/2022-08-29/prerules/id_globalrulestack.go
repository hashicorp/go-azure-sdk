package prerules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GlobalRuleStackId{}

// GlobalRuleStackId is a struct representing the Resource ID for a Global Rule Stack
type GlobalRuleStackId struct {
	GlobalRuleStackName string
}

// NewGlobalRuleStackID returns a new GlobalRuleStackId struct
func NewGlobalRuleStackID(globalRuleStackName string) GlobalRuleStackId {
	return GlobalRuleStackId{
		GlobalRuleStackName: globalRuleStackName,
	}
}

// ParseGlobalRuleStackID parses 'input' into a GlobalRuleStackId
func ParseGlobalRuleStackID(input string) (*GlobalRuleStackId, error) {
	parser := resourceids.NewParserFromResourceIdType(GlobalRuleStackId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GlobalRuleStackId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	return &id, nil
}

// ParseGlobalRuleStackIDInsensitively parses 'input' case-insensitively into a GlobalRuleStackId
// note: this method should only be used for API response data and not user input
func ParseGlobalRuleStackIDInsensitively(input string) (*GlobalRuleStackId, error) {
	parser := resourceids.NewParserFromResourceIdType(GlobalRuleStackId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GlobalRuleStackId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	return &id, nil
}

// ValidateGlobalRuleStackID checks that 'input' can be parsed as a Global Rule Stack ID
func ValidateGlobalRuleStackID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGlobalRuleStackID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Global Rule Stack ID
func (id GlobalRuleStackId) ID() string {
	fmtString := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/%s"
	return fmt.Sprintf(fmtString, id.GlobalRuleStackName)
}

// Segments returns a slice of Resource ID Segments which comprise this Global Rule Stack ID
func (id GlobalRuleStackId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticPaloAltoNetworksCloudNGFW", "PaloAltoNetworks.CloudNGFW", "PaloAltoNetworks.CloudNGFW"),
		resourceids.StaticSegment("staticGlobalRuleStacks", "globalRuleStacks", "globalRuleStacks"),
		resourceids.UserSpecifiedSegment("globalRuleStackName", "globalRuleStackValue"),
	}
}

// String returns a human-readable description of this Global Rule Stack ID
func (id GlobalRuleStackId) String() string {
	components := []string{
		fmt.Sprintf("Global Rule Stack Name: %q", id.GlobalRuleStackName),
	}
	return fmt.Sprintf("Global Rule Stack (%s)", strings.Join(components, "\n"))
}
