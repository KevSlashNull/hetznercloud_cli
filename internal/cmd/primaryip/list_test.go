package primaryip

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hetznercloud/cli/internal/testutil"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	fx := testutil.NewFixture(t)
	defer fx.Finish()

	cmd := listCmd.CobraCommand(context.Background(), fx.Client, fx.TokenEnsurer)

	fx.ExpectEnsureToken()
	fx.Client.PrimaryIPClient.EXPECT().
		AllWithOpts(
			gomock.Any(),
			hcloud.PrimaryIPListOpts{
				ListOpts: hcloud.ListOpts{
					PerPage:       50,
					LabelSelector: "foo=bar",
				},
				Sort: []string{"id:asc"},
			},
		).
		Return([]*hcloud.PrimaryIP{
			{
				ID:         123,
				Name:       "test-net",
				AutoDelete: true,
				Type:       hcloud.PrimaryIPTypeIPv4,
				IP:         net.ParseIP("127.0.0.1"),
				Created:    time.Now().Add(-10 * time.Second),
			},
		},
			nil)

	out, err := fx.Run(cmd, []string{"--selector", "foo=bar"})

	expOut := `ID    TYPE   NAME       IP          ASSIGNEE   DNS   AUTO DELETE   AGE
123   ipv4   test-net   127.0.0.1   -          -     yes           10s
`

	assert.NoError(t, err)
	assert.Equal(t, expOut, out)
}
