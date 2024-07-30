package people

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

type ListPeoplesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Person
}

type ListPeoplesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Person
}

type ListPeoplesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPeoplesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPeoples ...
func (c PeopleClient) ListPeoples(ctx context.Context, id UserId) (result ListPeoplesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPeoplesCustomPager{},
		Path:       fmt.Sprintf("%s/people", id.ID()),
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
		Values *[]beta.Person `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPeoplesComplete retrieves all the results into a single object
func (c PeopleClient) ListPeoplesComplete(ctx context.Context, id UserId) (ListPeoplesCompleteResult, error) {
	return c.ListPeoplesCompleteMatchingPredicate(ctx, id, PersonOperationPredicate{})
}

// ListPeoplesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PeopleClient) ListPeoplesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate PersonOperationPredicate) (result ListPeoplesCompleteResult, err error) {
	items := make([]beta.Person, 0)

	resp, err := c.ListPeoples(ctx, id)
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

	result = ListPeoplesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
