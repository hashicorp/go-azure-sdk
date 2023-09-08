package grouponenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteResourceClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteResourceClient: %+v", err)
	}

	return &GroupOnenoteResourceClient{
		Client: client,
	}, nil
}
