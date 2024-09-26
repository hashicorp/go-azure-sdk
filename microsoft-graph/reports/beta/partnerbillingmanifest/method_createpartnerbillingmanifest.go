package partnerbillingmanifest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingManifestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PartnersBillingManifest
}

type CreatePartnerBillingManifestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePartnerBillingManifestOperationOptions() CreatePartnerBillingManifestOperationOptions {
	return CreatePartnerBillingManifestOperationOptions{}
}

func (o CreatePartnerBillingManifestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePartnerBillingManifestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePartnerBillingManifestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePartnerBillingManifest - Create new navigation property to manifests for reports
func (c PartnerBillingManifestClient) CreatePartnerBillingManifest(ctx context.Context, input beta.PartnersBillingManifest, options CreatePartnerBillingManifestOperationOptions) (result CreatePartnerBillingManifestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/partners/billing/manifests",
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

	var model beta.PartnersBillingManifest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
