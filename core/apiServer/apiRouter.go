package apiServer

import (
	//_ "cloudplatform/core/apiServer/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) InitRoute() {
	s.eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
