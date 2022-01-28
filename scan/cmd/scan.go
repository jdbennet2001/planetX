/*
 Scan a directory for all comics.
Save the output in a file.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// MinFileSize Set minimum file size when looking for comics
const MinFileSize = 32 * 1024

/*
  scanCmd represents the scan command
*/
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Cache directory contents to disk",
	Long: `Scan CBZ catalog, storing names/sizes to disk to enable fast searches.
		   Sample Usage: catGrep scan /Volumes/Seagate/Comics`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("<catGrep scan> requires a directory as a parameter")
			os.Exit(-1)
		}

		dir := args[0]

		// Check that the directory exists
		_, err := os.Stat(dir)
		if err != nil{
			fmt.Println("Directory ", dir, " not reachable")
			os.Exit(-1)
		}

		println( "Scanning...")

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}




