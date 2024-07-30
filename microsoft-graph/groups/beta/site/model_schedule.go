package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Schedule struct {
	DayNotes                  *[]DayNote                 `json:"dayNotes,omitempty"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Id                        *string                    `json:"id,omitempty"`
	ODataType                 *string                    `json:"@odata.type,omitempty"`
	OfferShiftRequests        *[]OfferShiftRequest       `json:"offerShiftRequests,omitempty"`
	OfferShiftRequestsEnabled *bool                      `json:"offerShiftRequestsEnabled,omitempty"`
	OpenShiftChangeRequests   *[]OpenShiftChangeRequest  `json:"openShiftChangeRequests,omitempty"`
	OpenShifts                *[]OpenShift               `json:"openShifts,omitempty"`
	OpenShiftsEnabled         *bool                      `json:"openShiftsEnabled,omitempty"`
	ProvisionStatus           *ScheduleProvisionStatus   `json:"provisionStatus,omitempty"`
	ProvisionStatusCode       *string                    `json:"provisionStatusCode,omitempty"`
	SchedulingGroups          *[]SchedulingGroup         `json:"schedulingGroups,omitempty"`
	Shifts                    *[]Shift                   `json:"shifts,omitempty"`
	SwapShiftsChangeRequests  *[]SwapShiftsChangeRequest `json:"swapShiftsChangeRequests,omitempty"`
	SwapShiftsRequestsEnabled *bool                      `json:"swapShiftsRequestsEnabled,omitempty"`
	TimeCards                 *[]TimeCard                `json:"timeCards,omitempty"`
	TimeClockEnabled          *bool                      `json:"timeClockEnabled,omitempty"`
	TimeClockSettings         *TimeClockSettings         `json:"timeClockSettings,omitempty"`
	TimeOffReasons            *[]TimeOffReason           `json:"timeOffReasons,omitempty"`
	TimeOffRequests           *[]TimeOffRequest          `json:"timeOffRequests,omitempty"`
	TimeOffRequestsEnabled    *bool                      `json:"timeOffRequestsEnabled,omitempty"`
	TimeZone                  *string                    `json:"timeZone,omitempty"`
	TimesOff                  *[]TimeOff                 `json:"timesOff,omitempty"`
	WorkforceIntegrationIds   *[]string                  `json:"workforceIntegrationIds,omitempty"`
}
