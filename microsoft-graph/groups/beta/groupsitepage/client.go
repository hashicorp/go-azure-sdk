package groupsitepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSitePageClient struct {
	Client *msgraph.Client
}

func NewGroupSitePageClientWithBaseURI(api sdkEnv.Api) (*GroupSitePageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSitePageClient: %+v", err)
	}

	return &GroupSitePageClient{
		Client: client,
	}, nil
}
