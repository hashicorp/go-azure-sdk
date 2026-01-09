package onlinemeetingmeetingattendancereportattendancerecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingMeetingAttendanceReportAttendanceRecordClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingMeetingAttendanceReportAttendanceRecordClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingmeetingattendancereportattendancerecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingMeetingAttendanceReportAttendanceRecordClient: %+v", err)
	}

	return &OnlineMeetingMeetingAttendanceReportAttendanceRecordClient{
		Client: client,
	}, nil
}
