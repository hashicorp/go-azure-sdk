package extendedueinformation

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExtendedUeInfoProperties = UeInfo5G{}

type UeInfo5G struct {
	Info UeInfo5GProperties `json:"info"`

	// Fields inherited from ExtendedUeInfoProperties
	LastReadAt *string `json:"lastReadAt,omitempty"`
}

func (o *UeInfo5G) GetLastReadAtAsTime() (*time.Time, error) {
	if o.LastReadAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastReadAt, "2006-01-02T15:04:05Z07:00")
}

func (o *UeInfo5G) SetLastReadAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastReadAt = &formatted
}

var _ json.Marshaler = UeInfo5G{}

func (s UeInfo5G) MarshalJSON() ([]byte, error) {
	type wrapper UeInfo5G
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UeInfo5G: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UeInfo5G: %+v", err)
	}
	decoded["ratType"] = "5G"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UeInfo5G: %+v", err)
	}

	return encoded, nil
}
