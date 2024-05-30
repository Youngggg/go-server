package utils

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"apple/common/log"

	"golang.org/x/net/proxy"
)

var Http = &_Http{}

type _Http struct {
}

// 构建一个transport 支持代理 支持随机谷歌CipherSuites
func (_Http) BuildCommonHttpTransportWithProxyWithGoogleCipherSuites(timeOut int, proxyAddr string) *http.Transport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.DisableKeepAlives = true
	transport.IdleConnTimeout = 90 * time.Second
	transport.TLSHandshakeTimeout = 60 * time.Second
	transport.ExpectContinueTimeout = 2 * time.Second
	transport.TLSClientConfig = &tls.Config{
		CipherSuites:       CipherSuites.GetGoogleCipherSuites(),
		InsecureSkipVerify: true,
	}

	header := http.Header{}
	header.Add("Connection", "close")
	transport.ProxyConnectHeader = header

	transport.DialContext = (&net.Dialer{
		Timeout: time.Duration(timeOut) * time.Second,
	}).DialContext

	if proxyAddr != "" {
		//proxyUrl, err := url.Parse("socks5://" + proxyAddr)
		//if err != nil {
		//	utils.TaskLogger.Error("parse proxyAddr err" + err.Error())
		//	return transport
		//}
		//transport.Proxy = http.ProxyURL(proxyUrl)
		dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			log.Error("proxy.SOCKS5 err" + err.Error())
			return transport
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			return dialer.Dial(network, addr)
		}
	}

	return transport
}

func (_Http) BuildCommonHttpTransportWithProxy(timeOut int, proxyAddr string) *http.Transport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.DisableKeepAlives = true
	transport.IdleConnTimeout = 90 * time.Second
	transport.TLSHandshakeTimeout = 30 * time.Second
	transport.ExpectContinueTimeout = 2 * time.Second
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	header := http.Header{}
	header.Add("Connection", "close")
	transport.ProxyConnectHeader = header

	transport.DialContext = (&net.Dialer{
		Timeout: time.Duration(timeOut) * time.Second,
	}).DialContext

	if proxyAddr != "" {
		//proxyUrl, err := url.Parse("socks5://" + proxyAddr)
		//if err != nil {
		//	utils.TaskLogger.Error("parse proxyAddr err" + err.Error())
		//	return transport
		//}
		//transport.Proxy = http.ProxyURL(proxyUrl)
		dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
		if err != nil {
			log.Error("proxy.SOCKS5 err" + err.Error())
			return transport
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			return dialer.Dial(network, addr)
		}
	}

	return transport
}
