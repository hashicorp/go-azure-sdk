package extendedueinformation

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExtendedUeInfoProperties = UeInfo4G{}

type UeInfo4G struct {
	Info UeInfo4GProperties `json:"info"`

	// Fields inherited from ExtendedUeInfoProperties
	LastReadAt *string `json:"lastReadAt,omitempty"`
}

func (o *UeInfo4G) GetLastReadAtAsTime() (*time.Time, error) {
	if o.LastReadAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastReadAt, "2006-01-02T15:04:05Z07:00")
}

func (o *UeInfo4G) SetLastReadAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastReadAt = &formatted
}

var _ json.Marshaler = UeInfo4G{}

func (s UeInfo4G) MarshalJSON() ([]byte, error) {
	type wrapper UeInfo4G
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UeInfo4G: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UeInfo4G: %+v", err)
	}
	decoded["ratType"] = "4G"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UeInfo4G: %+v", err)
	}

	return encoded, nil
}
