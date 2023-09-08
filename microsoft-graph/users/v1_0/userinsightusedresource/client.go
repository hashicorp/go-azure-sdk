package userinsightusedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightUsedResourceClient struct {
	Client *msgraph.Client
}

func NewUserInsightUsedResourceClientWithBaseURI(api sdkEnv.Api) (*UserInsightUsedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsightusedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightUsedResourceClient: %+v", err)
	}

	return &UserInsightUsedResourceClient{
		Client: client,
	}, nil
}
