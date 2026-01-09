package siteonenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteOperationClient: %+v", err)
	}

	return &SiteOnenoteOperationClient{
		Client: client,
	}, nil
}
