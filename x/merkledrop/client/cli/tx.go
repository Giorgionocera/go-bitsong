package cli

import (
	"encoding/json"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/merkledrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"strconv"
	"strings"
)

// NewTxCmd returns the transaction commands for the merkledrop module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "merkledrop transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		//GetCmdGenerate(),
		GetCmdCreate(),
		GetCmdClaim(),
		GetCmdWithdraw(),
	)

	return txCmd
}

func GetCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [file-json] [out-list-json] --denom=xxx --start-height=xxx --end-height=xxx",
		Short: "Create a merkledrop from json file",
		Long: `Create a merkledrop from json file
Parameters:
	file-json: input file list
	out-list-json: output list with proofs

Flags:
	denom: the coin denom to distribuite
	start-height: the height when the merkledrop will begin (0 for immediatally)
	end-height: the height when the merkledrop will ends
		`,
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s tx merkledrop create accounts.json out-list.json \
	--denom=ubtsg \
	--start-height=1 \
	--end-height=10

where accounts.json contains
{
	"bitsong1vgpsha4f8grmsqr6krfdxwpcf3x20h0q3ztaj2": "1000000",
	"bitsong1zm6wlhr622yr9d7hh4t70acdfg6c32kcv34duw": "2000000",
	"bitsong1nzxmsks45e55d5edj4mcd08u8dycaxq5eplakw": "3000000"
}

after the computation the out-list.json should be similar to this output
{
  "bitsong1nzxmsks45e55d5edj4mcd08u8dycaxq5eplakw": {
    "index": 2,
    "amount": "3000000",
    "proof": [
      "3346fbddeb1d097311651f5615d3b2528a3893fb79b2ce40b740e6d470296d85"
    ]
  },
  "bitsong1vgpsha4f8grmsqr6krfdxwpcf3x20h0q3ztaj2": {
    "index": 0,
    "amount": "1000000",
    "proof": [
      "a258c32bee9b0bbb7a2d1999ab4698294844e7440aa6dcd067e0d5142fa20522",
      "7f0b92cc8318e4fb0db9052325b474e2eabb80d79e6e1abab92093d3a88fe029"
    ]
  },
  "bitsong1zm6wlhr622yr9d7hh4t70acdfg6c32kcv34duw": {
    "index": 1,
    "amount": "2000000",
    "proof": [
      "7a807e653a5d63556f46fd66a2ac9af6bddaa6864611e6b8da2ccf8389a91345",
      "7f0b92cc8318e4fb0db9052325b474e2eabb80d79e6e1abab92093d3a88fe029"
    ]
  }
}
`,
			version.AppName,
		)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			listBytes, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			var stringList map[string]string
			if err := json.Unmarshal(listBytes, &stringList); err != nil {
				return fmt.Errorf("Could not unmarshal json: %v", err)
			}

			accMap, err := AccountsFromMap(stringList)
			if err != nil {
				return fmt.Errorf("Could not get accounts from map")
			}

			tree, claimInfo, totalAmt, err := CreateDistributionList(accMap)
			if err != nil {
				return fmt.Errorf("Could not create distribution list: %v", err)
			}

			if _, err := createFile(args[1], claimInfo); err != nil {
				return fmt.Errorf("Could not create file: %v", err)
			}

			startHeight, endHeight, denom, err := parseGenerateFlags(cmd.Flags())
			merkleRoot := fmt.Sprintf("%x", tree.Root())

			if denom == "" {
				return fmt.Errorf("denom is required")
			}

			coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("%s%s", totalAmt.String(), denom))
			if err != nil {
				return err
			}

			msg := types.NewMsgCreate(clientCtx.GetFromAddress(), merkleRoot, startHeight, endHeight, coin)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagsCreate())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseGenerateFlags(flags *flag.FlagSet) (int64, int64, string, error) {
	startHeight, err := flags.GetInt64(FlagStartHeight)
	if err != nil {
		return 0, 0, "", err
	}

	endHeight, err := flags.GetInt64(FlagEndHeight)
	if err != nil {
		return 0, 0, "", err
	}

	denom, err := flags.GetString(FlagDenom)
	if err != nil {
		return 0, 0, "", err
	}

	return startHeight, endHeight, denom, nil
}

func GetCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [id] --proofs=\"...\" --amount=xxx --index=xxx",
		Short: "Claim a merkledrop from provided params",
		Long: `Claim a merkledrop from provided params
Parameters:
	id: merkledrop id

Flags:
	proofs: merkle-proofs to claim the merkledrop
	amount: the amount of the merkledrop to claim
	index: the index of the merkledrop to claim
		`,
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s tx merkledrop claim 1 \
	--proofs="a258c32bee9b0bbb7a2d1999ab4698294844e7440aa6dcd067e0d5142fa20522,7f0b92cc8318e4fb0db9052325b474e2eabb80d79e6e1abab92093d3a88fe029" \
	--amount=20000 \
	--index=1
	--from=<key-name>
`,
			version.AppName,
		)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merkledropId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			proofsStr, err := cmd.Flags().GetString(FlagProofs)
			if err != nil {
				return err
			}
			proofs := []string{}
			if proofsStr != "" {
				proofs = strings.Split(proofsStr, ",")
			}

			amount, err := cmd.Flags().GetInt64(FlagAmount)
			if err != nil {
				return err
			}

			index, err := cmd.Flags().GetUint64(FlagIndex)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaim(index, merkledropId, sdk.NewInt(amount), proofs, clientCtx.GetFromAddress())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagClaimMerkledrop())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw",
		Long: "Withdraw funds from an expired merkledrop",
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(`
$ %s tx merkledrop withdraw [id]
`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merkledropId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(clientCtx.GetFromAddress(), merkledropId)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagClaimMerkledrop())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}