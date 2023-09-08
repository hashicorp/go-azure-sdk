package groupsitecontenttypebase

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteContentTypeBaseClient struct {
	Client *msgraph.Client
}

func NewGroupSiteContentTypeBaseClientWithBaseURI(api sdkEnv.Api) (*GroupSiteContentTypeBaseClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecontenttypebase", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteContentTypeBaseClient: %+v", err)
	}

	return &GroupSiteContentTypeBaseClient{
		Client: client,
	}, nil
}
