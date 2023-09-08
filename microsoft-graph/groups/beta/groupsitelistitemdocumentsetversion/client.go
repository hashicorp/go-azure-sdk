package groupsitelistitemdocumentsetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemDocumentSetVersionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemDocumentSetVersionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemDocumentSetVersionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemdocumentsetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemDocumentSetVersionClient: %+v", err)
	}

	return &GroupSiteListItemDocumentSetVersionClient{
		Client: client,
	}, nil
}
