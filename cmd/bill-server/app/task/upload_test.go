package task

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	u = "http://r3sdtsxu6.hn-bkt.clouddn.com/7dbc406c9317c4558d9166d3f5785103.png"
)

func TestUpload(t *testing.T) {
	localFile := "../../../../resource/fapiao.png"
	url, err := Upload(localFile)
	log.Info(gctx.New(), url)
	assert.NoError(t, err)
}

func TestDeleteAfterDay(t *testing.T) {
	err := DeleteAfterDay(u, 7)
	assert.NoError(t, err)
}
