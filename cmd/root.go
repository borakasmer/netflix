/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/borakasmer/netflix/core"
	"github.com/borakasmer/netflix/parser"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netflix",
	Short: "Bu Cli Tool'u ile, Netflix'deki güncel filmler kategori bazlı çekilir.",
	Long: `
"netflix"
"netflix -c 2 -r 6 -f"

*** c Max: 5, '-f' var ise Max: 3 ve r Max:10 değerini alabilir ***
------------------
"-c" Listelecek Katgori sayısı. Default değeri 1 dir. İlk 3 kategori Grubu listelenir
"-c 2" => 3 - 6 arası kategori grubu.
"-c 3 -f" => 10 - 15 arası kategori grubu. (-f => 5'li grup demek)
"-c 5" => 12 - 15 arası kategori grubu. (-f yok => 3'lu grup demek)
"-r" => Listelecek film sayısı. Default değeri 5 dir. İlk 5 film ekrana basılır.
"netflix" komutu => Default olarak "netflix --category 1 --rowcount 5" anlamına gelmektedir.
Netflix üzerindeki film listesi, anlık olarak Parse Edilerek ekrana, kategori bazında basılır. 

**Netflix.com'da bir sorun olması durumunda, bu servis hizmet veremez!!'

Örnek kullanım:
 ."netflix -c 2 -r 6"
 ."netflix -r 7 -f"
 ."netflix --category 3 --rowcount 8 --five"
 ."netflix" => default: 'netflix -c 1 -r 5'
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		isFive, err := cmd.Flags().GetBool("five")

		category, err := cmd.Flags().GetInt("category")
		if err != nil || (!isFive && category > 5) || (isFive && category > 3) || category < 1 {
			category = core.NetflixCategory.IlkUc
		}
		rowcount, err := cmd.Flags().GetInt("rowcount")
		if err != nil || rowcount > 10 {
			rowcount = 5
		}
		getMovieByCategory(category, rowcount, isFive)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.netflix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("five", "f", false, "Kategori sayısı 3 yerine => 5 kabul edilir.")
	rootCmd.Flags().IntP("category", "c", core.NetflixCategory.IlkUc, "Netflix'den çekilecek Kategori Grubu: '-f'(5li grup) var ise 1 ile 3 yok is 1 ile 5 arasında")
	rootCmd.Flags().IntP("rowcount", "r", 5, "Kategori bazlı toplam çekilecek film sayısı. 1 ile 10 arasında")
}

func getMovieByCategory(categoryID int, rowCount int, isFive bool) {
	var moveData = parser.ParseWeb(categoryID, rowCount, isFive)

	categories := make([]string, 0)
	movieList := make([][]string, 0)
	keys := core.SortedKeys(moveData) //range moveData => We will sort categorys by their keys(names)

	for i := 0; i < rowCount; i++ {
		movieLine := make([]string, 0)
		//for category, title := range moveData {
		for _, category := range keys { //We will loop categories by theirs sorted names for every rowCount(moview)
			categories = core.UniqueAppend(categories, category)
			for titleIndex, item := range moveData[category] { //Loop int Category by its sorted key(Name)
				if i == titleIndex {
					movieLine = append(movieLine, item)
				} else if titleIndex > i {
					break
				}
			}
		}
		movieList = append(movieList, movieLine)
	}
	//Create Header of Table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(categories)
	//---------------------
	//Set Table Title
	table.SetCaption(true, "Netflix'de Kategori Bazlı Top "+strconv.Itoa(rowCount)+" Film\n ®coderbora => www.borakasmer.com")
	table.AppendBulk(movieList)
	if !isWindows() {
		if !isFive {
			table.SetHeaderColor(tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgMagentaColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgYellowColor})
		} else {
			table.SetHeaderColor(tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgMagentaColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgYellowColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgBlueColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgRedColor})
		}
	}
	table.Render()
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
