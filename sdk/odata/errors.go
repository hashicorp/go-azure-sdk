// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MPL-2.0 License. See NOTICE.txt in the project root for license information.

package odata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	ErrorAddedObjectReferencesAlreadyExist      = "One or more added object references already exist"
	ErrorCannotDeleteOrUpdateEnabledEntitlement = "Permission (scope or role) cannot be deleted or updated unless disabled first"
	ErrorConflictingObjectPresentInDirectory    = "A conflicting object with one or more of the specified property values is present in the directory"
	ErrorNotValidReferenceUpdate                = "Not a valid reference update"
	ErrorPropertyValuesAreInvalid               = "One or more property values specified are invalid"
	ErrorResourceDoesNotExist                   = "Resource '.+' does not exist or one of its queried reference-property objects are not present"
	ErrorRemovedObjectReferencesDoNotExist      = "One or more removed object references do not exist"
	ErrorServicePrincipalAppInOtherTenant       = "When using this permission, the backing application of the service principal being created must in the local tenant"
	ErrorServicePrincipalInvalidAppId           = "The appId '.+' of the service principal does not reference a valid application object"
	ErrorUnknownUnsupportedQuery                = "UnknownError: Unsupported Query"
)

// Error is used to unmarshal an API error message.
type Error struct {
	Code            *string          `json:"code"`
	Date            *string          `json:"date"`
	Message         *string          `json:"-"`
	RawMessage      *json.RawMessage `json:"message"` // sometimes a string, sometimes an object :/
	ClientRequestId *string          `json:"client-request-id"`
	RequestId       *string          `json:"request-id"`

	InnerError *Error `json:"innerError"` // nested errors

	Details    *[]ErrorDetails  `json:"-"`
	RawDetails *json.RawMessage `json:"details"` // should be an array, but sometimes an object :S

	Values *[]struct {
		Item  string `json:"item"`
		Value string `json:"value"`
	} `json:"values"`

	// Raw is the error exactly as the API sent it, which FullString uses to show the parts that are not modelled above.
	Raw json.RawMessage `json:"-"`
}

type ErrorDetails struct {
	Code   *string `json:"code"`
	Target *string `json:"target"`
}

func (e *Error) UnmarshalJSON(data []byte) error {
	// Perform unmarshalling using a local type
	type error Error
	var e2 error
	if err := json.Unmarshal(data, &e2); err != nil {
		return err
	}
	*e = Error(e2)
	e.Raw = append(json.RawMessage{}, data...)

	// Unmarshal the message, which can be a plain string or an object wrapping a message
	if raw := e.RawMessage; raw != nil && len(*raw) > 0 {
		switch string((*raw)[0]) {
		case "\"":
			var s string
			if err := json.Unmarshal(*raw, &s); err != nil {
				return err
			}
			e.Message = &s
		case "{":
			var m map[string]interface{}
			if err := json.Unmarshal(*raw, &m); err != nil {
				return err
			}
			if v, ok := m["value"]; ok {
				if s, ok := v.(string); ok {
					e.Message = &s
				}
			}
		default:
			return fmt.Errorf("unrecognised error message: %#v", string(*raw))
		}
	}

	// Unmarshal the details, which should be an array but might be an object
	if raw := e.RawDetails; raw != nil && len(*raw) > 0 {
		var details ErrorDetails
		err := json.Unmarshal(*raw, &details)
		if err == nil {
			e.Details = &[]ErrorDetails{
				details,
			}
			return nil
		}

		var detailsSlice []ErrorDetails
		if errSlice := json.Unmarshal(*raw, &detailsSlice); errSlice != nil {
			return fmt.Errorf("failed to unmarshal `RawDetails` as either `ErrorDetails` (%+v) or `[]ErrorDetails` (%+v) - type: %+v", err, errSlice, *raw)
		}
		e.Details = &detailsSlice
	}

	return nil
}

func (e Error) String() string {
	sl := make([]string, 0)
	if e.Code != nil {
		sl = append(sl, *e.Code)
	}
	if e.Message != nil {
		sl = append(sl, *e.Message)
	}
	if e.InnerError != nil {
		if is := e.InnerError.String(); is != "" {
			sl = append(sl, is)
		}
	}
	return strings.Join(sl, ": ")
}

// FullString is String followed by everything else the API sent with the error, one item per line: the target, nested details, the rest of the inner
// error and any additional info. APIs often put the real reason there and leave the top message generic.
func (e Error) FullString() string {
	if len(e.Raw) == 0 {
		return e.String()
	}

	decoder := json.NewDecoder(bytes.NewReader(e.Raw))
	decoder.UseNumber()

	var body map[string]interface{}
	if err := decoder.Decode(&body); err != nil {
		return e.String()
	}

	lines := describeError(body, "", true)
	if len(lines) == 0 {
		return e.String()
	}

	return e.String() + "\n" + strings.Join(lines, "\n")
}

