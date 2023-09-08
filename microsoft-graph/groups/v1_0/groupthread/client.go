package groupthread

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadClient struct {
	Client *msgraph.Client
}

func NewGroupThreadClientWithBaseURI(api sdkEnv.Api) (*GroupThreadClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthread", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadClient: %+v", err)
	}

	return &GroupThreadClient{
		Client: client,
	}, nil
}
