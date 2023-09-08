package groupsitelistitemdocumentsetversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemDocumentSetVersionFieldClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemDocumentSetVersionFieldClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemDocumentSetVersionFieldClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemdocumentsetversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemDocumentSetVersionFieldClient: %+v", err)
	}

	return &GroupSiteListItemDocumentSetVersionFieldClient{
		Client: client,
	}, nil
}
