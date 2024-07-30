package managedappregistration

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

type ListManagedAppRegistrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedAppRegistration
}

type ListManagedAppRegistrationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedAppRegistration
}

type ListManagedAppRegistrationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedAppRegistrationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedAppRegistrations ...
func (c ManagedAppRegistrationClient) ListManagedAppRegistrations(ctx context.Context, id UserId) (result ListManagedAppRegistrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedAppRegistrationsCustomPager{},
		Path:       fmt.Sprintf("%s/managedAppRegistrations", id.ID()),
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
		Values *[]beta.ManagedAppRegistration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedAppRegistrationsComplete retrieves all the results into a single object
func (c ManagedAppRegistrationClient) ListManagedAppRegistrationsComplete(ctx context.Context, id UserId) (ListManagedAppRegistrationsCompleteResult, error) {
	return c.ListManagedAppRegistrationsCompleteMatchingPredicate(ctx, id, ManagedAppRegistrationOperationPredicate{})
}

// ListManagedAppRegistrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedAppRegistrationClient) ListManagedAppRegistrationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ManagedAppRegistrationOperationPredicate) (result ListManagedAppRegistrationsCompleteResult, err error) {
	items := make([]beta.ManagedAppRegistration, 0)

	resp, err := c.ListManagedAppRegistrations(ctx, id)
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

	result = ListManagedAppRegistrationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
