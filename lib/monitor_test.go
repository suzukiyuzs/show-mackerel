package showmackerel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMonitor(t *testing.T) {
	var data [][]string
	file := "./testdata/monitor"

	data, _ = parseMonitor(data, file)

	assert.Equal(t, data[0][0], "0123456789a")
	assert.Equal(t, data[0][1], "disk.aa-00.writes.delta")
	assert.Equal(t, data[0][2], "This monitor is ...")
	assert.Equal(t, data[0][3], "host")
	assert.Equal(t, data[0][4], "false")
	assert.Equal(t, data[0][5], "60")
	assert.Equal(t, data[0][6], "disk.aa-00.writes.delta")
	assert.Equal(t, data[0][7], ">")
	assert.Equal(t, data[0][8], "20000")
	assert.Equal(t, data[0][9], "400000")
	assert.Equal(t, data[0][10], "3")
	assert.Equal(t, data[0][11], "3")
	assert.Equal(t, data[0][12], "Hatena-Blog")
	assert.Equal(t, data[0][13], "Hatena-Bookmark: db-master")
}
