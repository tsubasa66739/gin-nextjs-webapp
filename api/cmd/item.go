package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
)

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "サンプルコマンド",
	Long:  `ノートを新規作成する。 `,
	Run:   runItem,
}

func init() {
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().StringP("Name", "t", "", "名前")
	itemCmd.Flags().StringP("Price", "b", "", "1000")
}

func runItem(cmd *cobra.Command, args []string) {

	itemRepo := repository.NewItemRepository(db)
	itemSvc := service.NewItemService(itemRepo)

	req := &schema.CreateItemInput{
		Name:  cmd.Flag("Name").Value.String(),
		Price: cmd.Flag("Price").Value.uint(),
	}
	note, err := itemSvc.Create(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("PostNote completed.")
	fmt.Printf("Note: %v\n", note)
}
