package repository

import "sync"

type IdDistribute struct {
	lock    sync.Mutex
	postId  int64
	topicId int64
}

var (
	idDistribute     IdDistribute
	idDistributeOnce sync.Once
)

func (d *IdDistribute) generatePostId() int64 {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.postId++
	return d.postId
}

// updatePostId 更新Id占用最大值
func (d *IdDistribute) updatePostId(newId int64) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.postId < newId {
		d.postId = newId
	}
}

func newIdDistributeInstance() *IdDistribute {
	idDistributeOnce.Do(func() {
		idDistribute = IdDistribute{}
	})
	return &idDistribute
}
