package setgroupsitetermstoregroupsetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetChildrenChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
