package service

import (
	"fmt"
	"github.com/PedroGao/jerry/model"
	"github.com/PedroGao/jerry/utils"
	"sync"
)

func ListUser() ([]*model.UserModel, error) {
	// 1.新建存放用户信息的数组
	infos := make([]*model.UserModel, 0)
	// 2.从数据库查询用户信息
	users := []*model.UserModel{
		{
			Id:       1,
			Username: "pedro",
			SayHello: "world",
			Password: "123456",
		},
		{
			Id:       2,
			Username: "pedro1",
			SayHello: "world6",
			Password: "123456",
		},
		{
			Id:       3,
			Username: "pedro2",
			SayHello: "world0",
			Password: "123456",
		},
	}
	// 3. 新建存放用户id的数组
	var ids []uint64
	//ids := make([]uint64, 0)
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	// 4. 新建goroutine等待队列
	wg := sync.WaitGroup{}
	// 5. 新建用户列表映射
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserModel, len(users)),
	}

	// 6. 新建错误、完成通道
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	// 7. 并发处理数据
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := utils.GenShortId()
			// 如果获取id错误，就写入错误通道
			if err != nil {
				errChan <- err
				return
			}

			//map是并发不安全的，因此需要加锁
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			// 以 key 为 id，value 为 userInfo填充map
			userList.IdMap[u.Id] = &model.UserModel{
				Id:       u.Id,
				Username: u.Username,
				SayHello: fmt.Sprintf("Hello %s", shortId),
				Password: u.Password,
			}
		}(u)
	}

	// 开启一个goroutine等待wg
	go func() {
		wg.Wait()
		//fmt.Println("close finished here!!!")
		//close(finished)
		finished <- true
	}()

	// 阻塞
	select {
	case <-finished:
		//fmt.Println("received finished here!!!")
	case err := <-errChan: // 当通道收到错误，就会立即退出当前函数，并返回错误
		return nil, err
	}

	for _, id := range ids {
		//fmt.Println("received an id , here")
		infos = append(infos, userList.IdMap[id])
	}

	return infos, nil
}
