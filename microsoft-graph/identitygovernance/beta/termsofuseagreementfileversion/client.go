package termsofuseagreementfileversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseAgreementFileVersionClient struct {
	Client *msgraph.Client
}

func NewTermsOfUseAgreementFileVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsOfUseAgreementFileVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsofuseagreementfileversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsOfUseAgreementFileVersionClient: %+v", err)
	}

	return &TermsOfUseAgreementFileVersionClient{
		Client: client,
	}, nil
}
