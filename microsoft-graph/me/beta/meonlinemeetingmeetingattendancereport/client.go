package meonlinemeetingmeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingMeetingAttendanceReportClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingmeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingMeetingAttendanceReportClient: %+v", err)
	}

	return &MeOnlineMeetingMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
