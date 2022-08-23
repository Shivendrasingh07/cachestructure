package serverProvider

import (
	"example.com/m/provider"
	cacheprovider "example.com/m/provider/cache"
)

type Server struct {
	Cache provider.CacheProviderInterface
}

func SrvInit() *Server {
	cache := cacheprovider.NewCacheProvider()

	return &Server{
		Cache: cache,
	}
}

//func (srv *Server) Start() {
//	addr := ":4444"
//	httpSrv := &http.Server{
//		Addr:    addr,
//		Handler: srv.SetupRoutes(),
//	}
//
//	logrus.Info("Server running at PORT ", addr)
//	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//		logrus.Fatalf("SubscribeAllPartitions %v", err)
//		return
//	}
//}

//func (srv *Server) Stop() {
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	logrus.Info("closing server...")
//	//_ = srv.httpServer.Shutdown(ctx)
//	logrus.Info("Done")
//}
