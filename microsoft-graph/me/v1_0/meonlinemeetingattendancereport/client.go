package meonlinemeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingAttendanceReportClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingAttendanceReportClient: %+v", err)
	}

	return &MeOnlineMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
