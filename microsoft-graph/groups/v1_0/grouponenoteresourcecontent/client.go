package grouponenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteResourceContentClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteResourceContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteResourceContentClient: %+v", err)
	}

	return &GroupOnenoteResourceContentClient{
		Client: client,
	}, nil
}
