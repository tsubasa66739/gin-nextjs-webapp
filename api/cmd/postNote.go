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
var postNoteCmd = &cobra.Command{ //cobraライブラリを使用、postNoteCmdコマンドを定義
	Use:   "postNote",
	Short: "ノート新規作成",
	Long: `ノートを新規作成する
ex)
$ ./batchMain postNote --title="タイトル" --body="内容"`, //batchMainでpostNoteを実行し、titleとbodyに入力
	Run: runPostNote,
}

func init() {
	rootCmd.AddCommand(postNoteCmd) //rootCmdのサブコマンドにpostNoteCmdを定義

	postNoteCmd.Flags().StringP("title", "t", "", "タイトル") //titleにフラグ設定
	postNoteCmd.Flags().StringP("body", "b", "", "内容")    //bodyにフラグ設定
}

func runPostNote(cmd *cobra.Command, args []string) { //runPostNote関数をcobraライブラリを使用して定義
	fmt.Println("postNote called.")
	fmt.Printf("Title: %s\n", cmd.Flag("title").Value)
	fmt.Printf("Body: %s\n", cmd.Flag("body").Value)

	noteRepo := repository.NewNoteRepository(db) //レポジトリのNewNoteRepositoryのセットアップを呼び出すnoteRepoを設定
	noteSvc := service.NewNoteService(noteRepo)  //サービスのNewNoteServiceでレポジトリのセットアップを呼び出すnoteSvcを設定

	req := &schema.PostNoteReq{ //PostNoteReqのスキーマ定義をreqに格納
		Title: cmd.Flag("title").Value.String(), //cmdのtitleフラグの値を取得
		Body:  cmd.Flag("body").Value.String(),  //cmdのbodyフラグの値を取得
	}
	note, err := noteSvc.CreateNote(req) //フラグ取得のスキーマを持って、サービスでレポジトリのセットアップを呼び出すnoteSvcのCreateNoteの実行を行う、またはエラー
	if err != nil {                      //errが無でないならば
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("PostNote completed.")
	fmt.Printf("Note: %v\n", note)
}
