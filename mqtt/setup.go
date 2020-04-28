package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"

	mq "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli/v2"
)

type (
	Client = mq.Client

	ClientConfig struct {
		WillTopic               string
		WillPayload             string
		WillQoS                 int
		WillRetain              bool
		ClientID                string
		OnConnectHandler        func(mq.Client)
		OnConnectionLostHandler func(mq.Client, error)
	}
)

func NewPersistentMqtt(config ClientConfig, cliCtx *cli.Context) (mqttClient mq.Client, err error) {
	if config.ClientID == "" {
		config.ClientID = NewUniqueIdentifier()
	}

	opts := mq.NewClientOptions().
		SetCleanSession(false).
		SetAutoReconnect(true).
		SetResumeSubs(true).
		SetClientID(config.ClientID).
		SetConnectTimeout(cliCtx.Duration("mqtt.connection_timeout")).
		SetKeepAlive(cliCtx.Duration("mqtt.keepalive")).
		SetMaxReconnectInterval(cliCtx.Duration("mqtt.max_reconnect_interval")).
		SetMessageChannelDepth(100).
		SetPingTimeout(cliCtx.Duration("mqtt.ping_timeout")).
		SetProtocolVersion(4).
		SetOrderMatters(false).
		SetWriteTimeout(cliCtx.Duration("mqtt.write_timeout"))

	if config.OnConnectHandler != nil {
		opts.SetOnConnectHandler(config.OnConnectHandler)
	}

	if config.OnConnectionLostHandler != nil {
		opts.SetConnectionLostHandler(config.OnConnectionLostHandler)
	}

	if config.WillTopic != "" {
		opts.SetWill(
			config.WillTopic,
			config.WillPayload,
			byte(config.WillQoS&0xff),
			config.WillRetain,
		)
	}

	if cliCtx.IsSet("mqtt.username") && cliCtx.String("mqtt.username") != "" {
		opts.SetUsername(cliCtx.String("mqtt.username"))
	}

	if cliCtx.IsSet("mqtt.password") && cliCtx.String("mqtt.password") != "" {
		opts.SetPassword(cliCtx.String("mqtt.password"))
	}

	if cliCtx.IsSet("mqtt.tls") && cliCtx.Bool("mqtt.tls") {
		tlsCfg, err := setupTLS(cliCtx.String("mqtt.ca"), cliCtx.String("mqtt.cer"), cliCtx.String("mqtt.key"))
		if err != nil {
			return nil, err
		}

		opts.SetTLSConfig(tlsCfg)
		opts.AddBroker(fmt.Sprintf("ssl://%s", cliCtx.String("mqtt.address")))
	} else {
		opts.AddBroker(fmt.Sprintf("tcp://%s", cliCtx.String("mqtt.address")))
	}

	return mq.NewClient(opts), nil
}

// NewUniqueIdentifier returns a unique identifier that the client can use.
// This identifier is what should be set for the lastWillID for anything
// that is bridging more than one device
func NewUniqueIdentifier() string {
	return uuid.NewV4().String()
}

func setupTLS(caPath, certPath, keyPath string) (*tls.Config, error) {
	tlsCfg := &tls.Config{}

	if caPath != "" {
		caPem, err := ioutil.ReadFile(caPath)
		if err != nil {
			return nil, err
		}

		tlsCfg.RootCAs = x509.NewCertPool()
		tlsCfg.RootCAs.AppendCertsFromPEM(caPem)
	}

	if certPath != "" && keyPath != "" {
		if keyPath == "" {
			return nil, errors.New("certificate path specified, but key path missing")
		}

		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}

		tlsCfg.Certificates = []tls.Certificate{cert}
	}

	return tlsCfg, nil
}
