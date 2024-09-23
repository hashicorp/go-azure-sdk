package informationprotectionsensitivitylabelsublabel

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

type ListInformationProtectionSensitivityLabelSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SensitivityLabel
}

type ListInformationProtectionSensitivityLabelSublabelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SensitivityLabel
}

type ListInformationProtectionSensitivityLabelSublabelsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListInformationProtectionSensitivityLabelSublabelsOperationOptions() ListInformationProtectionSensitivityLabelSublabelsOperationOptions {
	return ListInformationProtectionSensitivityLabelSublabelsOperationOptions{}
}

func (o ListInformationProtectionSensitivityLabelSublabelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInformationProtectionSensitivityLabelSublabelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListInformationProtectionSensitivityLabelSublabelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInformationProtectionSensitivityLabelSublabelsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInformationProtectionSensitivityLabelSublabelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInformationProtectionSensitivityLabelSublabels - Get sublabels from me
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabels(ctx context.Context, id beta.MeInformationProtectionSensitivityLabelId, options ListInformationProtectionSensitivityLabelSublabelsOperationOptions) (result ListInformationProtectionSensitivityLabelSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInformationProtectionSensitivityLabelSublabelsCustomPager{},
		Path:          fmt.Sprintf("%s/sublabels", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListInformationProtectionSensitivityLabelSublabelsComplete retrieves all the results into a single object
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabelsComplete(ctx context.Context, id beta.MeInformationProtectionSensitivityLabelId, options ListInformationProtectionSensitivityLabelSublabelsOperationOptions) (ListInformationProtectionSensitivityLabelSublabelsCompleteResult, error) {
	return c.ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx, id, options, SensitivityLabelOperationPredicate{})
}

// ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InformationProtectionSensitivityLabelSublabelClient) ListInformationProtectionSensitivityLabelSublabelsCompleteMatchingPredicate(ctx context.Context, id beta.MeInformationProtectionSensitivityLabelId, options ListInformationProtectionSensitivityLabelSublabelsOperationOptions, predicate SensitivityLabelOperationPredicate) (result ListInformationProtectionSensitivityLabelSublabelsCompleteResult, err error) {
	items := make([]beta.SensitivityLabel, 0)

	resp, err := c.ListInformationProtectionSensitivityLabelSublabels(ctx, id, options)
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

	result = ListInformationProtectionSensitivityLabelSublabelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
