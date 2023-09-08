package groupsiteonenoteresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteResourceClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteResourceClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenoteresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteResourceClient: %+v", err)
	}

	return &GroupSiteOnenoteResourceClient{
		Client: client,
	}, nil
}