// describeError is one error object from an API response as lines indented by pad, descending into anything nested. headPrinted is set for the error
// and the inner errors that String has already covered, so their code and message are not repeated.
func describeError(obj map[string]interface{}, pad string, headPrinted bool) []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	inside := pad + "  "
	target := ""
	fields := make([]string, 0)
	details := make([]string, 0)
	inner := make([]string, 0)
	info := make([]string, 0)

	for _, k := range keys {
		v := obj[k]

		switch strings.ToLower(k) {
		case "code", "message":
			// part of the head, below

		case "target":
			target = errorText(v)

		case "details":
			details = describeNested(k, v, inside, false)

		case "innererror":
			inner = describeNested(k, v, inside, headPrinted)

		case "additionalinfo":
			info = describeAdditionalInfo(k, v, inside)

		default:
			fields = append(fields, describeField(k, v, inside)...)
		}
	}

	lines := make([]string, 0)

	head := errorHead(obj)
	switch {
	case headPrinted || head == "":
		lines = append(lines, describeField("target", target, inside)...)

	case target != "":
		lines = append(lines, fmt.Sprintf("%s%s (target: %s)", pad, indentLines(head, inside), target))

	default:
		lines = append(lines, pad+indentLines(head, inside))
	}

	lines = append(lines, fields...)
	lines = append(lines, details...)
	lines = append(lines, inner...)

	return append(lines, info...)
}

// describeNested is the lines for a details list or an inner error, which may be one object, a list of them or plain text. An object with no code or
// message of its own to lead with is listed under its key instead.
func describeNested(key string, v interface{}, pad string, headPrinted bool) []string {
	switch val := v.(type) {
	case map[string]interface{}:
		lines := describeError(val, pad, headPrinted)
		if len(lines) > 0 && (headPrinted || errorHead(val) == "") {
			lines = append([]string{fmt.Sprintf("%s%s:", pad, key)}, lines...)
		}

		return lines

	case []interface{}:
		lines := make([]string, 0)
		for _, item := range val {
			lines = append(lines, describeNested(key, item, pad, false)...)
		}

		return lines
	}

	return describeField(key, v, pad)
}

// describeAdditionalInfo is one line per item of an additionalInfo list, named after the item's type, e.g. PolicyViolation.
func describeAdditionalInfo(key string, v interface{}, pad string) []string {
	items, ok := v.([]interface{})
	if !ok {
		return describeField(key, v, pad)
	}

	lines := make([]string, 0)
	for _, item := range items {
		info, _ := item.(map[string]interface{})
		if label := errorText(info["type"]); label != "" {
			lines = append(lines, describeField(label, info["info"], pad)...)
			continue
		}

		lines = append(lines, describeField(key, item, pad)...)
	}

	return lines
}

// describeField is one "name: value" line, or nothing when there is no value. Anything but text, a number or a boolean is shown as compact JSON.
func describeField(key string, v interface{}, pad string) []string {
	switch val := v.(type) {
	case nil:
		return nil

	case string:
		if val == "" {
			return nil
		}

		return []string{fmt.Sprintf("%s%s: %s", pad, key, indentLines(val, pad+"  "))}

	case json.Number, bool:
		return []string{fmt.Sprintf("%s%s: %v", pad, key, val)}
	}

	b, err := json.Marshal(v)
	if err != nil || string(b) == "{}" || string(b) == "[]" {
		return nil
	}

	return []string{fmt.Sprintf("%s%s: %s", pad, key, b)}
}

// errorHead is the "code: message" an error object leads with, or whichever of the two it has.
func errorHead(obj map[string]interface{}) string {
	head := make([]string, 0, 2)

	for _, name := range []string{"code", "message"} {
		for k, v := range obj {
			if !strings.EqualFold(k, name) {
				continue
			}

			// a message is sometimes an object wrapping the text
			if m, ok := v.(map[string]interface{}); ok {
				v = m["value"]
			}

			if s := errorText(v); s != "" {
				head = append(head, s)
			}
		}
	}

	return strings.Join(head, ": ")
}

// errorText is v as text when it is text or a number, and empty otherwise.
func errorText(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val

	case json.Number:
		return val.String()
	}

	return ""
}

// indentLines keeps a value that runs over several lines under the line it starts on.
func indentLines(s, pad string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(s), "\r\n", "\n"), "\n", "\n"+pad)
}

func (e Error) Match(errorText string) bool {
	re := regexp.MustCompile(errorText)
	return re.MatchString(e.String())
}
