package allowedconnections

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedConnectionsResourceProperties struct {
	CalculatedDateTime   *string                `json:"calculatedDateTime,omitempty"`
	ConnectableResources *[]ConnectableResource `json:"connectableResources,omitempty"`
}

func (o *AllowedConnectionsResourceProperties) GetCalculatedDateTimeAsTime() (*time.Time, error) {
	if o.CalculatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CalculatedDateTime, "2006-01-02T15:04:05Z07:00")
}
