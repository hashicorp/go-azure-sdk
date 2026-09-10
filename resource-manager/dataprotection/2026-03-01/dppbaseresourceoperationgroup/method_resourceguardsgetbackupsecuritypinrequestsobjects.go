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

type ResourceGuardsGetBackupSecurityPINRequestsObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type ResourceGuardsGetBackupSecurityPINRequestsObjectsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceGuardsGetBackupSecurityPINRequestsObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceGuardsGetBackupSecurityPINRequestsObjects ...
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetBackupSecurityPINRequestsObjects(ctx context.Context, id ResourceGuardId) (result ResourceGuardsGetBackupSecurityPINRequestsObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceGuardsGetBackupSecurityPINRequestsObjectsCustomPager{},
		Path:       fmt.Sprintf("%s/getBackupSecurityPINRequests", id.ID()),
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

// ResourceGuardsGetBackupSecurityPINRequestsObjectsComplete retrieves all the results into a single object
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetBackupSecurityPINRequestsObjectsComplete(ctx context.Context, id ResourceGuardId) (ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteResult, error) {
	return c.ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteMatchingPredicate(ctx context.Context, id ResourceGuardId, predicate ResourceOperationPredicate) (result ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.ResourceGuardsGetBackupSecurityPINRequestsObjects(ctx, id)
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

	result = ResourceGuardsGetBackupSecurityPINRequestsObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
