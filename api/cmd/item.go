package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "サンプルコマンド",
	Long:  `サンプルのコマンドです`,
	Run:   runSample,
}

func init() {
	rootCmd.AddCommand(itemCmd) //rootCmdのサブコマンドにitemCmdを追加

	itemCmd.Flags().StringP("Name", "f", "", "ふが") //itemCmdにFlagを設定
}

func runSample(cmd *cobra.Command, args []string) { //cobraライブラリの構造を使用
	fmt.Println("item called.", "TEST TEST")         //文字列毎にスペース、改行あり
	fmt.Printf("Name: %s\n", cmd.Flag("Name").Value) //引数ごとにスペース、改行無し
}
