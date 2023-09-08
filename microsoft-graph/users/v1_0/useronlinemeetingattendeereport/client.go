package useronlinemeetingattendeereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingAttendeeReportClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingAttendeeReportClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingAttendeeReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingattendeereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingAttendeeReportClient: %+v", err)
	}

	return &UserOnlineMeetingAttendeeReportClient{
		Client: client,
	}, nil
}
