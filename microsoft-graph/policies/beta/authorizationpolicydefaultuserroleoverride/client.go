package authorizationpolicydefaultuserroleoverride

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicyDefaultUserRoleOverrideClient struct {
	Client *msgraph.Client
}

func NewAuthorizationPolicyDefaultUserRoleOverrideClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthorizationPolicyDefaultUserRoleOverrideClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authorizationpolicydefaultuserroleoverride", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationPolicyDefaultUserRoleOverrideClient: %+v", err)
	}

	return &AuthorizationPolicyDefaultUserRoleOverrideClient{
		Client: client,
	}, nil
}
