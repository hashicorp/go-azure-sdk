package alerts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	ActionTaken        *string                  `json:"actionTaken,omitempty"`
	AlertDisplayName   *string                  `json:"alertDisplayName,omitempty"`
	AlertName          *string                  `json:"alertName,omitempty"`
	AssociatedResource *string                  `json:"associatedResource,omitempty"`
	CanBeInvestigated  *bool                    `json:"canBeInvestigated,omitempty"`
	CompromisedEntity  *string                  `json:"compromisedEntity,omitempty"`
	ConfidenceReasons  *[]AlertConfidenceReason `json:"confidenceReasons,omitempty"`
	ConfidenceScore    *float64                 `json:"confidenceScore,omitempty"`
	CorrelationKey     *string                  `json:"correlationKey,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	DetectedTimeUtc    *string                  `json:"detectedTimeUtc,omitempty"`
	Entities           *[]AlertEntity           `json:"entities,omitempty"`
	ExtendedProperties *interface{}             `json:"extendedProperties,omitempty"`
	InstanceId         *string                  `json:"instanceId,omitempty"`
	IsIncident         *bool                    `json:"isIncident,omitempty"`
	RemediationSteps   *string                  `json:"remediationSteps,omitempty"`
	ReportedSeverity   *ReportedSeverity        `json:"reportedSeverity,omitempty"`
	ReportedTimeUtc    *string                  `json:"reportedTimeUtc,omitempty"`
	State              *string                  `json:"state,omitempty"`
	SubscriptionId     *string                  `json:"subscriptionId,omitempty"`
	SystemSource       *string                  `json:"systemSource,omitempty"`
	VendorName         *string                  `json:"vendorName,omitempty"`
	WorkspaceArmId     *string                  `json:"workspaceArmId,omitempty"`
}

func (o *AlertProperties) GetDetectedTimeUtcAsTime() (*time.Time, error) {
	if o.DetectedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DetectedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertProperties) SetDetectedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DetectedTimeUtc = &formatted
}

func (o *AlertProperties) GetReportedTimeUtcAsTime() (*time.Time, error) {
	if o.ReportedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReportedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertProperties) SetReportedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReportedTimeUtc = &formatted
}
