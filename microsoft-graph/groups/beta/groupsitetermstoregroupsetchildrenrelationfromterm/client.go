package groupsitetermstoregroupsetchildrenrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenRelationFromTermClient{
		Client: client,
	}, nil
}
