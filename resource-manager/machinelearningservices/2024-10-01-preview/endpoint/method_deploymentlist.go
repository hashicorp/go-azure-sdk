package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointDeploymentResourcePropertiesBasicResource
}

type DeploymentListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointDeploymentResourcePropertiesBasicResource
}

type DeploymentListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeploymentListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeploymentList ...
func (c EndpointClient) DeploymentList(ctx context.Context, id EndpointId) (result DeploymentListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DeploymentListCustomPager{},
		Path:       fmt.Sprintf("%s/deployments", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]EndpointDeploymentResourcePropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeploymentListComplete retrieves all the results into a single object
func (c EndpointClient) DeploymentListComplete(ctx context.Context, id EndpointId) (DeploymentListCompleteResult, error) {
	return c.DeploymentListCompleteMatchingPredicate(ctx, id, EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate{})
}

// DeploymentListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) DeploymentListCompleteMatchingPredicate(ctx context.Context, id EndpointId, predicate EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate) (result DeploymentListCompleteResult, err error) {
	items := make([]EndpointDeploymentResourcePropertiesBasicResource, 0)

	resp, err := c.DeploymentList(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = DeploymentListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
