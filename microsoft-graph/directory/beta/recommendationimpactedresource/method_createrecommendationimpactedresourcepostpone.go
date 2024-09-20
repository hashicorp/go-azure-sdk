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

type CreateRecommendationImpactedResourcePostponeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ImpactedResource
}

type CreateRecommendationImpactedResourcePostponeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateRecommendationImpactedResourcePostponeOperationOptions() CreateRecommendationImpactedResourcePostponeOperationOptions {
	return CreateRecommendationImpactedResourcePostponeOperationOptions{}
}

func (o CreateRecommendationImpactedResourcePostponeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRecommendationImpactedResourcePostponeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRecommendationImpactedResourcePostponeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRecommendationImpactedResourcePostpone - Invoke action postpone. Postpone action on an impactedResource object
// to a specified future date and time by marking its status as postponed. On the specified date and time, Microsoft
// Entra ID will automatically mark the status of the impactedResource object to active.
func (c RecommendationImpactedResourceClient) CreateRecommendationImpactedResourcePostpone(ctx context.Context, id beta.DirectoryRecommendationIdImpactedResourceId, input CreateRecommendationImpactedResourcePostponeRequest, options CreateRecommendationImpactedResourcePostponeOperationOptions) (result CreateRecommendationImpactedResourcePostponeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/postpone", id.ID()),
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

	var model beta.ImpactedResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
