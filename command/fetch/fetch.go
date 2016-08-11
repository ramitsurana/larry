package fetch

import (
	"fmt"				
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Used to fetch all the information of your cloud provider",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
var provider string
	str1 := "Name of your Cloud provider:"
    	fmt.Println(str1)
    	fmt.Scanf("%s", &provider)	

var awscomponent string
	str1 := "Name of your AWS service:"
    	fmt.Println(str1)
    	fmt.Scanf("%s", &awscomponent)

var gcloudcomponent string
	str1 := "Name of your GCP service:"
    	fmt.Println(str1)
    	fmt.Scanf("%s", &provider)

	},
}

func init() {
	RootCmd.AddCommand(fetchfetch)
}
