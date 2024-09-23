package termsofuseagreement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseAgreementClient struct {
	Client *msgraph.Client
}

func NewTermsOfUseAgreementClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsOfUseAgreementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsofuseagreement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsOfUseAgreementClient: %+v", err)
	}

	return &TermsOfUseAgreementClient{
		Client: client,
	}, nil
}
