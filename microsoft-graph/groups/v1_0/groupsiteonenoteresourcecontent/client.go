package groupsiteonenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteResourceContentClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteResourceContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteResourceContentClient: %+v", err)
	}

	return &GroupSiteOnenoteResourceContentClient{
		Client: client,
	}, nil
}
