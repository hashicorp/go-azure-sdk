package currentquotalimitbases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentQuotaLimitBasesClient struct {
	Client *resourcemanager.Client
}

func NewCurrentQuotaLimitBasesClientWithBaseURI(sdkApi sdkEnv.Api) (*CurrentQuotaLimitBasesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "currentquotalimitbases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CurrentQuotaLimitBasesClient: %+v", err)
	}

	return &CurrentQuotaLimitBasesClient{
		Client: client,
	}, nil
}
