package workers

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/models"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func longRunningTask(sym []models.Symbols, workerId int) <-chan string {
	r := make(chan string)
	go func() {
		defer close(r)
		// Simulate a workload.
		dataBaseSynchronizator(sym)
		r <- "Done Thread: " + strconv.Itoa(workerId)
	}()
	return r
}
func threadStarter(sym []models.Symbols, id int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	dataBaseSynchronizator(sym)

	fmt.Printf("Worker %d done\n", id)
}

func Ticker(numOfThread int) {
	var wg sync.WaitGroup
	threadLoopCounter := 0
	fmt.Println("Start at: ", time.Now().Format("2006-01-02 3:4:5"))
	sym := []models.Symbols{}
	reslut := global.BRC_DB.Find(&sym)
	// SELECT * FROM `users`
	println(reslut.Error) // returns found records count, equals `len(users)`

	size := len(sym) / numOfThread
	end := len(sym)
	//iszero:= end%numOfThread
	//if iszero != 0{
	//	numOfThread = numOfThread + 2
	//}
	//var threadCounter int
	type slice struct {
		left int
		right int
	}
	sliceSize :=[]slice{}
	for{
	var j int
	for i := 0; i < end; i+=size {
		j += size
		if j > end {
			j = end
		}
		sliceSize = append(sliceSize,slice{left: i,right: j})
	}
		wg.Add(len(sliceSize))
		for index :=range sliceSize{
			go threadStarter(sym[sliceSize[index].left:sliceSize[index].right], threadLoopCounter, &wg)
			threadLoopCounter = threadLoopCounter + 1
		}
		wg.Wait()
		fmt.Println("End at: ", time.Now().Format("2006-01-02 3:4:5"))
		time.Sleep(10 * time.Second)
	}
}

//DataBaseSynchronizator DataBase Sync
func dataBaseSynchronizator(sym []models.Symbols) {
	defer handlepanic()
	var brokenSymbol []models.Symbols
	for i := 0; i < len(sym); i++ {
		symData := RealTimeDataDownloaderKey(sym[i])
		if symData != nil {
			symSatus := Analyser(symData)
			symData.SymbolsID = sym[i].ID
			symbolRealTimeTable := global.BRC_DB.Create(symData)
			symSatus.SymbolsID = sym[i].ID
			symbolStatusTable := global.BRC_DB.Create(symSatus)
			//println("org: ",result.Error)
			if symbolRealTimeTable == nil {
				global.BRC_LOG.Error(" dataBaseSynchronizator: "+ sym[i].Name)
			}
			if symbolStatusTable == nil {
				global.BRC_LOG.Error(" dataBaseSynchronizator: "+ sym[i].Name)
			}
		} else {
			brokenSymbol = append(brokenSymbol, sym[i])
		}
	}
}
func getBrokenLink(sym []models.Symbols) {
	for i := 0; i < len(sym); i++ {
		rawData := RealTimeDataDownloaderKey(sym[i])
		result := global.BRC_DB.Create(rawData)
		if result == nil {
			global.BRC_LOG.Error(" dataBaseSynchronizator: "+ sym[i].Name)
		}
	}
}
