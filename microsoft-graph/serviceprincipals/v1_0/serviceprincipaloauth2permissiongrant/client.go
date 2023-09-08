package serviceprincipaloauth2permissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalOauth2PermissionGrantClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalOauth2PermissionGrantClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalOauth2PermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipaloauth2permissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalOauth2PermissionGrantClient: %+v", err)
	}

	return &ServicePrincipalOauth2PermissionGrantClient{
		Client: client,
	}, nil
}
