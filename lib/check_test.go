package showmackerel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCheckPlugin(t *testing.T) {
	var data [][]string
	file := "./testdata/check"

	data, _ = parseCheckPlugin(data, file)

	assert.Equal(t, data[0][0], "ssh")
	assert.Equal(t, data[0][1], "ruby /path/to/check-ssh.rb")
	assert.Equal(t, data[0][2], "60")
	assert.Equal(t, data[0][3], "1")
	assert.Equal(t, data[0][4], "5")
	assert.Equal(t, data[0][5], "45s")
	assert.Equal(t, data[0][6], "true")
	assert.Equal(t, data[0][7], "HOST=hostname, PORT=port")
	assert.Equal(t, data[0][8], "ruby /path/to/notify_something.rb")
	assert.Equal(t, data[0][9], "0s")
	assert.Equal(t, data[0][10], "NOTIFY_API_KEY=API_KEY")
	assert.Equal(t, data[0][11], "This check monitor is ...")
}
