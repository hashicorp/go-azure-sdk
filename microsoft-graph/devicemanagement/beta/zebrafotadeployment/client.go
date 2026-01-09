package zebrafotadeployment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentClient struct {
	Client *msgraph.Client
}

func NewZebraFotaDeploymentClientWithBaseURI(sdkApi sdkEnv.Api) (*ZebraFotaDeploymentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "zebrafotadeployment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ZebraFotaDeploymentClient: %+v", err)
	}

	return &ZebraFotaDeploymentClient{
		Client: client,
	}, nil
}
