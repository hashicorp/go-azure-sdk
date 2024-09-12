package microsofttunnelsiteserver

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesRequest struct {
	EndDateTime   *string               `json:"endDateTime,omitempty"`
	MetricName    nullable.Type[string] `json:"metricName,omitempty"`
	StartDateTime *string               `json:"startDateTime,omitempty"`
}
