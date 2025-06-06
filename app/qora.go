package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	. "qora/conf"
	"strconv"
	"syscall"
	"time"
)

type Qora struct {
	conf *Config
}

func New() *Qora {
	return &Qora{
		conf: NewConfig("./conf/qora_conf.yaml"),
	}
}

func (qora *Qora) Init() {
	// load configure
	err := qora.conf.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %s\n", err)
		os.Exit(1)
	}
}

func (qora *Qora) Start() {
	// apply default Gin service
	router := gin.Default()
	// apply Gin logger & recovery middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// create router group for nova
	qoraService := router.Group("qora/v1")
	{
		qoraService.GET("/test", func(c *gin.Context) { c.String(http.StatusOK, "hello Qora\n") })
	}
	// enable tls settings
	var tlsConfig *tls.Config
	tlsSettings := qora.conf.Configure.TLS
	if tlsSettings.TLSType != "non-tls" {
		var minVersion uint16
		// tls version
		switch tlsSettings.TLSMinVersion {
		case "1.2":
			minVersion = tls.VersionTLS12
		case "1.3":
			minVersion = tls.VersionTLS13
		default:
			minVersion = tls.VersionTLS13
		}
		// one-way tls
		tlsConfig = &tls.Config{
			MinVersion: minVersion,
			CurvePreferences: []tls.CurveID{
				tls.X25519,
				tls.CurveP256,
			},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
		}
		// mutual tls
		if tlsSettings.TLSType == "mutual-tls" {
			// read CA certificate
			caFile := qora.conf.Configure.TLS.CAFile
			caCert, err := os.ReadFile(caFile)
			if err != nil {
				panic(err)
			}
			// create CA certificate pool
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			// specific CA certificate pool
			tlsConfig.ClientCAs = caCertPool
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}
		// start https service
		port := qora.conf.Configure.Port
		certFile := tlsSettings.CertFile
		keyFile := tlsSettings.KeyFile
		server := &http.Server{
			Addr:      ":" + strconv.Itoa(port),
			Handler:   router,
			TLSConfig: tlsConfig,
		}
		// listen and server
		go func() {
			if err := server.ListenAndServeTLS(certFile, keyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}()
		// waiting for close server signal
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// creat timeout context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// stop & clean up resources
		if err := qora.Stop(); err != nil {
			panic(err)
		}
		// graceful Shutting down server
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	} else {
		// start http service
		port := qora.conf.Configure.Port
		server := &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: router,
		}
		// listen and server
		go func() {
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}()
		// waiting for close server signal
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// creat timeout context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// stop & clean up resources
		if err := qora.Stop(); err != nil {
			panic(err)
		}
		// graceful Shutting down server
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}

func (qora *Qora) Stop() error {
	return nil
}
