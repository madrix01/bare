package cmd

import (
	"bare/parser"
	"bare/styles"
	"bare/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd);
}
var addCmd = &cobra.Command{
	Use: "add",
	Short: "add you current bare to repo",
	Long: "same as short for addCmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("Bare Add"))
		addBare();
	},
}

func addBare(){
	recipePath := "./recipe.json"
	parser.Parser(recipePath)
	barePath := parser.BareObj.BarePath
	for _, objPath := range parser.BareObj.Include {
		fmt.Print(styles.AddFileStlyle.Render(objPath))
		sourcePath := "./" + objPath
		destiPath := barePath + "/" + objPath
		err := utils.CopyFileDirectory(sourcePath, destiPath)
		if err != nil {
			fmt.Print(" > ", styles.InitError.Render("Error"))
			fmt.Println("")
		}else{
			fmt.Print(" > ", styles.InitSuccess.Render("Success"))
			fmt.Println("")
		}
	}	

}
