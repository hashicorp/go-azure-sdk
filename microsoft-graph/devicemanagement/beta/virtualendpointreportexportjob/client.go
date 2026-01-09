package virtualendpointreportexportjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointReportExportJobClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointReportExportJobClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointReportExportJobClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointreportexportjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointReportExportJobClient: %+v", err)
	}

	return &VirtualEndpointReportExportJobClient{
		Client: client,
	}, nil
}
