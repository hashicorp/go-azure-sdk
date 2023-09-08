package groupsitetermstoresetchildrenrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenRelationFromTermClient{
		Client: client,
	}, nil
}
