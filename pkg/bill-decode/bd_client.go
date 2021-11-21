package bill_decode

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type BdClient struct {
	url        string
	cache      map[string]string
	baseParams map[string]string
	client     *http.Client
}

var bdClient = NewBdClient()

func NewBdClient() *BdClient {
	p := map[string]string{
		"type":             "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice",
		"aiPortalDemoType": "normal",
	}

	return &BdClient{
		url:        "https://ai.baidu.com/aidemo",
		cache:      make(map[string]string),
		baseParams: p,
		client:     &http.Client{},
	}
}

func (c *BdClient) GetTextByImageURL(u string) (*BillImageBaiduResult, error) {
	params := map[string]string{
		"image_url": u,
	}
	var bdResp BaiduOCRResponse
	err := c.post(c.url, params, &bdResp)
	if err != nil {
		return nil, err
	}
	return bdResp.getResult()
}

func (c *BdClient) GetTextByImageFile(f string) (*BillImageBaiduResult, error) {
	payload := "data:image/png;base64," + getBase64(f)
	params := map[string]string{
		"image": payload,
	}
	var bdResp BaiduOCRResponse
	err := c.post(c.url, params, &bdResp)
	if err != nil {
		return nil, err
	}
	return bdResp.getResult()
}

// image file to base64
func getBase64(f string) string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

func (c *BdClient) post(u string, params map[string]string, ptr interface{}) error {
	var param = url.Values{}
	for k, v := range c.baseParams {
		param.Set(k, v)
	}
	for k, v := range params {
		param.Set(k, v)
	}
	resp, err := c.client.Post(c.url, "application/x-www-form-urlencoded", strings.NewReader(param.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, ptr)
}
