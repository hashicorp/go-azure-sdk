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

type DismissRecommendationImpactedResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ImpactedResource
}

type DismissRecommendationImpactedResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDismissRecommendationImpactedResourceOperationOptions() DismissRecommendationImpactedResourceOperationOptions {
	return DismissRecommendationImpactedResourceOperationOptions{}
}

func (o DismissRecommendationImpactedResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissRecommendationImpactedResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissRecommendationImpactedResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissRecommendationImpactedResource - Invoke action dismiss. Dismiss an impactedResources object and update its
// status to dismissed.
func (c RecommendationImpactedResourceClient) DismissRecommendationImpactedResource(ctx context.Context, id beta.DirectoryRecommendationIdImpactedResourceId, input DismissRecommendationImpactedResourceRequest, options DismissRecommendationImpactedResourceOperationOptions) (result DismissRecommendationImpactedResourceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/dismiss", id.ID()),
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

	var model beta.ImpactedResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
