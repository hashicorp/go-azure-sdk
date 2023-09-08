package groupsitecontenttype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteContentTypeClient struct {
	Client *msgraph.Client
}

func NewGroupSiteContentTypeClientWithBaseURI(api sdkEnv.Api) (*GroupSiteContentTypeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecontenttype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteContentTypeClient: %+v", err)
	}

	return &GroupSiteContentTypeClient{
		Client: client,
	}, nil
}
