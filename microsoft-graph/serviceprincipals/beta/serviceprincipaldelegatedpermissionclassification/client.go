package serviceprincipaldelegatedpermissionclassification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalDelegatedPermissionClassificationClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalDelegatedPermissionClassificationClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalDelegatedPermissionClassificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipaldelegatedpermissionclassification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalDelegatedPermissionClassificationClient: %+v", err)
	}

	return &ServicePrincipalDelegatedPermissionClassificationClient{
		Client: client,
	}, nil
}
