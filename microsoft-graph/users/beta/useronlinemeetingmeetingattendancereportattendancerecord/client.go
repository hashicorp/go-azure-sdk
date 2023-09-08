package useronlinemeetingmeetingattendancereportattendancerecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingMeetingAttendanceReportAttendanceRecordClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingMeetingAttendanceReportAttendanceRecordClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingMeetingAttendanceReportAttendanceRecordClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingmeetingattendancereportattendancerecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingMeetingAttendanceReportAttendanceRecordClient: %+v", err)
	}

	return &UserOnlineMeetingMeetingAttendanceReportAttendanceRecordClient{
		Client: client,
	}, nil
}
