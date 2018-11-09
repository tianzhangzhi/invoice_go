package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
15_4MbZVAFP4IKAyh77G0RAt7UWiv0z_pQcOvqoLxT5EzjIKI_bsOhQtfmYCKrtZBHP3w77j9sqdog-I2rmuYQl270TI-9Zog6mpW2GBpMPG6Ad963XtGeFds2I5HZeg8Qp5RU32OTs1j9Gos6DDAJgAGATQZ

"https://mp.weixin.qq.com/intp/invoice/usertitlewxa?action=list_auth&invoice_key=Scan_test"



"https://api.weixin.qq.com/card/invoice/scantitle?access_token=AccessToken"
*/
func main() {
	data := `{"scan_text":"https://w.url.cn/s/ANLcn53?cak=qCOVPfzsxvVVuHv7KEG4hg"}`
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/card/invoice/scantitle?access_token=15_WdMDfVxOAeA0JJlYaccA76y3y31S8I_2nDJKWDyXOQ6DVa0Dz04xPd_dYDlHA2_uwJ3PLnWVDuZFPC32vdEC4L3I0kT5PvQ6DfwgtGpi3gCqK8j3C_aG73rdzaPwf3ajpGiCKaB7T3mgjV1DWKLgAIAQEF",
		strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	strings.NewReader("scan_text=https://w.url.cn/s/ANLcn53?cak=qCOVPfzsxvVVuHv7KEG4hg"))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Header)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
