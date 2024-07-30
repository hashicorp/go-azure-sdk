package termsofuseagreementacceptance

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

type ListTermsOfUseAgreementAcceptancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementAcceptance
}

type ListTermsOfUseAgreementAcceptancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementAcceptance
}

type ListTermsOfUseAgreementAcceptancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementAcceptancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementAcceptances ...
func (c TermsOfUseAgreementAcceptanceClient) ListTermsOfUseAgreementAcceptances(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId) (result ListTermsOfUseAgreementAcceptancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsOfUseAgreementAcceptancesCustomPager{},
		Path:       fmt.Sprintf("%s/acceptances", id.ID()),
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

// ListTermsOfUseAgreementAcceptancesComplete retrieves all the results into a single object
func (c TermsOfUseAgreementAcceptanceClient) ListTermsOfUseAgreementAcceptancesComplete(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId) (ListTermsOfUseAgreementAcceptancesCompleteResult, error) {
	return c.ListTermsOfUseAgreementAcceptancesCompleteMatchingPredicate(ctx, id, AgreementAcceptanceOperationPredicate{})
}

// ListTermsOfUseAgreementAcceptancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementAcceptanceClient) ListTermsOfUseAgreementAcceptancesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId, predicate AgreementAcceptanceOperationPredicate) (result ListTermsOfUseAgreementAcceptancesCompleteResult, err error) {
	items := make([]stable.AgreementAcceptance, 0)

	resp, err := c.ListTermsOfUseAgreementAcceptances(ctx, id)
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

	result = ListTermsOfUseAgreementAcceptancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
