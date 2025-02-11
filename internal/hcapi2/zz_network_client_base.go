// Code generated by interfacer; DO NOT EDIT

package hcapi2

import (
	"context"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

// NetworkClientBase is an interface generated for "github.com/hetznercloud/hcloud-go/v2/hcloud.NetworkClient".
type NetworkClientBase interface {
	AddRoute(context.Context, *hcloud.Network, hcloud.NetworkAddRouteOpts) (*hcloud.Action, *hcloud.Response, error)
	AddSubnet(context.Context, *hcloud.Network, hcloud.NetworkAddSubnetOpts) (*hcloud.Action, *hcloud.Response, error)
	All(context.Context) ([]*hcloud.Network, error)
	AllWithOpts(context.Context, hcloud.NetworkListOpts) ([]*hcloud.Network, error)
	ChangeIPRange(context.Context, *hcloud.Network, hcloud.NetworkChangeIPRangeOpts) (*hcloud.Action, *hcloud.Response, error)
	ChangeProtection(context.Context, *hcloud.Network, hcloud.NetworkChangeProtectionOpts) (*hcloud.Action, *hcloud.Response, error)
	Create(context.Context, hcloud.NetworkCreateOpts) (*hcloud.Network, *hcloud.Response, error)
	Delete(context.Context, *hcloud.Network) (*hcloud.Response, error)
	DeleteRoute(context.Context, *hcloud.Network, hcloud.NetworkDeleteRouteOpts) (*hcloud.Action, *hcloud.Response, error)
	DeleteSubnet(context.Context, *hcloud.Network, hcloud.NetworkDeleteSubnetOpts) (*hcloud.Action, *hcloud.Response, error)
	Get(context.Context, string) (*hcloud.Network, *hcloud.Response, error)
	GetByID(context.Context, int64) (*hcloud.Network, *hcloud.Response, error)
	GetByName(context.Context, string) (*hcloud.Network, *hcloud.Response, error)
	List(context.Context, hcloud.NetworkListOpts) ([]*hcloud.Network, *hcloud.Response, error)
	Update(context.Context, *hcloud.Network, hcloud.NetworkUpdateOpts) (*hcloud.Network, *hcloud.Response, error)
}
