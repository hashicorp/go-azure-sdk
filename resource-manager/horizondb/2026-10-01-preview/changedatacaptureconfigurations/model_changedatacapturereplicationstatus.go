package changedatacaptureconfigurations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureReplicationStatus struct {
	ExternalArtifactId *string                           `json:"externalArtifactId,omitempty"`
	LastConnectedAtUtc *string                           `json:"lastConnectedAtUtc,omitempty"`
	State              ChangeDataCaptureReplicationState `json:"state"`
}

func (o *ChangeDataCaptureReplicationStatus) GetLastConnectedAtUtcAsTime() (*time.Time, error) {
	if o.LastConnectedAtUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastConnectedAtUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ChangeDataCaptureReplicationStatus) SetLastConnectedAtUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastConnectedAtUtc = &formatted
}
