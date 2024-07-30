package telecomexpensemanagementpartner

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

type ListTelecomExpenseManagementPartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TelecomExpenseManagementPartner
}

type ListTelecomExpenseManagementPartnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TelecomExpenseManagementPartner
}

type ListTelecomExpenseManagementPartnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTelecomExpenseManagementPartnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTelecomExpenseManagementPartners ...
func (c TelecomExpenseManagementPartnerClient) ListTelecomExpenseManagementPartners(ctx context.Context) (result ListTelecomExpenseManagementPartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTelecomExpenseManagementPartnersCustomPager{},
		Path:       "/deviceManagement/telecomExpenseManagementPartners",
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
		Values *[]beta.TelecomExpenseManagementPartner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTelecomExpenseManagementPartnersComplete retrieves all the results into a single object
func (c TelecomExpenseManagementPartnerClient) ListTelecomExpenseManagementPartnersComplete(ctx context.Context) (ListTelecomExpenseManagementPartnersCompleteResult, error) {
	return c.ListTelecomExpenseManagementPartnersCompleteMatchingPredicate(ctx, TelecomExpenseManagementPartnerOperationPredicate{})
}

// ListTelecomExpenseManagementPartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TelecomExpenseManagementPartnerClient) ListTelecomExpenseManagementPartnersCompleteMatchingPredicate(ctx context.Context, predicate TelecomExpenseManagementPartnerOperationPredicate) (result ListTelecomExpenseManagementPartnersCompleteResult, err error) {
	items := make([]beta.TelecomExpenseManagementPartner, 0)

	resp, err := c.ListTelecomExpenseManagementPartners(ctx)
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

	result = ListTelecomExpenseManagementPartnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
