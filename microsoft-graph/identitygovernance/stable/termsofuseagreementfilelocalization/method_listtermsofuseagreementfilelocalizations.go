package termsofuseagreementfilelocalization

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

type ListTermsOfUseAgreementFileLocalizationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementFileLocalization
}

type ListTermsOfUseAgreementFileLocalizationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementFileLocalization
}

type ListTermsOfUseAgreementFileLocalizationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFileLocalizationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFileLocalizations ...
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizations(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId) (result ListTermsOfUseAgreementFileLocalizationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsOfUseAgreementFileLocalizationsCustomPager{},
		Path:       fmt.Sprintf("%s/file/localizations", id.ID()),
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
		Values *[]stable.AgreementFileLocalization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFileLocalizationsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizationsComplete(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId) (ListTermsOfUseAgreementFileLocalizationsCompleteResult, error) {
	return c.ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate(ctx, id, AgreementFileLocalizationOperationPredicate{})
}

// ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementId, predicate AgreementFileLocalizationOperationPredicate) (result ListTermsOfUseAgreementFileLocalizationsCompleteResult, err error) {
	items := make([]stable.AgreementFileLocalization, 0)

	resp, err := c.ListTermsOfUseAgreementFileLocalizations(ctx, id)
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

	result = ListTermsOfUseAgreementFileLocalizationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
