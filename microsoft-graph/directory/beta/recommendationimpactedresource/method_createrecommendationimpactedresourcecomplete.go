package recommendationimpactedresource

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

type CreateRecommendationImpactedResourceCompleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ImpactedResource
}

type CreateRecommendationImpactedResourceCompleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateRecommendationImpactedResourceCompleteOperationOptions() CreateRecommendationImpactedResourceCompleteOperationOptions {
	return CreateRecommendationImpactedResourceCompleteOperationOptions{}
}

func (o CreateRecommendationImpactedResourceCompleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRecommendationImpactedResourceCompleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRecommendationImpactedResourceCompleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRecommendationImpactedResourceComplete - Invoke action complete. Complete an impactedResource object and update
// its status to completedByUser.
func (c RecommendationImpactedResourceClient) CreateRecommendationImpactedResourceComplete(ctx context.Context, id beta.DirectoryRecommendationIdImpactedResourceId, options CreateRecommendationImpactedResourceCompleteOperationOptions) (result CreateRecommendationImpactedResourceCompleteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/complete", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model beta.ImpactedResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
