package dppbaseresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceGuardsGetUpdateProtectionPolicyRequestsObjects ...
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectionPolicyRequestsObjects(ctx context.Context, id ResourceGuardId) (result ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCustomPager{},
		Path:       fmt.Sprintf("%s/updateProtectionPolicyRequests", id.ID()),
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
		Values *[]Resource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsComplete retrieves all the results into a single object
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsComplete(ctx context.Context, id ResourceGuardId) (ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteResult, error) {
	return c.ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteMatchingPredicate(ctx context.Context, id ResourceGuardId, predicate ResourceOperationPredicate) (result ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.ResourceGuardsGetUpdateProtectionPolicyRequestsObjects(ctx, id)
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

	result = ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
