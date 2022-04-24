package parser

//https://zetcode.com/golang/goquery/
//$ go get github.com/olekukonko/tablewriter  Console'da Tablo için Harika
import (
	"github.com/PuerkitoBio/goquery"
	"github.com/borakasmer/netflix/core"
	"log"
	"net/http"
	"strconv"
	"time"
)

var movieList = make(map[string][]string, 0)

func ParseWeb(categoryID int, rowCount int, isFive bool) map[string][]string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Get("https://www.netflix.com/tr/browse/genre/34399")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			/*FIND ALL
			data := doc.Find(".nm-collections-title-name")
			exchangeList = make(map[string]string, data.Length())
			data.Each(func(i int, s *goquery.Selection) {
				exchangeList[strconv.Itoa(i)] = s.Text()
			})
			return exchangeList
			*/

			section := doc.Find(".nm-collections-row")
			section.Each(func(i int, s *goquery.Selection) {
				if !isFive && (i < 3 && categoryID == core.NetflixCategory.IkinciUc || i < 6 && categoryID == core.NetflixCategory.UcuncuUc ||
					i < 9 && categoryID == core.NetflixCategory.DorduncuUc || i < 12 && categoryID == core.NetflixCategory.BesinciUc) {
					//skip
				} else if isFive && (i < 5 && categoryID == core.NetflixCategory.IkinciBes || i < 10 && categoryID == core.NetflixCategory.UcuncuBes) {
					//skip
				} else {
					if (!isFive && i >= 3*categoryID) || (isFive && i >= 5*categoryID) {
						return
					}
					name := strconv.Itoa(i+1) + "-" + s.Find(".nm-collections-row-name").Last().Text()
					liList := s.Find(".nm-content-horizontal-row ul li")

					var list = make([]string, 0)
					liList.Each(func(i2 int, s2 *goquery.Selection) {
						if i2 >= rowCount {
							return
						}
						var maxlength = len(s2.Text())
						if len(s2.Text()) > 25 {
							maxlength = 25
						}
						list = append(list, strconv.Itoa(i2+1)+"."+s2.Text()[:maxlength])
					})
					movieList[name] = list
				}
			})
			return movieList
		}
	} else {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return make(map[string][]string, 0)
}
