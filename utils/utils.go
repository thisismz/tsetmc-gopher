package utils

import (
	"tsetmc-gopher/global"
	"tsetmc-gopher/models"
	"tsetmc-gopher/tools"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
func JsonToFile(filename string, data []byte) {
	//file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(filename+".json", data, 0644)
}
func UpdateSymbolJson() {
	result := []models.Symbols{}
	defer handlepanic()
	data, err := ioutil.ReadFile("data/MarketWatchInit.txt")
	check(err)
	temp := string(data)
	AllData := strings.Split(temp, ";")
	for i := 1; i < len(AllData); i++ {
		symdata := models.Symbols{}
		spil := strings.Split(AllData[i], ",")
		if _, err := strconv.Atoi(spil[2]); err == nil {
			fmt.Printf("%q looks like a number.\n", spil[2])
		} else {
			symdata.Name = spil[2]
			symdata.FullName = spil[3]
			symdata.Key = spil[0]
			symdata.IsniId = spil[1]
			result = append(result, symdata)
			if result != nil {
				global.BRC_LOG.Error(" dataBaseSynchronizator: "+ symdata.Name)
			}
		}
	}
	//result :=workers.GetAllSymbolRealtimeData()
	jsonString, _ := json.Marshal(result)
	err = ioutil.WriteFile("allsymbol.json", jsonString, os.ModePerm)

}

//UpdataSymbolDbs get data from testmcFilter section
func UpdataSymbolDbs() {
	defer handlepanic()
	data := tools.Responser(global.BRC_CONFIG.Link.TSE_FILTER_ALL_DATA_NAME)
	temp := string(data)
	AllData := strings.Split(temp, ";")
	for i := 1; i < len(AllData); i++ {
		symdata := models.Symbols{}
		spil := strings.Split(AllData[i], ",")
		if _, err := strconv.Atoi(spil[2]); err == nil {
			//	fmt.Printf("%q looks like a number.\n", spil[2])
		} else {
			symdata.Name = spil[2]
			symdata.FullName = spil[3]
			symdata.Key = spil[0]
			symdata.IsniId = spil[1]
			var result []string
			result = strings.Split(spil[2], "")
			sort.Strings(result)
			if _, err := strconv.Atoi(result[1]); err == nil {
				symdata.IsCompany = false
			} else if strings.Contains(spil[1], "IRT") {
				symdata.IsCompany = false
			} else {
				symdata.IsCompany = true
			}
			if global.BRC_DB.Model(&symdata).Where("Name = ?", symdata.Name).Updates(&symdata).RowsAffected == 0 {
				result := global.BRC_DB.Create(&symdata)
				if result != nil {
					global.BRC_LOG.Error(" dataBaseSynchronizator: "+ symdata.Name)
				} else {
					println("row effected symname:%s", symdata.Name)
				}
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func handlepanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

/*
func CreateOrUpdate(db *gorm.DB, model interface{}, where interface{}, update interface{}) (interface{}, error) {
    var result interface{}
    err := db.Model(model).Where(where).First(result).Error
    if err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        } else {
            //insert
            if err = db.Model(model).Create(update).Error; err != nil {
                return nil, err
            }
        }
    }
    //not update some field
    reflect.ValueOf(update).Elem().FieldByName("someField").SetInt(0)
    if err = db.Model(model).Where(where).Updates(update).Error; err != nil {
        return nil, err
    }
    return update, nil
}
*/
