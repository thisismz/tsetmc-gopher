package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"strconv"
	"strings"
	"tsetmc-gopher/models"
	"tsetmc-gopher/workers"
)

func longRunningTask(text string) <-chan []byte {
	r := make(chan []byte)
	go func() {
		defer close(r)
		// Simulate a workload.
		res := workers.RealTimeDataDownloaderName(text)
		res2B, _ := json.Marshal(res)
		r <- res2B
	}()
	return r
}
func Input(messege string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(messege)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Trim(text, "\r\n")
	text = strings.Replace(text, "\"", "", -1)
	result := <-longRunningTask(text)
	tablo := models.RealTimeData{}
	json.Unmarshal(result, &tablo)
	//println("nemad : " ,text)
	//println("----------------_______---------------")
	drawer(tablo)
	println("ctrl + c  for exit and type other symbol name")
}
func drawer(tablo models.RealTimeData) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	p4 := widgets.NewParagraph()
	p4.Title = "راهنما"
	p4.Text = "press q or C,c for quit"
	p4.SetRect(40, 60, 70, 20)
	p4.BorderStyle.Fg = ui.ColorBlue
	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		[]string{"Tetadeh Foroshdeh", "Hajme Forosh", "Gimat Forosh"},
		[]string{strconv.FormatInt(tablo.RowOneCountOfSeller, 10), strconv.FormatInt(tablo.RowOneSupplyVol, 10), strconv.FormatInt(tablo.RowOneSupplyPrice, 10)},
		[]string{strconv.FormatInt(tablo.RowTwoCountOfSeller, 10), strconv.FormatInt(tablo.RowTwoSupplyVol, 10), strconv.FormatInt(tablo.RowTwoSupplyPrice, 10)},
		[]string{strconv.FormatInt(tablo.RowThreeCountOfSeller, 10), strconv.FormatInt(tablo.RowThreeSupplyVol, 10), strconv.FormatInt(tablo.RowThreeSupplyPrice, 10)},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)

	table2 := widgets.NewTable()
	table2.Rows = [][]string{
		[]string{"Tetadeh Karid", "Hajme Karid", "Gimat Karid"},
		[]string{strconv.FormatInt(tablo.RowOneCountOfBuyer, 10), strconv.FormatInt(tablo.RowOneDemandVol, 10), strconv.FormatInt(tablo.RowOneDemandPrice, 10)},
		[]string{strconv.FormatInt(tablo.RowTwoCountOfBuyer, 10), strconv.FormatInt(tablo.RowTwoDemandVol, 10), strconv.FormatInt(tablo.RowTwoDemandPrice, 10)},
		[]string{strconv.FormatInt(tablo.RowThreeCountOfBuyer, 10), strconv.FormatInt(tablo.RowThreeDemandVol, 10), strconv.FormatInt(tablo.RowThreeDemandPrice, 10)},
	}
	table2.TextStyle = ui.NewStyle(ui.ColorWhite)
	table2.SetRect(0, 10, 60, 20)
	ui.Render(p4, table1, table2)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
