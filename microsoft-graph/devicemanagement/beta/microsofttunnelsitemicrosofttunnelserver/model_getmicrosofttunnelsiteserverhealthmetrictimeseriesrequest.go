package microsofttunnelsitemicrosofttunnelserver

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesRequest struct {
	EndDateTime   *string               `json:"endDateTime,omitempty"`
	MetricName    nullable.Type[string] `json:"metricName,omitempty"`
	StartDateTime *string               `json:"startDateTime,omitempty"`
}
