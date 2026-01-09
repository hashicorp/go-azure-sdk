package createdobject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatedObjectClient struct {
	Client *msgraph.Client
}

func NewCreatedObjectClientWithBaseURI(sdkApi sdkEnv.Api) (*CreatedObjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "createdobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CreatedObjectClient: %+v", err)
	}

	return &CreatedObjectClient{
		Client: client,
	}, nil
}
