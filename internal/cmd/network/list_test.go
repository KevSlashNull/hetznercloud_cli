package network_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hetznercloud/cli/internal/cmd/network"
	"github.com/hetznercloud/cli/internal/testutil"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	fx := testutil.NewFixture(t)
	defer fx.Finish()

	cmd := network.ListCmd.CobraCommand(context.Background(), fx.Client, fx.TokenEnsurer)

	fx.ExpectEnsureToken()
	fx.Client.NetworkClient.EXPECT().
		AllWithOpts(
			gomock.Any(),
			hcloud.NetworkListOpts{
				ListOpts: hcloud.ListOpts{
					PerPage:       50,
					LabelSelector: "foo=bar",
				},
				Sort: []string{"id:asc"},
			},
		).
		Return([]*hcloud.Network{
			{
				ID:      123,
				Name:    "test-net",
				IPRange: &net.IPNet{IP: net.ParseIP("192.0.2.1"), Mask: net.CIDRMask(24, 32)},
				Servers: []*hcloud.Server{{ID: 3421}},
				Created: time.Now().Add(-10 * time.Second),
			},
		},
			nil)

	out, err := fx.Run(cmd, []string{"--selector", "foo=bar"})

	expOut := `ID    NAME       IP RANGE       SERVERS    AGE
123   test-net   192.0.2.1/24   1 server   10s
`

	assert.NoError(t, err)
	assert.Equal(t, expOut, out)
}
