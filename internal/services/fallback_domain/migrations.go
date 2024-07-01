// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r FallbackDomainResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"policy_id": schema.StringAttribute{
						Description: "Device ID.",
						Required:    true,
					},
					"suffix": schema.StringAttribute{
						Description: "The domain suffix to match when resolving locally.",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "A description of the fallback domain, displayed in the client UI.",
						Optional:    true,
					},
					"dns_server": schema.ListAttribute{
						Description: "A list of IP addresses to handle domain resolution.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state FallbackDomainModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
