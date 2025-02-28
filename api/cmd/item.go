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
	Short: "アイテムコマンド",
	Long:  `アイテムを新規作成する。 `,
	Run:   runItem,
}

func init() {
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().StringP("Name", "n", "デフォルト", "名前")
	itemCmd.Flags().UintP("Price", "p", 1000, "価格")
	itemCmd.Flags().StringP("Description", "d", "デフォルトの説明", "説明")
}

func runItem(cmd *cobra.Command, args []string) {
	fmt.Println("item called.")
	fmt.Printf("Name: %s\n", cmd.Flag("Name").Value)
	fmt.Printf("Price: %s\n", cmd.Flag("Price").Value)
	fmt.Printf("Description: %s\n", cmd.Flag("Description").Value)

	itemRepo := repository.NewItemRepository(db)
	itemSvc := service.NewItemService(itemRepo)

	price, err := cmd.Flags().GetUint("Price")
	if err != nil {
		fmt.Printf("Error getting Price flag: %v\n", err)
		return
	}

	req := &schema.CreateItemInput{
		Name:        cmd.Flag("Name").Value.String(),
		Price:       price,
		Description: cmd.Flag("Description").Value.String(),
	}
	items, err := itemSvc.Create(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Iteme completed.")
	fmt.Printf("Item: %v\n", items)
}
