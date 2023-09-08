package groupsitetermstoresettermchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetTermChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetTermChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetTermChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresettermchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetTermChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetTermChildrenRelationClient{
		Client: client,
	}, nil
}
