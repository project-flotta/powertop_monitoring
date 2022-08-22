package stats

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

//parse to get tunables data
func ParseTunables(data [][]string) []string {
	var tunData []string
	k := 0
	for _, line := range data {
		if strings.Contains(
			line[0],
			"*  *  *   Software Settings in Need of Tuning   *  *  *",
		) {
			//when
			k = 1
		}
		if strings.Compare(
			line[0],
			"____________________________________________________________________",
		) == 0 {
			k = 0
		}

		if k == 1 {
			tunData = append(
				tunData,
				line[0],
			)
		}

	}
	return tunData
}

//get number of tunables
func GeNumOfTunables(data []string) uint32 {
	length := uint32(len(data))
	//fmt.Printf(
	//	"the length is : %v",
	//	length,
	//)
	return length
}

func TunableLogs(tundata []string) {

	for _, element := range tundata {
		log.Println(element)
	}

}

//reads the csv file
func ReadCSV(path string) ([][]string, error) {
	//opens file
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//closes file
	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			log.Printf(
				"%v\n",
				err,
			)
			return err
		}
		return nil
	}(f)

	//reads file
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = ';'
	csvReader.Comment = '*'

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Printf(
			"%v\n",
			err,
		)
		return nil, err
	}

	//fmt.Println(data)
	return data, nil

}
