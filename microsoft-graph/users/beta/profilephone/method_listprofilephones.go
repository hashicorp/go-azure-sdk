package profilephone

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

type ListProfilePhonesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemPhone
}

type ListProfilePhonesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemPhone
}

type ListProfilePhonesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfilePhonesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfilePhones ...
func (c ProfilePhoneClient) ListProfilePhones(ctx context.Context, id UserId) (result ListProfilePhonesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfilePhonesCustomPager{},
		Path:       fmt.Sprintf("%s/profile/phones", id.ID()),
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
		Values *[]beta.ItemPhone `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfilePhonesComplete retrieves all the results into a single object
func (c ProfilePhoneClient) ListProfilePhonesComplete(ctx context.Context, id UserId) (ListProfilePhonesCompleteResult, error) {
	return c.ListProfilePhonesCompleteMatchingPredicate(ctx, id, ItemPhoneOperationPredicate{})
}

// ListProfilePhonesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfilePhoneClient) ListProfilePhonesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ItemPhoneOperationPredicate) (result ListProfilePhonesCompleteResult, err error) {
	items := make([]beta.ItemPhone, 0)

	resp, err := c.ListProfilePhones(ctx, id)
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

	result = ListProfilePhonesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
