package groupsitecontenttypecolumnposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteContentTypeColumnPositionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteContentTypeColumnPositionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteContentTypeColumnPositionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecontenttypecolumnposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteContentTypeColumnPositionClient: %+v", err)
	}

	return &GroupSiteContentTypeColumnPositionClient{
		Client: client,
	}, nil
}
