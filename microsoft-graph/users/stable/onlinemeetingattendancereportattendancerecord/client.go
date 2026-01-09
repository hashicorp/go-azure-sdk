package onlinemeetingattendancereportattendancerecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAttendanceReportAttendanceRecordClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingAttendanceReportAttendanceRecordClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingattendancereportattendancerecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingAttendanceReportAttendanceRecordClient: %+v", err)
	}

	return &OnlineMeetingAttendanceReportAttendanceRecordClient{
		Client: client,
	}, nil
}
