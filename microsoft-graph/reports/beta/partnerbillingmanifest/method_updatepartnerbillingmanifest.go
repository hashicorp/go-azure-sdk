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

type UpdatePartnerBillingManifestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePartnerBillingManifestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdatePartnerBillingManifestOperationOptions() UpdatePartnerBillingManifestOperationOptions {
	return UpdatePartnerBillingManifestOperationOptions{}
}

func (o UpdatePartnerBillingManifestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePartnerBillingManifestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePartnerBillingManifestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePartnerBillingManifest - Update the navigation property manifests in reports
func (c PartnerBillingManifestClient) UpdatePartnerBillingManifest(ctx context.Context, id beta.ReportPartnerBillingManifestId, input beta.PartnersBillingManifest, options UpdatePartnerBillingManifestOperationOptions) (result UpdatePartnerBillingManifestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
