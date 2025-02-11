package server

import (
	"context"
	"fmt"

	"github.com/hetznercloud/cli/internal/cmd/base"
	"github.com/hetznercloud/cli/internal/cmd/cmpl"
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/cli/internal/state"
	"github.com/spf13/cobra"
)

var DetachISOCommand = base.Cmd{
	BaseCobraCommand: func(client hcapi2.Client) *cobra.Command {
		return &cobra.Command{
			Use:                   "detach-iso [FLAGS] SERVER",
			Short:                 "Detach an ISO from a server",
			Args:                  cobra.ExactArgs(1),
			ValidArgsFunction:     cmpl.SuggestArgs(cmpl.SuggestCandidatesF(client.Server().Names)),
			TraverseChildren:      true,
			DisableFlagsInUseLine: true,
		}
	},
	Run: func(ctx context.Context, client hcapi2.Client, waiter state.ActionWaiter, cmd *cobra.Command, args []string) error {
		idOrName := args[0]
		server, _, err := client.Server().Get(ctx, idOrName)
		if err != nil {
			return err
		}
		if server == nil {
			return fmt.Errorf("server not found: %s", idOrName)
		}

		action, _, err := client.Server().DetachISO(ctx, server)
		if err != nil {
			return err
		}

		if err := waiter.ActionProgress(ctx, action); err != nil {
			return err
		}

		fmt.Printf("ISO detached from server %d\n", server.ID)
		return nil
	},
}
