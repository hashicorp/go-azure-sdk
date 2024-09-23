package contact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactClient struct {
	Client *msgraph.Client
}

func NewContactClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactClient: %+v", err)
	}

	return &ContactClient{
		Client: client,
	}, nil
}
