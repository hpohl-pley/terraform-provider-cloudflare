// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ipsec_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IPSECTunnelResultDataSourceEnvelope struct {
	Result IPSECTunnelDataSourceModel `json:"result,computed"`
}

type IPSECTunnelDataSourceModel struct {
	AccountID     types.String         `tfsdk:"account_id" path:"account_id"`
	IPSECTunnelID types.String         `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id"`
	IPSECTunnel   jsontypes.Normalized `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel"`
}
