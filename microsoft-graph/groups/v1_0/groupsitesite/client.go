package groupsitesite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteSiteClient struct {
	Client *msgraph.Client
}

func NewGroupSiteSiteClientWithBaseURI(api sdkEnv.Api) (*GroupSiteSiteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitesite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteSiteClient: %+v", err)
	}

	return &GroupSiteSiteClient{
		Client: client,
	}, nil
}
