package meonlinemeetingattendeereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingAttendeeReportClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingAttendeeReportClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingAttendeeReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingattendeereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingAttendeeReportClient: %+v", err)
	}

	return &MeOnlineMeetingAttendeeReportClient{
		Client: client,
	}, nil
}
