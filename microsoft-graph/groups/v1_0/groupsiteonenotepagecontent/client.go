package groupsiteonenotepagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenotePageContentClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenotePageContentClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenotePageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotepagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenotePageContentClient: %+v", err)
	}

	return &GroupSiteOnenotePageContentClient{
		Client: client,
	}, nil
}
