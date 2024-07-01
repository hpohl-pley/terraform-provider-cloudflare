// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"list_id": schema.StringAttribute{
						Description: "The unique ID of the list.",
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ip", "redirect", "hostname", "asn"),
						},
					},
					"name": schema.StringAttribute{
						Description: "An informative name for the list. Use this name in filter and rule expressions.",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "An informative summary of the list.",
						Optional:    true,
					},
					"id": schema.StringAttribute{
						Description: "The unique ID of the item in the List.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state ListModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
