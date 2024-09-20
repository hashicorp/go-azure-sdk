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

type CreateRecommendationImpactedResourceReactivateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ImpactedResource
}

type CreateRecommendationImpactedResourceReactivateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateRecommendationImpactedResourceReactivateOperationOptions() CreateRecommendationImpactedResourceReactivateOperationOptions {
	return CreateRecommendationImpactedResourceReactivateOperationOptions{}
}

func (o CreateRecommendationImpactedResourceReactivateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRecommendationImpactedResourceReactivateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRecommendationImpactedResourceReactivateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRecommendationImpactedResourceReactivate - Invoke action reactivate. Reactivate an accidentally dismissed,
// completed, or postponed impactedResource object. This action updates the status of the resource to active. This
// method is relevant only if the status of the impactedResource object is dismissed, postponed, or completedByUser.
func (c RecommendationImpactedResourceClient) CreateRecommendationImpactedResourceReactivate(ctx context.Context, id beta.DirectoryRecommendationIdImpactedResourceId, options CreateRecommendationImpactedResourceReactivateOperationOptions) (result CreateRecommendationImpactedResourceReactivateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reactivate", id.ID()),
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
