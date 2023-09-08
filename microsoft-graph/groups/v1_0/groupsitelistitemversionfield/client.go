package groupsitelistitemversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemVersionFieldClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemVersionFieldClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemVersionFieldClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemVersionFieldClient: %+v", err)
	}

	return &GroupSiteListItemVersionFieldClient{
		Client: client,
	}, nil
}
