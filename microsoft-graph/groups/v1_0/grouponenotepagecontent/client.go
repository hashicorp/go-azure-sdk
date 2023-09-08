package grouponenotepagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenotePageContentClient struct {
	Client *msgraph.Client
}

func NewGroupOnenotePageContentClientWithBaseURI(api sdkEnv.Api) (*GroupOnenotePageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotepagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenotePageContentClient: %+v", err)
	}

	return &GroupOnenotePageContentClient{
		Client: client,
	}, nil
}
