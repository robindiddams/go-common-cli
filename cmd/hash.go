package cmd

import (
	"fmt"

	"github.com/pinpt/go-common/hash"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "hash package",
}

var hashValuesCmd = &cobra.Command{
	Use:   "Values",
	Short: "hash.Values([]interface{}...)",
	Long: `Will run hash.Values on the args in order.
Examples:

gc hash Values 1234 abcd # standard hash
gc hash Values robin nil # "nil" will be swapped for go nil
gc hash Values a empty b # "empty" will be swapped for ""
`,
	Run: func(cmd *cobra.Command, args []string) {
		var values []interface{}
		for _, arg := range args {
			var val interface{} = arg
			if arg == "nil" {
				val = nil
			} else if arg == "empty" {
				val = ""
			}
			values = append(values, val)
		}
		fmt.Printf("hashing:")
		for _, val := range values {
			fmt.Printf(" %v", val)
		}
		fmt.Printf("\n")
		fmt.Println(hash.Values(values...))
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.AddCommand(hashValuesCmd)
}
