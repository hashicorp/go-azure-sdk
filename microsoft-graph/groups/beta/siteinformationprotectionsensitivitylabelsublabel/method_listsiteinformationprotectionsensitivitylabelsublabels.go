package siteinformationprotectionsensitivitylabelsublabel

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

type ListSiteInformationProtectionSensitivityLabelSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SensitivityLabel
}

type ListSiteInformationProtectionSensitivityLabelSublabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SensitivityLabel
}

type ListSiteInformationProtectionSensitivityLabelSublabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteInformationProtectionSensitivityLabelSublabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteInformationProtectionSensitivityLabelSublabels ...
func (c SiteInformationProtectionSensitivityLabelSublabelClient) ListSiteInformationProtectionSensitivityLabelSublabels(ctx context.Context, id GroupIdSiteIdInformationProtectionSensitivityLabelId) (result ListSiteInformationProtectionSensitivityLabelSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteInformationProtectionSensitivityLabelSublabelsCustomPager{},
		Path:       fmt.Sprintf("%s/sublabels", id.ID()),
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
		Values *[]beta.SensitivityLabel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteInformationProtectionSensitivityLabelSublabelsComplete retrieves all the results into a single object
func (c SiteInformationProtectionSensitivityLabelSublabelClient) ListSiteInformationProtectionSensitivityLabelSublabelsComplete(ctx context.Context, id GroupIdSiteIdInformationProtectionSensitivityLabelId) (ListSiteInformationProtectionSensitivityLabelSublabelsCompleteResult, error) {
	return c.ListSiteInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx, id, SensitivityLabelOperationPredicate{})
}

// ListSiteInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteInformationProtectionSensitivityLabelSublabelClient) ListSiteInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdInformationProtectionSensitivityLabelId, predicate SensitivityLabelOperationPredicate) (result ListSiteInformationProtectionSensitivityLabelSublabelsCompleteResult, err error) {
	items := make([]beta.SensitivityLabel, 0)

	resp, err := c.ListSiteInformationProtectionSensitivityLabelSublabels(ctx, id)
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

	result = ListSiteInformationProtectionSensitivityLabelSublabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
