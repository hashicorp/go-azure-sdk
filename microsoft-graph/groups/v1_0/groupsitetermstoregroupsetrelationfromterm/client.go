package groupsitetermstoregroupsetrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetRelationFromTermClient{
		Client: client,
	}, nil
}
