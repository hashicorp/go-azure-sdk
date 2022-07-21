package odata

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"net/url"
	"regexp"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const ODataVersion = "4.0" // TODO: support 4.01 - https://docs.oasis-open.org/odata/odata-json-format/v4.01/cs01/odata-json-format-v4.01-cs01.html#_Toc499720587

type Id string

func (o Id) MarshalJSON() ([]byte, error) {
	id := regexp.MustCompile(`/v2/`).ReplaceAllString(string(o), `/v1.0/`)

	u, err := url.Parse(id)
	if err != nil || u.Scheme == "" || u.Host == "" {
		matches := regexp.MustCompile(`([^()'"]+)\(['"]([^'"]+)['"]\)`).FindStringSubmatch(id)
		if len(matches) != 3 {
			return nil, fmt.Errorf("Marshaling odata.Id: could not match a GUID")
		}

		objectType := matches[1]
		guid := matches[2]
		if _, err = uuid.ParseUUID(guid); err != nil {
			return nil, fmt.Errorf("Marshaling odata.Id: %+v", err)
		}

		// Although we're hard-coding `graph.microsoft.com` here, this doesn't _appear_ to be a problem
		// The host can seemingly be anything, even complete nonsense, and the API will accept it provided
		// it can parse out a version number, an object type and a GUID.
		id = fmt.Sprintf("https://graph.microsoft.com/v1.0/%s/%s", objectType, guid)
	}

	return json.Marshal(id)
}

func (o *Id) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	*o = Id(regexp.MustCompile(`/v2/`).ReplaceAllString(id, `/v1.0/`))

	return nil
}

type Link string

func (o *Link) UnmarshalJSON(data []byte) error {
	var link string
	if err := json.Unmarshal(data, &link); err != nil {
		return err
	}
	*o = Link(regexp.MustCompile(`/v2/`).ReplaceAllString(link, `/v1.0/`))
	return nil
}

// OData is used to unmarshall OData metadata from an API response.
type OData struct {
	Context      *string `json:"@odata.context"`
	MetadataEtag *string `json:"@odata.metadataEtag"`
	Type         *Type   `json:"@odata.type"`
	Count        *int    `json:"@odata.count"`
	NextLink     *string `json:"@odata.nextLink"`
	Delta        *string `json:"@odata.delta"`
	DeltaLink    *string `json:"@odata.deltaLink"`
	Id           *Id     `json:"@odata.id"`
	EditLink     *Link   `json:"@odata.editLink"`
	Etag         *string `json:"@odata.etag"`

	Error *Error `json:"-"`

	Value interface{} `json:"value"`
}

func (o *OData) UnmarshalJSON(data []byte) error {
	// Perform unmarshalling using a local type
	type odata OData
	var o2 odata
	if err := json.Unmarshal(data, &o2); err != nil {
		return err
	}
	*o = OData(o2)

	// Look for errors in the "error" and "odata.error" fields
	var e map[string]json.RawMessage
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	for _, k := range []string{"error", "odata.error"} {
		if v, ok := e[k]; ok {
			var e2 Error
			if err := json.Unmarshal(v, &e2); err != nil {
				return err
			}
			o.Error = &e2
			break
		}
	}
	return nil
}
