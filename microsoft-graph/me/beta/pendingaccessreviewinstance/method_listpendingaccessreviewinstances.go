package pendingaccessreviewinstance

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

type ListPendingAccessReviewInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstance
}

type ListPendingAccessReviewInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstance
}

type ListPendingAccessReviewInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstances ...
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstances(ctx context.Context) (result ListPendingAccessReviewInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstancesCustomPager{},
		Path:       "/me/pendingAccessReviewInstances",
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
		Values *[]beta.AccessReviewInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstancesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstancesComplete(ctx context.Context) (ListPendingAccessReviewInstancesCompleteResult, error) {
	return c.ListPendingAccessReviewInstancesCompleteMatchingPredicate(ctx, AccessReviewInstanceOperationPredicate{})
}

// ListPendingAccessReviewInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceClient) ListPendingAccessReviewInstancesCompleteMatchingPredicate(ctx context.Context, predicate AccessReviewInstanceOperationPredicate) (result ListPendingAccessReviewInstancesCompleteResult, err error) {
	items := make([]beta.AccessReviewInstance, 0)

	resp, err := c.ListPendingAccessReviewInstances(ctx)
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

	result = ListPendingAccessReviewInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
