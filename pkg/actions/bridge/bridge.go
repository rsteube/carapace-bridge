package bridge

import (
	"slices"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	"github.com/carapace-sh/carapace-bridge/pkg/env"
)

// Bridges bridges completions as defined in bridges.yaml and CARAPACE_BRIDGE environment variable
func ActionBridges(command ...string) carapace.Action {
	return actionCommand(command...)(func(command ...string) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if bridge, ok := bridges.Config()[command[0]]; ok {
				switch bridge {
				case "argcomplete":
					return ActionArgcomplete(command...)
				case "bash":
					return ActionBash(command...)
				case "carapace":
					return ActionCarapace(command...)
				case "clap":
					return ActionClap(command...)
				case "click":
					return ActionClick(command...)
				case "cobra":
					return ActionCobra(command...)
				case "complete":
					return ActionComplete(command...)
				case "fish":
					return ActionFish(command...)
				case "inshellisense":
					return ActionInshellisense(command...)
				case "kingpin":
					return ActionKingpin(command...)
				case "powershell":
					return ActionPowershell(command...)
				case "urfavecli":
					return ActionUrfavecli(command...)
				case "yargs":
					return ActionYargs(command...)
				case "zsh":
					return ActionZsh(command...)
				default:
					return carapace.ActionMessage("unknown bridge: %v", bridge)
				}
			}

			for _, b := range env.Bridges() {
				switch b {
				case "bash":
					if slices.Contains(bridges.Bash(), command[0]) {
						return ActionBash(command...)
					}
				case "fish":
					if slices.Contains(bridges.Fish(), command[0]) {
						return ActionFish(command...)
					}
				case "inshellisense":
					if slices.Contains(bridges.Inshellisense(), command[0]) {
						return ActionInshellisense(command...)
					}
				case "zsh":
					if slices.Contains(bridges.Zsh(), command[0]) {
						return ActionZsh(command...)
					}
				}
			}
			return carapace.ActionValues()
		})
	})
}
