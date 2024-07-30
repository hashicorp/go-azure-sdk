package agreementacceptance

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

type ListAgreementAcceptancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementAcceptance
}

type ListAgreementAcceptancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementAcceptance
}

type ListAgreementAcceptancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAgreementAcceptancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAgreementAcceptances ...
func (c AgreementAcceptanceClient) ListAgreementAcceptances(ctx context.Context) (result ListAgreementAcceptancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAgreementAcceptancesCustomPager{},
		Path:       "/me/agreementAcceptances",
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
		Values *[]stable.AgreementAcceptance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAgreementAcceptancesComplete retrieves all the results into a single object
func (c AgreementAcceptanceClient) ListAgreementAcceptancesComplete(ctx context.Context) (ListAgreementAcceptancesCompleteResult, error) {
	return c.ListAgreementAcceptancesCompleteMatchingPredicate(ctx, AgreementAcceptanceOperationPredicate{})
}

// ListAgreementAcceptancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AgreementAcceptanceClient) ListAgreementAcceptancesCompleteMatchingPredicate(ctx context.Context, predicate AgreementAcceptanceOperationPredicate) (result ListAgreementAcceptancesCompleteResult, err error) {
	items := make([]stable.AgreementAcceptance, 0)

	resp, err := c.ListAgreementAcceptances(ctx)
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

	result = ListAgreementAcceptancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
