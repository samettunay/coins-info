package coins

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	s "strings"

	"time"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"

	"github.com/atomicgo/cursor"
	"github.com/inancgumus/screen"
)

type CryptoCurrency struct {
	url      string
	Price    []string
	Name     []string
	Percent  []string
	UpOrDown []string
}

var (
	crypto    = new(CryptoCurrency)
	ErrLenght = errors.New("Lenght Error! must coin name > 3")
)

func clearInfo() {
	crypto.Name = nil
	crypto.UpOrDown = nil
	crypto.Price = nil
	crypto.Percent = nil
}

func getCoinsInfo() {
	c := colly.NewCollector()

	c.OnHTML("li a div div div span", func(e *colly.HTMLElement) {
		re := regexp.MustCompile("[0-9]")
		if re.MatchString(e.Text) {
			temp := s.TrimSpace(e.Text)
			if s.Contains(e.Text, "%") {
				crypto.Percent = append(crypto.Percent, temp)
			} else if s.Contains(e.Text, "+") || s.Contains(e.Text, "-") {
				if s.Contains(e.Text, "+") {
					crypto.UpOrDown = append(crypto.UpOrDown, color.GreenString(temp))
				} else if s.Contains(e.Text, "-") {
					crypto.UpOrDown = append(crypto.UpOrDown, color.RedString(temp))
				}
			} else {
				crypto.Price = append(crypto.Price, temp)
			}
		}
	})

	c.OnHTML("li a div div div div div", func(e *colly.HTMLElement) {
		re := regexp.MustCompile("[a-z]")
		if re.MatchString(e.Text) && s.Contains(e.Text, "USD") {
			crypto.Name = append(crypto.Name, e.Text)
		}
	})

	c.Visit("https://www.google.com/finance/markets/cryptocurrencies")

}

func createTable(row int) {
	data := [][]string{}

	for i := 0; i < row; i++ {
		data = append(data, []string{strconv.Itoa(i), crypto.Name[i], crypto.Price[i], crypto.UpOrDown[i], crypto.Percent[i]})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Crypto Currency", "Price", "Up & Down", "Percent"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	clearInfo()
}

// Show coins limit 75
func ShowLiveTable(coin int) {
	screen.Clear()
	cursor.Hide()
	for {
		screen.MoveTopLeft()
		time.Sleep(time.Second)
		getCoinsInfo()
		createTable(coin)
	}
}

// Get coin info Info("BTC")
func Info(name string) {

	if len(name) < 3 {
		fmt.Print(ErrLenght)
		return
	}

	getCoinsInfo()

	for i, nm := range crypto.Name {
		con := nm[s.Index(nm, "(")+1 : s.Index(nm, "/")]
		if s.Contains(con, s.ToUpper(name)) {
			fmt.Println(nm, crypto.Price[i], crypto.UpOrDown[i], crypto.Percent[i])
		}
	}

}
