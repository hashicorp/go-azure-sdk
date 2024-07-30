package attributeset

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAttributeSetsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AttributeSet
}

type ListAttributeSetsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AttributeSet
}

type ListAttributeSetsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAttributeSetsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAttributeSets ...
func (c AttributeSetClient) ListAttributeSets(ctx context.Context) (result ListAttributeSetsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAttributeSetsCustomPager{},
		Path:       "/directory/attributeSets",
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
		Values *[]stable.AttributeSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAttributeSetsComplete retrieves all the results into a single object
func (c AttributeSetClient) ListAttributeSetsComplete(ctx context.Context) (ListAttributeSetsCompleteResult, error) {
	return c.ListAttributeSetsCompleteMatchingPredicate(ctx, AttributeSetOperationPredicate{})
}

// ListAttributeSetsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AttributeSetClient) ListAttributeSetsCompleteMatchingPredicate(ctx context.Context, predicate AttributeSetOperationPredicate) (result ListAttributeSetsCompleteResult, err error) {
	items := make([]stable.AttributeSet, 0)

	resp, err := c.ListAttributeSets(ctx)
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

	result = ListAttributeSetsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
