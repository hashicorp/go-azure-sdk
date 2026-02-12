package kqlscripts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KqlScriptId{}

// KqlScriptId is a struct representing the Resource ID for a Kql Script
type KqlScriptId struct {
	BaseURI       string
	KqlScriptName string
}

// NewKqlScriptID returns a new KqlScriptId struct
func NewKqlScriptID(baseURI string, kqlScriptName string) KqlScriptId {
	return KqlScriptId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		KqlScriptName: kqlScriptName,
	}
}

// ParseKqlScriptID parses 'input' into a KqlScriptId
func ParseKqlScriptID(input string) (*KqlScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KqlScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KqlScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseKqlScriptIDInsensitively parses 'input' case-insensitively into a KqlScriptId
// note: this method should only be used for API response data and not user input
func ParseKqlScriptIDInsensitively(input string) (*KqlScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KqlScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KqlScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *KqlScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.KqlScriptName, ok = input.Parsed["kqlScriptName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "kqlScriptName", input)
	}

	return nil
}

// ValidateKqlScriptID checks that 'input' can be parsed as a Kql Script ID
func ValidateKqlScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKqlScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Kql Script ID
func (id KqlScriptId) ID() string {
	fmtString := "%s/kqlScripts/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.KqlScriptName)
}

// Path returns the formatted Kql Script ID without the BaseURI
func (id KqlScriptId) Path() string {
	fmtString := "/kqlScripts/%s"
	return fmt.Sprintf(fmtString, id.KqlScriptName)
}

// PathElements returns the values of Kql Script ID Segments without the BaseURI
func (id KqlScriptId) PathElements() []any {
	return []any{id.KqlScriptName}
}

// Segments returns a slice of Resource ID Segments which comprise this Kql Script ID
func (id KqlScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticKqlScripts", "kqlScripts", "kqlScripts"),
		resourceids.UserSpecifiedSegment("kqlScriptName", "kqlScriptName"),
	}
}

// String returns a human-readable description of this Kql Script ID
func (id KqlScriptId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Kql Script Name: %q", id.KqlScriptName),
	}
	return fmt.Sprintf("Kql Script (%s)", strings.Join(components, "\n"))
}
