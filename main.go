package main

import (
	"bytes"
	"encoding/csv"
 "encoding/json"
 "fmt"
 "os"
 "crypto/sha256"
 // "strconv"
)
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

type Attributes struct{
  TraitType string  `json:"trait_type"`
  Value     string  `json:"value"`
}

type Collection struct{
  Name  string  `json:"name"`
  Id    string  `json:"id"`
  Attributes1    []Attributes1 `json:"attributes1"`
}

type Attributes1 struct{
  Type string  `json:"type"`
  Value     string  `json:"value"`
}

type Data  struct{
  ExampleData string `json:"example_data"`
}


func main() {
	csvFile, err := os.Open("./nft-data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -2

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var nft1 []NftRecord

	for _, each := range csvData {
		// nft.Format = "CHIP-0007"
		// nft.Number = each[0]
		// nft.Filename = each[1]
		// nft.UUID = each[2]
		// nft.Data = "some string"
		//
		//
		if each[1] == "Filename" {
         // skip header line
         continue
      }
		var nftRecord NftRecord
		nftRecord.Format = "CHIP-0007"
		nftRecord.Name = each[1]
		nftRecord.Description = "Electric-type Pokémon with stretchy cheeks"
		nftRecord.MintingTool = "SuperMinter/2.5.2"
		nftRecord.SensitiveContent = false
		nftRecord.SeriesNumber = each[0]
		nftRecord.SeriesTotal = each[2]


		nftRecord.Collection.Name = "Example Pokémon Collection"
		nftRecord.Collection.Id = "e43fcfe6-1d5c-4d6e-82da-5de3aa8b3b57"


		var attributes Attributes
		attributes.TraitType = "Color"
		attributes.Value = "Yellow"
		nftRecord.Attributes = append(nftRecord.Attributes, attributes)
		attributes.TraitType = "Color1"
		attributes.Value = "Yellow1"
		nftRecord.Attributes = append(nftRecord.Attributes, attributes)

		var attributes1 Attributes1
		attributes1.Type = "description"
		attributes1.Value = "Example Pokémon Collection is the best Pokémon collection. Get yours today!"
		nftRecord.Collection.Attributes1 = append(nftRecord.Collection.Attributes1, attributes1)
		attributes1.Type = "descriptionwe"
		attributes1.Value = "Example Pokémon Collection is the best Pokémon collection. Get yours todaywe!"
		nftRecord.Collection.Attributes1 = append(nftRecord.Collection.Attributes1, attributes1)

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(nftRecord)

		// fmt.Println(reqBodyBytes.Bytes()) // this is the []byte
		h := sha256.New()
		h.Write(reqBodyBytes.Bytes())

		bs := h.Sum(nil)

		tt := fmt.Sprintf("%x", bs)
		// fmt.Println("\n")
		// fmt.Println(tt)
nftRecord.Data.ExampleData = tt

		if each[1] == "" {
         // skip header line
         break
      }

		nft1 = append(nft1, nftRecord)

	}

	// Convert to JSON
	jsonData, err := json.Marshal(nft1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
