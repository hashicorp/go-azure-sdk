package groupsitetermstoregroupsetchildrenchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient{
		Client: client,
	}, nil
}
