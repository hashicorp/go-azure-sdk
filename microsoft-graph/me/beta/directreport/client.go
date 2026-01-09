package directreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectReportClient struct {
	Client *msgraph.Client
}

func NewDirectReportClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectReportClient: %+v", err)
	}

	return &DirectReportClient{
		Client: client,
	}, nil
}
