package virtualendpointreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointReportClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointReportClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointReportClient: %+v", err)
	}

	return &VirtualEndpointReportClient{
		Client: client,
	}, nil
}
