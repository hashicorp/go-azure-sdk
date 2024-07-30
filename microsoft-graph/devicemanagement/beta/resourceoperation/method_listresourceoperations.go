package resourceoperation

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

type ListResourceOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ResourceOperation
}

type ListResourceOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ResourceOperation
}

type ListResourceOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceOperations ...
func (c ResourceOperationClient) ListResourceOperations(ctx context.Context) (result ListResourceOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListResourceOperationsCustomPager{},
		Path:       "/deviceManagement/resourceOperations",
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
		Values *[]beta.ResourceOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourceOperationsComplete retrieves all the results into a single object
func (c ResourceOperationClient) ListResourceOperationsComplete(ctx context.Context) (ListResourceOperationsCompleteResult, error) {
	return c.ListResourceOperationsCompleteMatchingPredicate(ctx, ResourceOperationOperationPredicate{})
}

// ListResourceOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceOperationClient) ListResourceOperationsCompleteMatchingPredicate(ctx context.Context, predicate ResourceOperationOperationPredicate) (result ListResourceOperationsCompleteResult, err error) {
	items := make([]beta.ResourceOperation, 0)

	resp, err := c.ListResourceOperations(ctx)
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

	result = ListResourceOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
