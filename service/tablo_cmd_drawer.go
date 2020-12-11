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
	drawer(tablo)
}
func drawer(tablo models.RealTimeData) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		[]string{"low price", "high price", "close price"},
		[]string{strconv.FormatInt(tablo.DayMinPrice, 10), strconv.FormatInt(tablo.DayMaxPrice, 10), strconv.FormatInt(tablo.ClosePrice, 10)},
		[]string{"safe Kharid :", strconv.FormatInt(tablo.RowOneDemandVol, 10), strconv.FormatInt(tablo.RowOneDemandPrice, 10)},
		[]string{"safe frosh :", strconv.FormatInt(tablo.RowOneSupplyVol, 10), strconv.FormatInt(tablo.RowOneSupplyPrice, 10)},
		[]string{"ctrl + c  for exit and type other symbol name"},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)
	ui.Render(table1)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
