package siteinformationprotectionpolicylabel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionPolicyLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.InformationProtectionLabel
}

type CreateSiteInformationProtectionPolicyLabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSiteInformationProtectionPolicyLabelOperationOptions() CreateSiteInformationProtectionPolicyLabelOperationOptions {
	return CreateSiteInformationProtectionPolicyLabelOperationOptions{}
}

func (o CreateSiteInformationProtectionPolicyLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteInformationProtectionPolicyLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteInformationProtectionPolicyLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteInformationProtectionPolicyLabel - Create new navigation property to labels for groups
func (c SiteInformationProtectionPolicyLabelClient) CreateSiteInformationProtectionPolicyLabel(ctx context.Context, id beta.GroupIdSiteId, input beta.InformationProtectionLabel, options CreateSiteInformationProtectionPolicyLabelOperationOptions) (result CreateSiteInformationProtectionPolicyLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/informationProtection/policy/labels", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.InformationProtectionLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
