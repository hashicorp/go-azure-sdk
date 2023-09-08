package groupthreadpost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostClient: %+v", err)
	}

	return &GroupThreadPostClient{
		Client: client,
	}, nil
}
