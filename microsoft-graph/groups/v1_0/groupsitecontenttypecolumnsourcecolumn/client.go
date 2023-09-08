package groupsitecontenttypecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteContentTypeColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteContentTypeColumnSourceColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteContentTypeColumnSourceColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecontenttypecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteContentTypeColumnSourceColumnClient: %+v", err)
	}

	return &GroupSiteContentTypeColumnSourceColumnClient{
		Client: client,
	}, nil
}
