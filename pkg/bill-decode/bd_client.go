package bill_decode

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hades300/bill-center/cmd/bill-server/library/flow"
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
	fc         *flow.Leaky
}

var DefaultClient = NewBdClient(2, 3)

func NewBdClient(rate int64, gap int64) *BdClient {
	p := map[string]string{
		"type":             "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice",
		"aiPortalDemoType": "normal",
	}

	return &BdClient{
		url:        "https://ai.baidu.com/aidemo",
		cache:      make(map[string]string),
		baseParams: p,
		client:     &http.Client{},
		fc:         flow.NewLeaky(rate, gap),
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
	_ = c.fc.Wait(1) // flow control
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
