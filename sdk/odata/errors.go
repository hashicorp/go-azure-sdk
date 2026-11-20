// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MPL-2.0 License. See NOTICE.txt in the project root for license information.

package odata

import (
	"encoding/json"
	"fmt"
	"regexp"
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
}

type ErrorDetails struct {
	Code       *string          `json:"code"`
	Target     *string          `json:"target"`
	Message    *string          `json:"-"`
	RawMessage *json.RawMessage `json:"message"`
	Details    *[]ErrorDetails  `json:"-"`
	RawDetails *json.RawMessage `json:"details"`
}

func unmarshalRawMessage(raw *json.RawMessage) (*string, error) {
	if raw == nil || len(*raw) == 0 || string(*raw) == "null" {
		return nil, nil
	}
	switch string((*raw)[0]) {
	case "\"":
		var s string
		if err := json.Unmarshal(*raw, &s); err != nil {
			return nil, err
		}
		return &s, nil
	case "{":
		var m map[string]interface{}
		if err := json.Unmarshal(*raw, &m); err != nil {
			return nil, err
		}
		if v, ok := m["value"]; ok {
			if s, ok := v.(string); ok {
				return &s, nil
			}
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("unrecognised error message: %#v", string(*raw))
	}
}

func unmarshalRawDetails(raw *json.RawMessage) (*[]ErrorDetails, error) {
	if raw == nil || len(*raw) == 0 || string(*raw) == "null" {
		return nil, nil
	}
	var details ErrorDetails
	err := json.Unmarshal(*raw, &details)
	if err == nil {
		return &[]ErrorDetails{
			details,
		}, nil
	}

	var detailsSlice []ErrorDetails
	if errSlice := json.Unmarshal(*raw, &detailsSlice); errSlice != nil {
		return nil, fmt.Errorf("failed to unmarshal `RawDetails` as either `ErrorDetails` (%+v) or `[]ErrorDetails` (%+v) - type: %+v", err, errSlice, *raw)
	}
	return &detailsSlice, nil
}

func (d *ErrorDetails) UnmarshalJSON(data []byte) error {
	type errorDetails ErrorDetails
	var d2 errorDetails
	if err := json.Unmarshal(data, &d2); err != nil {
		return err
	}
	*d = ErrorDetails(d2)

	msg, err := unmarshalRawMessage(d.RawMessage)
	if err != nil {
		return err
	}
	d.Message = msg

	details, err := unmarshalRawDetails(d.RawDetails)
	if err != nil {
		return err
	}
	d.Details = details

	return nil
}

func (d ErrorDetails) String() string {
	sl := make([]string, 0)
	if d.Code != nil && *d.Code != "" {
		if d.Target != nil && *d.Target != "" {
			sl = append(sl, fmt.Sprintf("%s (%s)", *d.Code, *d.Target))
		} else {
			sl = append(sl, *d.Code)
		}
	} else if d.Target != nil && *d.Target != "" {
		sl = append(sl, *d.Target)
	}
	if d.Message != nil && *d.Message != "" {
		sl = append(sl, *d.Message)
	}
	if d.Details != nil {
		for _, child := range *d.Details {
			if cs := child.String(); cs != "" {
				sl = append(sl, cs)
			}
		}
	}
	return strings.Join(sl, ": ")
}

func (e *Error) UnmarshalJSON(data []byte) error {
	// Perform unmarshalling using a local type
	type error Error
	var e2 error
	if err := json.Unmarshal(data, &e2); err != nil {
		return err
	}
	*e = Error(e2)

	var altIds struct {
		ClientRequestId *string `json:"clientRequestId"`
		RequestId       *string `json:"requestId"`
	}
	if err := json.Unmarshal(data, &altIds); err == nil {
		if e.ClientRequestId == nil && altIds.ClientRequestId != nil {
			e.ClientRequestId = altIds.ClientRequestId
		}
		if e.RequestId == nil && altIds.RequestId != nil {
			e.RequestId = altIds.RequestId
		}
	}

	// Unmarshal the message, which can be a plain string or an object wrapping a message
	msg, err := unmarshalRawMessage(e.RawMessage)
	if err != nil {
		return err
	}
	e.Message = msg

	// Unmarshal the details, which should be an array but might be an object
	details, err := unmarshalRawDetails(e.RawDetails)
	if err != nil {
		return err
	}
	e.Details = details

	return nil
}

func (e Error) String() string {
	sl := make([]string, 0)
	if e.Code != nil && *e.Code != "" {
		sl = append(sl, *e.Code)
	}
	if e.Message != nil && *e.Message != "" {
		sl = append(sl, *e.Message)
	}
	if e.Details != nil {
		for _, d := range *e.Details {
			if ds := d.String(); ds != "" {
				sl = append(sl, ds)
			}
		}
	}
	if e.InnerError != nil {
		if is := e.InnerError.String(); is != "" {
			sl = append(sl, is)
		}
	}
	return strings.Join(sl, ": ")
}

func (e Error) Match(errorText string) bool {
	re := regexp.MustCompile(errorText)
	return re.MatchString(e.String())
}
