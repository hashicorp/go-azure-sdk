package useronlinemeetingattendancereportattendancerecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingAttendanceReportAttendanceRecordClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingAttendanceReportAttendanceRecordClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingattendancereportattendancerecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingAttendanceReportAttendanceRecordClient: %+v", err)
	}

	return &UserOnlineMeetingAttendanceReportAttendanceRecordClient{
		Client: client,
	}, nil
}
