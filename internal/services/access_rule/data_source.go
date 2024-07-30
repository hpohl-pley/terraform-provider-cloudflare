// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccessRuleDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccessRuleDataSource{}

func NewAccessRuleDataSource() datasource.DataSource {
	return &AccessRuleDataSource{}
}

func (d *AccessRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_rule"
}

func (r *AccessRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AccessRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccessRuleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := AccessRuleResultDataSourceEnvelope{*data}
		params := firewall.AccessRuleGetParams{}

		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		_, err := r.client.Firewall.AccessRules.Get(
			ctx,
			data.Identifier.ValueString(),
			params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}
		bytes, _ := io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
			return
		}
		data = &env.Result
	} else {
		params := firewall.AccessRuleListParams{
			Direction: cloudflare.F(firewall.AccessRuleListParamsDirection(data.Filter.Direction.ValueString())),
			EgsPagination: cloudflare.F(firewall.AccessRuleListParamsEgsPagination{
				Json: cloudflare.F(firewall.AccessRuleListParamsEgsPaginationJson{
					Page:    cloudflare.F(data.Filter.EgsPagination.Json.Page.ValueFloat64()),
					PerPage: cloudflare.F(data.Filter.EgsPagination.Json.PerPage.ValueFloat64()),
				}),
			}),
			Filters: cloudflare.F(firewall.AccessRuleListParamsFilters{
				ConfigurationTarget: cloudflare.F(firewall.AccessRuleListParamsFiltersConfigurationTarget(data.Filter.Filters.ConfigurationTarget.ValueString())),
				ConfigurationValue:  cloudflare.F(data.Filter.Filters.ConfigurationValue.ValueString()),
				Match:               cloudflare.F(firewall.AccessRuleListParamsFiltersMatch(data.Filter.Filters.Match.ValueString())),
				Mode:                cloudflare.F(firewall.AccessRuleListParamsFiltersMode(data.Filter.Filters.Mode.ValueString())),
				Notes:               cloudflare.F(data.Filter.Filters.Notes.ValueString()),
			}),
			Order:   cloudflare.F(firewall.AccessRuleListParamsOrder(data.Filter.Order.ValueString())),
			Page:    cloudflare.F(data.Filter.Page.ValueFloat64()),
			PerPage: cloudflare.F(data.Filter.PerPage.ValueFloat64()),
		}
		if !data.AccountID.IsNull() {
			params.AccountID = cloudflare.F(data.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
		}

		items := &[]*AccessRuleDataSourceModel{}
		env := AccessRuleResultListDataSourceEnvelope{items}

		page, err := r.client.Firewall.AccessRules.List(ctx, params)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(*items); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		data = (*items)[0]
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
