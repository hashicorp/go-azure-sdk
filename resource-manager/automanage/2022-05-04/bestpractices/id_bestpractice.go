package bestpractices

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BestPracticeId{})
}

var _ resourceids.ResourceId = &BestPracticeId{}

// BestPracticeId is a struct representing the Resource ID for a Best Practice
type BestPracticeId struct {
	BestPracticeName string
}

// NewBestPracticeID returns a new BestPracticeId struct
func NewBestPracticeID(bestPracticeName string) BestPracticeId {
	return BestPracticeId{
		BestPracticeName: bestPracticeName,
	}
}

// ParseBestPracticeID parses 'input' into a BestPracticeId
func ParseBestPracticeID(input string) (*BestPracticeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BestPracticeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BestPracticeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBestPracticeIDInsensitively parses 'input' case-insensitively into a BestPracticeId
// note: this method should only be used for API response data and not user input
func ParseBestPracticeIDInsensitively(input string) (*BestPracticeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BestPracticeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BestPracticeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BestPracticeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BestPracticeName, ok = input.Parsed["bestPracticeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bestPracticeName", input)
	}

	return nil
}

// ValidateBestPracticeID checks that 'input' can be parsed as a Best Practice ID
func ValidateBestPracticeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBestPracticeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Best Practice ID
func (id BestPracticeId) ID() string {
	fmtString := "/providers/Microsoft.AutoManage/bestPractices/%s"
	return fmt.Sprintf(fmtString, id.BestPracticeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Best Practice ID
func (id BestPracticeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticBestPractices", "bestPractices", "bestPractices"),
		resourceids.UserSpecifiedSegment("bestPracticeName", "bestPracticeValue"),
	}
}

// String returns a human-readable description of this Best Practice ID
func (id BestPracticeId) String() string {
	components := []string{
		fmt.Sprintf("Best Practice Name: %q", id.BestPracticeName),
	}
	return fmt.Sprintf("Best Practice (%s)", strings.Join(components, "\n"))
}
