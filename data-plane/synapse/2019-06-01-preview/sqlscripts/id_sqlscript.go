package sqlscripts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SqlScriptId{}

// SqlScriptId is a struct representing the Resource ID for a Sql Script
type SqlScriptId struct {
	BaseURI       string
	SqlScriptName string
}

// NewSqlScriptID returns a new SqlScriptId struct
func NewSqlScriptID(baseURI string, sqlScriptName string) SqlScriptId {
	return SqlScriptId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		SqlScriptName: sqlScriptName,
	}
}

// ParseSqlScriptID parses 'input' into a SqlScriptId
func ParseSqlScriptID(input string) (*SqlScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlScriptId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSqlScriptIDInsensitively parses 'input' case-insensitively into a SqlScriptId
// note: this method should only be used for API response data and not user input
func ParseSqlScriptIDInsensitively(input string) (*SqlScriptId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlScriptId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlScriptId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SqlScriptId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.SqlScriptName, ok = input.Parsed["sqlScriptName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sqlScriptName", input)
	}

	return nil
}

// ValidateSqlScriptID checks that 'input' can be parsed as a Sql Script ID
func ValidateSqlScriptID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlScriptID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Script ID
func (id SqlScriptId) ID() string {
	fmtString := "%s/sqlScripts/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.SqlScriptName)
}

// Path returns the formatted Sql Script ID without the BaseURI
func (id SqlScriptId) Path() string {
	fmtString := "/sqlScripts/%s"
	return fmt.Sprintf(fmtString, id.SqlScriptName)
}

// PathElements returns the values of Sql Script ID Segments without the BaseURI
func (id SqlScriptId) PathElements() []any {
	return []any{id.SqlScriptName}
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Script ID
func (id SqlScriptId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticSqlScripts", "sqlScripts", "sqlScripts"),
		resourceids.UserSpecifiedSegment("sqlScriptName", "sqlScriptName"),
	}
}

// String returns a human-readable description of this Sql Script ID
func (id SqlScriptId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Sql Script Name: %q", id.SqlScriptName),
	}
	return fmt.Sprintf("Sql Script (%s)", strings.Join(components, "\n"))
}
