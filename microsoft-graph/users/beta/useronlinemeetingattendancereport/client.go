package useronlinemeetingattendancereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingAttendanceReportClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingAttendanceReportClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingAttendanceReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingattendancereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingAttendanceReportClient: %+v", err)
	}

	return &UserOnlineMeetingAttendanceReportClient{
		Client: client,
	}, nil
}
