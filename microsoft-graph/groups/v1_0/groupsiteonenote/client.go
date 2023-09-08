package groupsiteonenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteClient: %+v", err)
	}

	return &GroupSiteOnenoteClient{
		Client: client,
	}, nil
}
