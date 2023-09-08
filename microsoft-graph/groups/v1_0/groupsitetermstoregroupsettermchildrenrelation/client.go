package groupsitetermstoregroupsettermchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermChildrenRelationClient{
		Client: client,
	}, nil
}
