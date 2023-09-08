package useronlinemeetingmeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingMeetingAttendanceReportClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingmeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingMeetingAttendanceReportClient: %+v", err)
	}

	return &UserOnlineMeetingMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
