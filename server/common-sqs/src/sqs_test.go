package src

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	REGION = "ap-northeast-2"
)

func Test_ListForQueue(t *testing.T) {

	config := New(REGION)

	list := config.ListForQueue()

	assert.Equal(t, len(list), 3)

	for _, url := range list {
		fmt.Println(url)
	}
}

func Test_RetrieveQueue(t *testing.T) {
	config := New(REGION)

	q1, q2 := "ex-queue", "nono-queue"

	q1Res, err := config.RetrieveQueue(q1)

	assert.Equal(t, err, nil)
	assert.NotNil(t, q1Res)

	_, err = config.RetrieveQueue(q2)

	assert.Equal(t, err.Error(), fmt.Sprintf("not Exists %s", q2))
}

func Test_RetriveQueueProperties(t *testing.T) {
	config := New(REGION)

	q, fifo_q := "ex-queue", "ex-queue.fifo"

	res1, err := config.RetrieveQueueProperty(q)

	assert.Nil(t, err)
	assert.Equal(t, strings.Contains(res1, q), true)

	res2, err := config.RetrieveQueueProperty(fifo_q)

	assert.Nil(t, err)
	assert.Equal(t, strings.Contains(res2, q), true)
}
