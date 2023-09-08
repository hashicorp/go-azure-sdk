package identityb2xuserflowlanguageoverridespage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowLanguageOverridesPageClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowLanguageOverridesPageClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowLanguageOverridesPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowlanguageoverridespage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowLanguageOverridesPageClient: %+v", err)
	}

	return &IdentityB2xUserFlowLanguageOverridesPageClient{
		Client: client,
	}, nil
}
