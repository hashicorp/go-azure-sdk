package defender

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderClient struct {
	Client *msgraph.Client
}

func NewDefenderClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defender", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderClient: %+v", err)
	}

	return &DefenderClient{
		Client: client,
	}, nil
}
