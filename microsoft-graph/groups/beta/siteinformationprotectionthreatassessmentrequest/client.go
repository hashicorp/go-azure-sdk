package siteinformationprotectionthreatassessmentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionThreatAssessmentRequestClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionThreatAssessmentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionthreatassessmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionThreatAssessmentRequestClient: %+v", err)
	}

	return &SiteInformationProtectionThreatAssessmentRequestClient{
		Client: client,
	}, nil
}
