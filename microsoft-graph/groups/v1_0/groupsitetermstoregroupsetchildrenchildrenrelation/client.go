package groupsitetermstoregroupsetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
