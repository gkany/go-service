package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/judwhite/go-svc/svc"
	"github.com/liyue201/go-logger"
	"github.com/liyue201/golib/cron"
	"go-service/cache"
	"go-service/models"
	"go-service/service"
	"go-service/service/routers"
	"net/http"
	"syscall"
	"time"
)

func main() {
	cfg := flag.String("C", "config.json", "configuration file")
	flag.Parse()
	if err := cache.InitConfig(*cfg); err != nil {
		fmt.Println("init config error:", err.Error())
		return
	}

	app := &program{}
	if err := svc.Run(app, syscall.SIGINT, syscall.SIGTERM); err != nil {
		logger.Println(err)
	}

	logger.Infof("main done.")
}

type program struct {
	models.WaitGroupWrapper
	cronTasks   *cron.Cron
	engine       *gin.Engine
	httpServer  *http.Server
}

func (p *program) Init(env svc.Environment) error {
	logger.Info("[Init] program init start")

	if err := cache.InitLog(cache.Cfg.Log); err != nil {
		fmt.Println("[Init] init log error:", err)
		return err
	}

	p.cronTasks = cron.NewCron(time.Second)
	p.cronTasks.AddTask(&cron.Task{F: service.CronTest, D: time.Second * time.Duration(10)})

	gin.SetMode(gin.ReleaseMode)
	p.engine = gin.Default()
	routers.InitRoute(p.engine)

	logger.Info("[Init] program init done.")
	return nil
}

func (p *program) Start() error {
	logger.Info("[Start] program start")
	p.Wrap(func() {
		addr := fmt.Sprintf("0.0.0.0:%d", cache.Cfg.HttpPort)
		p.httpServer = &http.Server{Addr: addr, Handler: p.engine}
		p.httpServer.ListenAndServe()
	})

	p.Wrap(func() {
		p.cronTasks.Run()
	})

	return nil
}

func (p *program) Stop() error {
	p.cronTasks.Stop()
	if p.httpServer != nil {
		p.httpServer.Shutdown(context.Background())
	}
	p.Wait()

	logger.Info("[Stop] program stopped")
	return nil
}


