package setgroupsitetermstoregroupsetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetRelationClient{
		Client: client,
	}, nil
}
