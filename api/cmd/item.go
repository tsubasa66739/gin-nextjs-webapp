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
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().StringP("Name", "f", "", "ふが")
}

// func add(a int, b int) int {
// 	return a + b
// }

func runSample(cmd *cobra.Command, args []string) {
	fmt.Println("item called.", "TEST TEST")
	fmt.Printf("Name: %s\n", cmd.Flag("Name").Value)
	// println("sample called.", "TEST TEST")
	// fmt.Println("Hello,World!")

	// num := 123
	// var num2 = 123456789
	// num3 := 1.23
	// var num4 float64 = 1.23456789

	// fmt.Println(reflect.TypeOf(num))
	// fmt.Println(reflect.TypeOf(num2))
	// fmt.Println(reflect.TypeOf(num3))
	// fmt.Println(reflect.TypeOf(num4))

	// var string_a string = "Hello World!"
	// fmt.Println(string_a)

	// a := 10 for文などで短縮記法を使う
	// b := 1

	// var num_bool bool = a > b
	// fmt.Println(num_bool)

	// a := [...]string{"sato", "suzuki", "takahashi"}
	// fmt.Println(a[0])

	// b := [2][2]string{{"sato", "suzuki"}, {"takahashi", "taro"}}
	// fmt.Println(b[1][0])

	// var x int = 10
	// var y int = 20
	// if x < y {
	// 	fmt.Println("yの方が大きい")
	// } else if x == y {
	// 	fmt.Println("xとyは同じ")
	// }

	// var arr [3]int = [3]int{1, 2, 3}
	// fmt.Println(arr)

	// var slice = []int{10, 20, 30, 40}
	// fmt.Println(slice)

	// slice = append(slice, 50)
	// fmt.Println(slice)

}
