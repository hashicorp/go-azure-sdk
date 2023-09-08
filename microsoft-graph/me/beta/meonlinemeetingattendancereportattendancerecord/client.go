package meonlinemeetingattendancereportattendancerecord

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingAttendanceReportAttendanceRecordClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingAttendanceReportAttendanceRecordClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingattendancereportattendancerecord", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingAttendanceReportAttendanceRecordClient: %+v", err)
	}

	return &MeOnlineMeetingAttendanceReportAttendanceRecordClient{
		Client: client,
	}, nil
}
