package metransitivereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTransitiveReportClient struct {
	Client *msgraph.Client
}

func NewMeTransitiveReportClientWithBaseURI(api sdkEnv.Api) (*MeTransitiveReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metransitivereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTransitiveReportClient: %+v", err)
	}

	return &MeTransitiveReportClient{
		Client: client,
	}, nil
}
