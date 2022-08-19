package main

import (
	"actor-mq/actor"
	"actor-mq/mq/cli"
	"actor-mq/mq/pb"

	console "github.com/asynkron/goconsole"
)

// func openFile(fileName string) *os.File {
// 	f, err := os.OpenFile(fileName+".csv",
// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return f
// }

func main() {
	const NumberRequests = 5000

	// file := openFile(fmt.Sprintf("responseTime%d", NumberRequests))
	// defer file.Close()

	cli.StartClient("127.0.0.1", "8091")
	usr1 := cli.NewUser()
	actor.PIDMgr.Register("usr1", usr1)
	usr2 := cli.NewUser()
	actor.PIDMgr.Register("usr2", usr2)
	usr3 := cli.NewUser()
	actor.PIDMgr.Register("usr3", usr3)
	cli.BindUsertoRouter(usr2, actor.NewPID("localhost:8090", "encrypt"), "encrypt")
	cli.BindUsertoRouter(usr3, actor.NewPID("localhost:8090", "encrypt"), "encrypt")
	cli.BindUsertoRouter(usr3, actor.NewPID("localhost:8090", "encrypted"), "encrypted")

	// var initialTime time.Time

	for i := 0; i < NumberRequests; i++ {
		// initialTime = time.Now()

		actor.NewPID("localhost:8090", "encrypt").SendMsg(&pb.CommonMsg{
			Sender: usr2,
			Target: actor.NewPID("localhost:8090", "encrypt"),
			Msg:    "Hello",
		})

		// if _, err := file.WriteString(fmt.Sprintf("%d\n", time.Since(initialTime).Nanoseconds())); err != nil {
		// 	log.Println(err)
		// }

	}
	console.ReadLine()
}
