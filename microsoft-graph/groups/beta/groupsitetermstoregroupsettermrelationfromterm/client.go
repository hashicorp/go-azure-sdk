package groupsitetermstoregroupsettermrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermRelationFromTermClient{
		Client: client,
	}, nil
}
