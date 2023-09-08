package groupsitetermstoresettermrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetTermRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetTermRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetTermRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresettermrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetTermRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetTermRelationFromTermClient{
		Client: client,
	}, nil
}
