package grouponenotesectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupClient{
		Client: client,
	}, nil
}
