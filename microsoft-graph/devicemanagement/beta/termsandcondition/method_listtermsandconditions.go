package termsandcondition

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

type ListTermsAndConditionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TermsAndConditions
}

type ListTermsAndConditionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TermsAndConditions
}

type ListTermsAndConditionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsAndConditionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsAndConditions ...
func (c TermsAndConditionClient) ListTermsAndConditions(ctx context.Context) (result ListTermsAndConditionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsAndConditionsCustomPager{},
		Path:       "/deviceManagement/termsAndConditions",
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
		Values *[]beta.TermsAndConditions `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsAndConditionsComplete retrieves all the results into a single object
func (c TermsAndConditionClient) ListTermsAndConditionsComplete(ctx context.Context) (ListTermsAndConditionsCompleteResult, error) {
	return c.ListTermsAndConditionsCompleteMatchingPredicate(ctx, TermsAndConditionsOperationPredicate{})
}

// ListTermsAndConditionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsAndConditionClient) ListTermsAndConditionsCompleteMatchingPredicate(ctx context.Context, predicate TermsAndConditionsOperationPredicate) (result ListTermsAndConditionsCompleteResult, err error) {
	items := make([]beta.TermsAndConditions, 0)

	resp, err := c.ListTermsAndConditions(ctx)
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

	result = ListTermsAndConditionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
