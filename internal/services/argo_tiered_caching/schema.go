// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ArgoTieredCachingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The identifier of the caching setting",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.StringAttribute{
				Description: "Enables Tiered Caching.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
			},
			"editable": schema.BoolAttribute{
				Description: "Whether the setting is editable",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The time when the setting was last modified",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}
