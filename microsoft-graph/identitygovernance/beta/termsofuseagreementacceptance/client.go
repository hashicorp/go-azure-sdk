package termsofuseagreementacceptance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseAgreementAcceptanceClient struct {
	Client *msgraph.Client
}

func NewTermsOfUseAgreementAcceptanceClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsOfUseAgreementAcceptanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsofuseagreementacceptance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsOfUseAgreementAcceptanceClient: %+v", err)
	}

	return &TermsOfUseAgreementAcceptanceClient{
		Client: client,
	}, nil
}
