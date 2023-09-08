package groupsitetermstoregroupsettermchildrenrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermChildrenRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermChildrenRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermChildrenRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermchildrenrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermChildrenRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermChildrenRelationFromTermClient{
		Client: client,
	}, nil
}
