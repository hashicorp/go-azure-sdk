package grouponenotesection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionClient: %+v", err)
	}

	return &GroupOnenoteSectionClient{
		Client: client,
	}, nil
}
