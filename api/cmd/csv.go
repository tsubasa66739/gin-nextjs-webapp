package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
)

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "csv出力",
	Long:  `csv出力をする`,
	Run:   runCsv,
}

func init() {
	rootCmd.AddCommand(csvCmd)
}

func runCsv(cmd *cobra.Command, args []string) {

	itemRepo := repository.NewItemRepository(db)
	itemSvc := service.NewItemService(itemRepo)

	err := itemSvc.ExportItemToCSV("item.csv")
	if err != nil {
		log.Fatalf("Error exporting item to CSV: %v", err)
	}
	log.Println("CSVファイルに書き出しました")
}
