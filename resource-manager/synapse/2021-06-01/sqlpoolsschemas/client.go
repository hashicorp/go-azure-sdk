package sqlpoolsschemas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSchemasClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSchemasClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSchemasClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolsschemas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSchemasClient: %+v", err)
	}

	return &SqlPoolsSchemasClient{
		Client: client,
	}, nil
}
