/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
)

// postNoteCmd represents the postNote command
var postNoteCmd = &cobra.Command{
	Use:   "postNote",
	Short: "ノート新規作成",
	Long: `ノートを新規作成する
ex)
$ ./batchMain postNote --title="タイトル" --body="内容"`,
	Run: runPostNote,
}

func init() {
	rootCmd.AddCommand(postNoteCmd)

	postNoteCmd.Flags().StringP("title", "t", "", "タイトル")
	postNoteCmd.Flags().StringP("body", "b", "", "内容")
}

func runPostNote(cmd *cobra.Command, args []string) {
	fmt.Println("postNote called.")
	fmt.Printf("Title: %s\n", cmd.Flag("title").Value)
	fmt.Printf("Body: %s\n", cmd.Flag("body").Value)

	noteRepo := repository.NewNoteRepository(db)
	noteSvc := service.NewNoteService(noteRepo)

	req := &schema.PostNoteReq{
		Title: cmd.Flag("title").Value.String(),
		Body:  cmd.Flag("body").Value.String(),
	}
	note, err := noteSvc.CreateNote(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("PostNote completed.")
	fmt.Printf("Note: %v\n", note)
}
