// Code generated by interfacer; DO NOT EDIT

package hcapi2

import (
	"context"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

// ServerTypeClientBase is an interface generated for "github.com/hetznercloud/hcloud-go/v2/hcloud.ServerTypeClient".
type ServerTypeClientBase interface {
	All(context.Context) ([]*hcloud.ServerType, error)
	AllWithOpts(context.Context, hcloud.ServerTypeListOpts) ([]*hcloud.ServerType, error)
	Get(context.Context, string) (*hcloud.ServerType, *hcloud.Response, error)
	GetByID(context.Context, int64) (*hcloud.ServerType, *hcloud.Response, error)
	GetByName(context.Context, string) (*hcloud.ServerType, *hcloud.Response, error)
	List(context.Context, hcloud.ServerTypeListOpts) ([]*hcloud.ServerType, *hcloud.Response, error)
}
