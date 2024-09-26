package recommendationimpactedresource

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRecommendationImpactedResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateRecommendationImpactedResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateRecommendationImpactedResourceOperationOptions() UpdateRecommendationImpactedResourceOperationOptions {
	return UpdateRecommendationImpactedResourceOperationOptions{}
}

func (o UpdateRecommendationImpactedResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateRecommendationImpactedResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateRecommendationImpactedResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateRecommendationImpactedResource - Update the navigation property impactedResources in directory
func (c RecommendationImpactedResourceClient) UpdateRecommendationImpactedResource(ctx context.Context, id beta.DirectoryRecommendationIdImpactedResourceId, input beta.ImpactedResource, options UpdateRecommendationImpactedResourceOperationOptions) (result UpdateRecommendationImpactedResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
