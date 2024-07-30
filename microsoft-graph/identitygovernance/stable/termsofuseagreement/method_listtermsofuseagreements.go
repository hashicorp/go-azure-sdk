package termsofuseagreement

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

type ListTermsOfUseAgreementsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Agreement
}

type ListTermsOfUseAgreementsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Agreement
}

type ListTermsOfUseAgreementsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreements ...
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreements(ctx context.Context) (result ListTermsOfUseAgreementsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsOfUseAgreementsCustomPager{},
		Path:       "/identityGovernance/termsOfUse/agreements",
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
		Values *[]stable.Agreement `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreementsComplete(ctx context.Context) (ListTermsOfUseAgreementsCompleteResult, error) {
	return c.ListTermsOfUseAgreementsCompleteMatchingPredicate(ctx, AgreementOperationPredicate{})
}

// ListTermsOfUseAgreementsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreementsCompleteMatchingPredicate(ctx context.Context, predicate AgreementOperationPredicate) (result ListTermsOfUseAgreementsCompleteResult, err error) {
	items := make([]stable.Agreement, 0)

	resp, err := c.ListTermsOfUseAgreements(ctx)
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

	result = ListTermsOfUseAgreementsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
