package profileinterest

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

type ListProfileInterestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonInterest
}

type ListProfileInterestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonInterest
}

type ListProfileInterestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileInterestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileInterests ...
func (c ProfileInterestClient) ListProfileInterests(ctx context.Context) (result ListProfileInterestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileInterestsCustomPager{},
		Path:       "/me/profile/interests",
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
		Values *[]beta.PersonInterest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileInterestsComplete retrieves all the results into a single object
func (c ProfileInterestClient) ListProfileInterestsComplete(ctx context.Context) (ListProfileInterestsCompleteResult, error) {
	return c.ListProfileInterestsCompleteMatchingPredicate(ctx, PersonInterestOperationPredicate{})
}

// ListProfileInterestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileInterestClient) ListProfileInterestsCompleteMatchingPredicate(ctx context.Context, predicate PersonInterestOperationPredicate) (result ListProfileInterestsCompleteResult, err error) {
	items := make([]beta.PersonInterest, 0)

	resp, err := c.ListProfileInterests(ctx)
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

	result = ListProfileInterestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
