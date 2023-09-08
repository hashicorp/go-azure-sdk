package identityb2xuserflowlanguagedefaultpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowLanguageDefaultPageClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowLanguageDefaultPageClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowLanguageDefaultPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowlanguagedefaultpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowLanguageDefaultPageClient: %+v", err)
	}

	return &IdentityB2xUserFlowLanguageDefaultPageClient{
		Client: client,
	}, nil
}
