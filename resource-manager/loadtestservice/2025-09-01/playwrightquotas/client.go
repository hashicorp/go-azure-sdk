package playwrightquotas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlaywrightQuotasClient struct {
	Client *resourcemanager.Client
}

func NewPlaywrightQuotasClientWithBaseURI(sdkApi sdkEnv.Api) (*PlaywrightQuotasClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "playwrightquotas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlaywrightQuotasClient: %+v", err)
	}

	return &PlaywrightQuotasClient{
		Client: client,
	}, nil
}
