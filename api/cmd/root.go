//ひな形をブートストラップcobra-cli initで自動生成されるものに追記

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"gorm.io/gorm"
)

var db *gorm.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gin-nextjs-webapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	config.Setup()
	db = repository.Setup() //dbにレポジトリのセットアップを格納
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute() //rootCmdの実行
	if err != nil {          //errが無でないならば
		fmt.Println("プログラムを終了します")
		os.Exit(1) //プログラムを修了する関数
	}
}
