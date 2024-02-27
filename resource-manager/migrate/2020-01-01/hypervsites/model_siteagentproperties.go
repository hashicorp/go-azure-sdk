package hypervsites

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAgentProperties struct {
	Id               *string `json:"id,omitempty"`
	KeyVaultId       *string `json:"keyVaultId,omitempty"`
	KeyVaultUri      *string `json:"keyVaultUri,omitempty"`
	LastHeartBeatUtc *string `json:"lastHeartBeatUtc,omitempty"`
	Version          *string `json:"version,omitempty"`
}

func (o *SiteAgentProperties) GetLastHeartBeatUtcAsTime() (*time.Time, error) {
	if o.LastHeartBeatUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastHeartBeatUtc, "2006-01-02T15:04:05Z07:00")
}
