package incidentbookmarks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailClusterEntityProperties struct {
	AdditionalData          *interface{} `json:"additionalData,omitempty"`
	ClusterGroup            *string      `json:"clusterGroup,omitempty"`
	ClusterQueryEndTime     *string      `json:"clusterQueryEndTime,omitempty"`
	ClusterQueryStartTime   *string      `json:"clusterQueryStartTime,omitempty"`
	ClusterSourceIdentifier *string      `json:"clusterSourceIdentifier,omitempty"`
	ClusterSourceType       *string      `json:"clusterSourceType,omitempty"`
	CountByDeliveryStatus   *interface{} `json:"countByDeliveryStatus,omitempty"`
	CountByProtectionStatus *interface{} `json:"countByProtectionStatus,omitempty"`
	CountByThreatType       *interface{} `json:"countByThreatType,omitempty"`
	FriendlyName            *string      `json:"friendlyName,omitempty"`
	IsVolumeAnomaly         *bool        `json:"isVolumeAnomaly,omitempty"`
	MailCount               *int64       `json:"mailCount,omitempty"`
	NetworkMessageIds       *[]string    `json:"networkMessageIds,omitempty"`
	Query                   *string      `json:"query,omitempty"`
	QueryTime               *string      `json:"queryTime,omitempty"`
	Source                  *string      `json:"source,omitempty"`
	Threats                 *[]string    `json:"threats,omitempty"`
}

func (o *MailClusterEntityProperties) GetClusterQueryEndTimeAsTime() (*time.Time, error) {
	if o.ClusterQueryEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ClusterQueryEndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MailClusterEntityProperties) SetClusterQueryEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ClusterQueryEndTime = &formatted
}

func (o *MailClusterEntityProperties) GetClusterQueryStartTimeAsTime() (*time.Time, error) {
	if o.ClusterQueryStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ClusterQueryStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MailClusterEntityProperties) SetClusterQueryStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ClusterQueryStartTime = &formatted
}

func (o *MailClusterEntityProperties) GetQueryTimeAsTime() (*time.Time, error) {
	if o.QueryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.QueryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MailClusterEntityProperties) SetQueryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.QueryTime = &formatted
}
