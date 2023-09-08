package meinsightusedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInsightUsedResourceClient struct {
	Client *msgraph.Client
}

func NewMeInsightUsedResourceClientWithBaseURI(api sdkEnv.Api) (*MeInsightUsedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinsightusedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInsightUsedResourceClient: %+v", err)
	}

	return &MeInsightUsedResourceClient{
		Client: client,
	}, nil
}
