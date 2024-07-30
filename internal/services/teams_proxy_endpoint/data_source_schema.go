// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_proxy_endpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &TeamsProxyEndpointDataSource{}
var _ datasource.DataSourceWithValidateConfig = &TeamsProxyEndpointDataSource{}

func (r TeamsProxyEndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"proxy_endpoint_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *TeamsProxyEndpointDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *TeamsProxyEndpointDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
