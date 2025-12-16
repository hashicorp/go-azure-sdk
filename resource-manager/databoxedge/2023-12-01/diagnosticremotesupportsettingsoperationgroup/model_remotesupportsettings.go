package diagnosticremotesupportsettingsoperationgroup

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteSupportSettings struct {
	AccessLevel              *AccessLevel           `json:"accessLevel,omitempty"`
	ExpirationTimeStampInUTC *string                `json:"expirationTimeStampInUTC,omitempty"`
	RemoteApplicationType    *RemoteApplicationType `json:"remoteApplicationType,omitempty"`
}

func (o *RemoteSupportSettings) GetExpirationTimeStampInUTCAsTime() (*time.Time, error) {
	if o.ExpirationTimeStampInUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTimeStampInUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *RemoteSupportSettings) SetExpirationTimeStampInUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTimeStampInUTC = &formatted
}
