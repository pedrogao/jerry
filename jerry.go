package main

import (
	"errors"
	"github.com/PedroGao/jerry/config"
	lv "github.com/PedroGao/jerry/libs/validator"
	"github.com/PedroGao/jerry/model"
	"github.com/PedroGao/jerry/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	wcfg = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	// parse the flags
	pflag.Parse()

	// init config from file
	if err := config.Init(*wcfg); err != nil {
		log.Fatalf("read config file error : %s  \n", err)
	}

	// init db
	model.Init()
	defer model.Close()

	// set gin app run mode
	gin.SetMode(viper.GetString("runmode"))

	// 更换 v8 至 v9
	binding.Validator = lv.New()

	app := gin.Default()

	// load middleware and routes
	router.Load(app)

	// test api
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "greeting from pedro",
		})
	})

	// ping goroutine for check app is alive or not
	go func() {
		if err := ping(); err != nil {
			log.Fatalf("The router has no response, or it might took too long to start up. erro: %s \n", err)
		}
		log.Infoln("The app has been deployed successfully.")
	}()

	// run
	log.Fatalln(app.Run(viper.GetString("addr")))
}

// check app self when start
func ping() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://localhost" + viper.GetString("addr") + "/")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infoln("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("app is not working")
}
