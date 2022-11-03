package main

// Importation of required files
import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"crypto/sha256"
	"log"
)


//Initialization of NftRecord variables
type NftRecord struct{
  Format        string  `json:"format"`
  Name          string  `json:"name"`
  Description   string  `json:"description"`
  MintingTool   string  `json:"minting_tool"`
  SensitiveContent bool `json:"sensitive_content"`
  SeriesNumber    string  `json:"series_number"`
  SeriesTotal     string `json:"series_total"`
  Attributes      []Attributes `json:"attributes"`
  Collection  Collection  `json:"collection"`
  Data    Data  `json:"data"`
}


// Initialization of attributes variables
type Attributes struct{
  TraitType string  `json:"trait_type"`
  Value     string  `json:"value"`
}


// initialization of collection variables
type Collection struct{
  Name  string  `json:"name"`
  Id    string  `json:"id"`
  Attributes1    []Attributes1 `json:"attributes1"`
}

// initialization of attributes of variables
type Attributes1 struct{
  Type string  `json:"type"`
  Value     string  `json:"value"`
}


// initialization of Data variables
type Data  struct{
  ExampleData string `json:"example_data"`
}


// where the main code is written
func main() {

	// put the file location in this variable
	workFile := "./nft-data.csv"

		// here we open the csv file containing the nfts
		csvFile, err := os.Open(workFile)
		if err != nil {
		fmt.Println(err)
		}
		defer csvFile.Close()

		// here we read the csv file containing the nfts
		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -2
		csvData, err := reader.ReadAll()
		if err != nil {
		fmt.Println(err)
		os.Exit(1)
		}

// we use this if we want to create one file
	// var nft1 []NftRecord

// here we iterate to create the csv files
	for _, each := range csvData {

// here we skip the header to read the rest of the files
		if each[1] == "Filename" {
         // skip header line
         continue
      }


	// here we assign values to the nftRecords	 struct
		var nftRecord NftRecord
		nftRecord.Format = "CHIP-0007"
		nftRecord.Name = each[1]
		nftRecord.Description = "Electric-type Pokémon with stretchy cheeks"
		nftRecord.MintingTool = "SuperMinter/2.5.2"
		nftRecord.SensitiveContent = false
		nftRecord.SeriesNumber = each[0]
		nftRecord.SeriesTotal = each[2]


//here we assign value to Collection record struct
		nftRecord.Collection.Name = "Example Pokémon Collection"
		nftRecord.Collection.Id = "e43fcfe6-1d5c-4d6e-82da-5de3aa8b3b57"

// here we append values to the attributes struct
		var attributes Attributes
		attributes.TraitType = "Color"
		attributes.Value = "Yellow"
		nftRecord.Attributes = append(nftRecord.Attributes, attributes)
		attributes.TraitType = "Color1"
		attributes.Value = "Yellow1"
		nftRecord.Attributes = append(nftRecord.Attributes, attributes)

// here we append values to the 2nd attributes struct
		var attributes1 Attributes1
		attributes1.Type = "description"
		attributes1.Value = "Example Pokémon Collection is the best Pokémon collection. Get yours today!"
		nftRecord.Collection.Attributes1 = append(nftRecord.Collection.Attributes1, attributes1)
		attributes1.Type = "descriptionwe"
		attributes1.Value = "Example Pokémon Collection is the best Pokémon collection. Get yours todaywe!"
		nftRecord.Collection.Attributes1 = append(nftRecord.Collection.Attributes1, attributes1)

//Here we read the read the whole json for  hashing
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(nftRecord)

// here we hash using hash256
		// fmt.Println(reqBodyBytes.Bytes()) // this is the []byte
		h := sha256.New()
		h.Write(reqBodyBytes.Bytes())
		bs := h.Sum(nil)
		tt := fmt.Sprintf("%x", bs)
		// fmt.Println("\n")
		// fmt.Println(tt)
		nftRecord.Data.ExampleData = tt


		 // here we convert the whole script to json
		jsonData, err := json.Marshal(nftRecord)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

// here i print the json files on my console
		fmt.Println(string(jsonData))

// here I create the output folder
		path := "output"
		err = os.Mkdir(path, os.ModePerm)
				if err != nil {
					log.Println(err)
				}

// here I create the json files needed in the output folder
		jsonFile, err := os.Create("./output/"+each[1]+".json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

// here I write the json data for the created files
		jsonFile.Write(jsonData)
		jsonFile.Close()

// here i break the code if the iteration meets an empty space
		if each[1] == "" {
         // skip header line
         break
      }

// if you want to create one file and append all
		// nft1 = append(nft1, nftRecord)

	}

}
