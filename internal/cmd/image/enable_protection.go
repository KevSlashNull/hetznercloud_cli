package image

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/cmd/cmpl"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/cli/internal/state"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/spf13/cobra"
)

var EnableProtectionCommand = base.Cmd{
	BaseCobraCommand: func(client hcapi2.Client) *cobra.Command {

		return &cobra.Command{
			Use:   "enable-protection [FLAGS] IMAGE PROTECTIONLEVEL [PROTECTIONLEVEL...]",
			Short: "Enable resource protection for an image",
			Args:  cobra.MinimumNArgs(2),
			ValidArgsFunction: cmpl.SuggestArgs(
				cmpl.SuggestCandidatesF(client.Image().Names),
				cmpl.SuggestCandidates("delete"),
			),
			TraverseChildren:      true,
			DisableFlagsInUseLine: true,
		}
	},
	Run: func(ctx context.Context, client hcapi2.Client, waiter state.ActionWaiter, command *cobra.Command, args []string) error {
		imageID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return errors.New("invalid image ID")
		}
		image := &hcloud.Image{ID: imageID}

		var unknown []string
		opts := hcloud.ImageChangeProtectionOpts{}
		for _, arg := range args[1:] {
			switch strings.ToLower(arg) {
			case "delete":
				opts.Delete = hcloud.Bool(true)
			default:
				unknown = append(unknown, arg)
			}
		}
		if len(unknown) > 0 {
			return fmt.Errorf("unknown protection level: %s", strings.Join(unknown, ", "))
		}

		action, _, err := client.Image().ChangeProtection(ctx, image, opts)
		if err != nil {
			return err
		}

		if err := waiter.ActionProgress(ctx, action); err != nil {
			return err
		}

		fmt.Printf("Resource protection enabled for image %d\n", image.ID)
		return nil
	},
}
