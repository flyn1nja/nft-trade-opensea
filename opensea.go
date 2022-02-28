package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

func GetCollection(offset, count int, colName string) {

	colBody := getColBody(colName)
	assetBody := getAssetsBody(offset, count, colName)

	// col := Collection{}

	// fmt.Println(res)
	// fmt.Println(string(body))

	// njson.Unmarshal(body, col)
	// fmt.Println(col)

	var wg sync.WaitGroup
	var idWg sync.WaitGroup

	var names []string
	var ids []int
	var tokenAddrs []string
	var basePrices []float64
	// var bidPrices []float64
	var floorPrice float64

	wg.Add(6)
	idWg.Add(1)

	go func() {
		ids = getTokenIds(assetBody)
		wg.Done()
		idWg.Done()
	}()
	go func() {
		names = getColNames(assetBody)
		wg.Done()
	}()
	go func() {
		floorPrice = getFloorPrice(colBody)
		wg.Done()
	}()
	go func() {
		basePrices = getColPrices(assetBody)
		wg.Done()
	}()
	go func() {
		tokenAddrs = getTokenAddresses(assetBody)
		wg.Done()
	}()
	go func() {
		idWg.Wait()
		// bidPrices = getColBids(ids)
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("Number of entries : %v\n", len(names))

	if len(names) == 0 {
		println(assetBody)
		return
	}

	fmt.Printf("Collection floor price: %v\n", floorPrice)
	for i := range names {
		fmt.Printf("%s\t", names[i])
		fmt.Printf("id: %d \t\t address: %s ", ids[i], tokenAddrs[i])
		fmt.Printf("Base: %v ETH\t| Curr: \n", basePrices[i]) //, bidPrices[i])
	}
	// if succ.Success {
	// } else {
	// 	fmt.Printf("La collection %s n'a pas été trouvé / la requête n'a pas fonctionné\n", colName)
	// }
}

func GetCollectionNoAsync(offset, count int, colName string) {

	url := fmt.Sprintf(GET_COL_ASSETS_URL, offset, count, colName)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	var names []string
	var ids []int
	var prices []float64

	names = getColNames(body)
	ids = getTokenIds(body)
	prices = getColPrices(body)

	for i := range names {
		fmt.Printf("%s\t", names[i])
		fmt.Printf("id: %d  ", ids[i])
		fmt.Printf("Base: %v ETH\t| Curr: \n", prices[i])
	}
}

func getColBody(colName string) []byte {
	colUrl := GET_COL_URL + colName

	reqCol, _ := http.NewRequest("GET", colUrl, nil)

	res, _ := http.DefaultClient.Do(reqCol)

	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func getAssetsBody(offset, count int, colName string) []byte {
	url := fmt.Sprintf(GET_COL_ASSETS_URL, offset, count, colName)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-API-KEY", API_KEY)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	return body
}

// var rID, _ = regexp.Compile("\"id\":(\\d+),\"token_id\"")
var rID, _ = regexp.Compile("\"id\":(\\d+)")
var rNAME, _ = regexp.Compile("\"name\":\"([^\"]+)\",\"description\"")
var rETH, _ = regexp.Compile("\"total_price\":\"(\\d+)\"")
var rFLOOR, _ = regexp.Compile("\"floor_price\": ?(\\d+\\.\\d+)")
var rADDRS, _ = regexp.Compile("\"payout_address\":\"(0x[0-9a-f]+)|(null)\"")

func getFloorPrice(body []byte) float64 {
	strs := rFLOOR.FindSubmatch(body)
	if len(strs) == 0 {
		return 0.0
	}

	resp, _ := strconv.ParseFloat(string(strs[1]), 64)

	return resp
}

func getColPrices(body []byte) []float64 {
	strs := rETH.FindAllSubmatch(body, -1)
	resp := []float64{}
	if len(strs) == 0 {
		resp = append(resp, -1.0)
		println("No match for Prices")
		return resp
	}

	i2 := 0
	for i := 0; i < len(strs)/2; i++ {
		i2 = i * 2
		// println(string(strs[i2]))
		ethVal, _ := strconv.ParseFloat(string(strs[i2][1]), 64)
		ethVal /= PRICE_MULT
		resp = append(resp, ethVal)
	}
	return resp
}

func getColNames(body []byte) []string {
	strs := rNAME.FindAllSubmatch(body, -1)
	resp := []string{}
	if len(strs) == 0 {
		println("No match for Names")
		return resp
	}

	i2 := 0
	for i := 0; i < len(strs)/2; i++ {
		i2 = i*2 + 1
		name := string(strs[i2][1])
		resp = append(resp, name)
	}
	return resp
}

func getTokenIds(body []byte) []int {
	strs := rID.FindAllSubmatch(body, -1)
	resp := []int{}
	if len(strs) == 0 {
		resp = append(resp, -1)
		println("No match for Ids")
		// println(string(body))
		return resp
	}

	i2 := 0
	for i := 0; i < len(strs)/2; i++ {
		i2 = i * 2
		// println(string(strs[i2]))
		id, _ := strconv.Atoi(string(strs[i2][1]))
		resp = append(resp, id)
	}
	return resp
}

func getTokenAddresses(body []byte) []string {
	strs := rADDRS.FindAllSubmatch(body, -1)
	resp := []string{}
	if len(strs) == 0 {
		println("No match for Addresses")
		return resp
	}

	i2 := 0
	for i := 0; i < len(strs)/2; i++ {
		i2 = i * 2
		// println(string(strs[i2]))
		addr := string(strs[i2][1])
		resp = append(resp, addr)
	}
	return resp
}

func getColBids(ids []int) []float64 {
	// TODO:
	// Aller chercher le bid (http) pour chaque id dans ids
	// greper les valeurs de tous les bids
	// les mettre dans un array
	// retourner le array

	resp := []float64{}

	// strs := rBIDS.FindAllSubmatch(body, -1)
	// resp := []int{}
	// if len(strs) == 0 {
	// 	resp = append(resp, -1)
	// 	println("No match for Ids")
	// 	return resp
	// }

	// i2 := 0
	// for i := 0; i < len(strs)/2; i++ {
	// 	i2 = i * 2
	// 	// println(string(strs[i2]))
	// 	id, _ := strconv.Atoi(string(strs[i2][1]))
	// 	resp = append(resp, id)
	// }

	// //
	// //
	// for i := 0; i < len(ids); i++ {
	// 	resp = append(resp, float64(ids[i]))
	// }
	// //
	// //

	return resp
}

var rUSD, _ = regexp.Compile("\"usd_price\":(\\d+)")

func getUSDPrice(body []byte) float64 {
	str := rUSD.FindSubmatch(body)
	if len(str) == 0 {
		return -1.0
	} else {
		println(string(str[1]))
	}
	usdVal, _ := strconv.ParseFloat(string(str[1]), 64)
	return usdVal / PRICE_MULT
}

// func readScript(fileToRead string) string {
// 	file, err := os.Open(fileToRead)

// 	var res strings.Builder

// 	buf := make([]byte, 8)

// 	for err == nil {
// 		res.Write(buf)
// 		_, err = file.Read(buf)
// 	}

// 	return res.String()
// }

// func StartClient() {

// 	// browser = openbrowser(fmt.Sprintf("https://osu.ppy.sh/oauth/authorize?client_id=%d&redirect_uri=%s&response_type=code", CLIEND_ID, CONN_ADDR))
// 	client = &http.Client{}
// }
