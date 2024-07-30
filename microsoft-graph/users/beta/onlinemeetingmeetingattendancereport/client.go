package onlinemeetingmeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingMeetingAttendanceReportClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingmeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingMeetingAttendanceReportClient: %+v", err)
	}

	return &OnlineMeetingMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
