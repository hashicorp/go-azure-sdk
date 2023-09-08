package meinsightused

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInsightUsedClient struct {
	Client *msgraph.Client
}

func NewMeInsightUsedClientWithBaseURI(api sdkEnv.Api) (*MeInsightUsedClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinsightused", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInsightUsedClient: %+v", err)
	}

	return &MeInsightUsedClient{
		Client: client,
	}, nil
}
