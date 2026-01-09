package report

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportClient struct {
	Client *msgraph.Client
}

func NewReportClientWithBaseURI(sdkApi sdkEnv.Api) (*ReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "report", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReportClient: %+v", err)
	}

	return &ReportClient{
		Client: client,
	}, nil
}
