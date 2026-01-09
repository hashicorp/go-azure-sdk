package transitivereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransitiveReportClient struct {
	Client *msgraph.Client
}

func NewTransitiveReportClientWithBaseURI(sdkApi sdkEnv.Api) (*TransitiveReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "transitivereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransitiveReportClient: %+v", err)
	}

	return &TransitiveReportClient{
		Client: client,
	}, nil
}
