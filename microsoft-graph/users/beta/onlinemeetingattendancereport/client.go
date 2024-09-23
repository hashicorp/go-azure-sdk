package onlinemeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingAttendanceReportClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingAttendanceReportClient: %+v", err)
	}

	return &OnlineMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
