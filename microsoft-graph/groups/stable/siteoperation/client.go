package siteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOperationClient struct {
	Client *msgraph.Client
}

func NewSiteOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOperationClient: %+v", err)
	}

	return &SiteOperationClient{
		Client: client,
	}, nil
}
