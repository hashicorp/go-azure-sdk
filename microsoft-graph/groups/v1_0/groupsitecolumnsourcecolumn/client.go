package groupsitecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteColumnSourceColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteColumnSourceColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteColumnSourceColumnClient: %+v", err)
	}

	return &GroupSiteColumnSourceColumnClient{
		Client: client,
	}, nil
}
