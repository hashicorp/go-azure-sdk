package profileemail

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

type ListProfileEmailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemEmail
}

type ListProfileEmailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemEmail
}

type ListProfileEmailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileEmailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileEmails ...
func (c ProfileEmailClient) ListProfileEmails(ctx context.Context) (result ListProfileEmailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileEmailsCustomPager{},
		Path:       "/me/profile/emails",
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
		Values *[]beta.ItemEmail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileEmailsComplete retrieves all the results into a single object
func (c ProfileEmailClient) ListProfileEmailsComplete(ctx context.Context) (ListProfileEmailsCompleteResult, error) {
	return c.ListProfileEmailsCompleteMatchingPredicate(ctx, ItemEmailOperationPredicate{})
}

// ListProfileEmailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileEmailClient) ListProfileEmailsCompleteMatchingPredicate(ctx context.Context, predicate ItemEmailOperationPredicate) (result ListProfileEmailsCompleteResult, err error) {
	items := make([]beta.ItemEmail, 0)

	resp, err := c.ListProfileEmails(ctx)
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

	result = ListProfileEmailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
