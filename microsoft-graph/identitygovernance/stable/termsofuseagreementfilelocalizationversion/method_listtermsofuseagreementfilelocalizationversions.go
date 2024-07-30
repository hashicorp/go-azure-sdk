package termsofuseagreementfilelocalizationversion

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

type ListTermsOfUseAgreementFileLocalizationVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementFileVersion
}

type ListTermsOfUseAgreementFileLocalizationVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementFileVersion
}

type ListTermsOfUseAgreementFileLocalizationVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFileLocalizationVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFileLocalizationVersions ...
func (c TermsOfUseAgreementFileLocalizationVersionClient) ListTermsOfUseAgreementFileLocalizationVersions(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) (result ListTermsOfUseAgreementFileLocalizationVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTermsOfUseAgreementFileLocalizationVersionsCustomPager{},
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]stable.AgreementFileVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFileLocalizationVersionsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileLocalizationVersionClient) ListTermsOfUseAgreementFileLocalizationVersionsComplete(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId) (ListTermsOfUseAgreementFileLocalizationVersionsCompleteResult, error) {
	return c.ListTermsOfUseAgreementFileLocalizationVersionsCompleteMatchingPredicate(ctx, id, AgreementFileVersionOperationPredicate{})
}

// ListTermsOfUseAgreementFileLocalizationVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileLocalizationVersionClient) ListTermsOfUseAgreementFileLocalizationVersionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId, predicate AgreementFileVersionOperationPredicate) (result ListTermsOfUseAgreementFileLocalizationVersionsCompleteResult, err error) {
	items := make([]stable.AgreementFileVersion, 0)

	resp, err := c.ListTermsOfUseAgreementFileLocalizationVersions(ctx, id)
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

	result = ListTermsOfUseAgreementFileLocalizationVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
