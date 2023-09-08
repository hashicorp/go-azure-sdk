package meinsightsharedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInsightSharedResourceClient struct {
	Client *msgraph.Client
}

func NewMeInsightSharedResourceClientWithBaseURI(api sdkEnv.Api) (*MeInsightSharedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinsightsharedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInsightSharedResourceClient: %+v", err)
	}

	return &MeInsightSharedResourceClient{
		Client: client,
	}, nil
}
