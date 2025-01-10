package scalingplanpersonalschedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingPlanPersonalScheduleProperties struct {
	DaysOfWeek                        *[]DayOfWeek              `json:"daysOfWeek,omitempty"`
	OffPeakActionOnDisconnect         *SessionHandlingOperation `json:"offPeakActionOnDisconnect,omitempty"`
	OffPeakActionOnLogoff             *SessionHandlingOperation `json:"offPeakActionOnLogoff,omitempty"`
	OffPeakMinutesToWaitOnDisconnect  *int64                    `json:"offPeakMinutesToWaitOnDisconnect,omitempty"`
	OffPeakMinutesToWaitOnLogoff      *int64                    `json:"offPeakMinutesToWaitOnLogoff,omitempty"`
	OffPeakStartTime                  *Time                     `json:"offPeakStartTime,omitempty"`
	OffPeakStartVMOnConnect           *SetStartVMOnConnect      `json:"offPeakStartVMOnConnect,omitempty"`
	PeakActionOnDisconnect            *SessionHandlingOperation `json:"peakActionOnDisconnect,omitempty"`
	PeakActionOnLogoff                *SessionHandlingOperation `json:"peakActionOnLogoff,omitempty"`
	PeakMinutesToWaitOnDisconnect     *int64                    `json:"peakMinutesToWaitOnDisconnect,omitempty"`
	PeakMinutesToWaitOnLogoff         *int64                    `json:"peakMinutesToWaitOnLogoff,omitempty"`
	PeakStartTime                     *Time                     `json:"peakStartTime,omitempty"`
	PeakStartVMOnConnect              *SetStartVMOnConnect      `json:"peakStartVMOnConnect,omitempty"`
	RampDownActionOnDisconnect        *SessionHandlingOperation `json:"rampDownActionOnDisconnect,omitempty"`
	RampDownActionOnLogoff            *SessionHandlingOperation `json:"rampDownActionOnLogoff,omitempty"`
	RampDownMinutesToWaitOnDisconnect *int64                    `json:"rampDownMinutesToWaitOnDisconnect,omitempty"`
	RampDownMinutesToWaitOnLogoff     *int64                    `json:"rampDownMinutesToWaitOnLogoff,omitempty"`
	RampDownStartTime                 *Time                     `json:"rampDownStartTime,omitempty"`
	RampDownStartVMOnConnect          *SetStartVMOnConnect      `json:"rampDownStartVMOnConnect,omitempty"`
	RampUpActionOnDisconnect          *SessionHandlingOperation `json:"rampUpActionOnDisconnect,omitempty"`
	RampUpActionOnLogoff              *SessionHandlingOperation `json:"rampUpActionOnLogoff,omitempty"`
	RampUpAutoStartHosts              *StartupBehavior          `json:"rampUpAutoStartHosts,omitempty"`
	RampUpMinutesToWaitOnDisconnect   *int64                    `json:"rampUpMinutesToWaitOnDisconnect,omitempty"`
	RampUpMinutesToWaitOnLogoff       *int64                    `json:"rampUpMinutesToWaitOnLogoff,omitempty"`
	RampUpStartTime                   *Time                     `json:"rampUpStartTime,omitempty"`
	RampUpStartVMOnConnect            *SetStartVMOnConnect      `json:"rampUpStartVMOnConnect,omitempty"`
}
