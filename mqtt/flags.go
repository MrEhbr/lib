package mqtt

import (
	"time"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var Flags = []cli.Flag{
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.address",
			Aliases: []string{"mqtt_address"},
			Value:   "localhost:1883",
			EnvVars: []string{"MQTT_ADDRESS"},
			Usage:   "Address to MQTT endpoint",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.username",
			Aliases: []string{"mqtt_username"},
			EnvVars: []string{"MQTT_USERNAME"},
			Usage:   "MQTT Username",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.password",
			Aliases: []string{"mqtt_password"},
			EnvVars: []string{"MQTT_PASSWORD"},
			Usage:   "MQTT Password",
		},
	),
	altsrc.NewBoolFlag(
		&cli.BoolFlag{
			Name:    "mqtt.tls",
			Aliases: []string{"mqtt_tls"},
			EnvVars: []string{"MQTT_TLS"},
			Usage:   "Enable TLS",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.ca",
			Aliases: []string{"mqtt_ca"},
			EnvVars: []string{"MQTT_CA"},
			Usage:   "Path to CA certificate",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.cert",
			Aliases: []string{"mqtt_cert"},
			EnvVars: []string{"MQTT_CERT"},
			Usage:   "Path to Client certificate",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "mqtt.key",
			Aliases: []string{"mqtt_key"},
			EnvVars: []string{"MQTT_KEY"},
			Usage:   "Path to Client certificate key",
		},
	),
	altsrc.NewDurationFlag(
		&cli.DurationFlag{
			Name:    "mqtt.connection_timeout",
			Value:   10 * time.Second,
			Aliases: []string{"mqtt_connection-timeout"},
			EnvVars: []string{"MQTT_CONNECTION_TIMEOUT"},
			Usage:   "Connection timeout",
		},
	),
	altsrc.NewDurationFlag(
		&cli.DurationFlag{
			Name:    "mqtt.keepalive",
			Value:   30 * time.Second,
			Aliases: []string{"mqtt_keepalive"},
			EnvVars: []string{"MQTT_KEEPALIVE"},
			Usage:   "Time between each PING packet",
		},
	),
	altsrc.NewDurationFlag(
		&cli.DurationFlag{
			Name:    "mqtt.max_reconnect_interval",
			Value:   2 * time.Minute,
			Aliases: []string{"mqtt_max_reconnect_interval"},
			EnvVars: []string{"MQTT_MAX_RECONNECT_INTERVAL"},
			Usage:   "Maximum time to wait between reconnect attempts",
		},
	),
	altsrc.NewDurationFlag(
		&cli.DurationFlag{
			Name:    "mqtt.ping_timeout",
			Value:   10 * time.Second,
			Aliases: []string{"mqtt_ping_timeout"},
			EnvVars: []string{"MQTT_PING_TIMEOUT"},
			Usage:   "Time after which a ping times out",
		},
	),
	altsrc.NewDurationFlag(
		&cli.DurationFlag{
			Name:    "mqtt.write_timeout",
			Value:   5 * time.Second,
			Aliases: []string{"mqtt_write_timeout"},
			EnvVars: []string{"MQTT_WRITE_TIMEOUT"},
			Usage:   "Time after which a write will time out",
		},
	),
}
