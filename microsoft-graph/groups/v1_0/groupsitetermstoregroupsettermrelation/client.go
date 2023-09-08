package groupsitetermstoregroupsettermrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermRelationClient{
		Client: client,
	}, nil
}
