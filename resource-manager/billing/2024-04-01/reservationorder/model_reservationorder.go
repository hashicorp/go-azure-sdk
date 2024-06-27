package reservationorder

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrder struct {
	Etag       *int64                    `json:"etag,omitempty"`
	Id         *string                   `json:"id,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Properties *ReservationOrderProperty `json:"properties,omitempty"`
	SystemData *systemdata.SystemData    `json:"systemData,omitempty"`
	Tags       *map[string]string        `json:"tags,omitempty"`
	Type       *string                   `json:"type,omitempty"`
}
