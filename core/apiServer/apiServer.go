package apiServer

import (
	"cloudplatform/config"
	"cloudplatform/log"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Server struct {
	config    *config.Server
	eng       *gin.Engine
	httpServe *http.Server
	ctx       context.Context
	closeOnce sync.Once
}

func New(config *config.Server) *Server {

	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	//最大文件
	engine.MaxMultipartMemory = 1 << 30
	return &Server{
		config: config,
		eng:    engine,
	}
}

func (s *Server) Init() error {
	s.InitRoute()

	return nil
}

func (s *Server) Name() string {
	return "APIServer"
}

func (s *Server) Startup(ctx context.Context) (err error) {
	s.ctx = ctx

	log.Infof("APIServer startup,Listen Address:%s...\r\n", s.config.Address)

	if !s.config.EnableTls {
		s.httpServe = &http.Server{
			Addr:    s.config.Address,
			Handler: s.eng,
		}
		//注册结构体绑定规则
		//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//	err := v.RegisterValidation("TID", req.TID)
		//	if err != nil {
		//		return err
		//	}
		//}
		//将路由API routes存放全局,让logger中间件接口使用

		if err := s.httpServe.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err)
		}
	} else {
		//if err = s.eng.RunTLS(s.config.ListenAddress,
		//	s.config.CertFile,
		//	s.config.KeyFile); err != nil {
		//
		//	log.Error(err)
		//	return err
		//}
	}
	return nil
}

func (s *Server) Close() error {
	s.closeOnce.Do(func() {
		log.Info("APIServer close..")
		if err := s.httpServe.Shutdown(s.ctx); err != nil {
			log.Error(err)
		}
	})
	return nil
}
