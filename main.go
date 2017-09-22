package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var description = `
kqm [OPTIONS] host:port [host:port]...

KQM is a command line tool to monitor Apache Kafka for lags.
It also comes with an option to send the lag statistics to Statsd.

Option               Description
------               -----------
--statsd-addr        Use this option if you need to send
                     the lag statistics to Statsd.

--statsd-prefix      This option is REQUIRED IF
                     --statsd-addr is specified.

--read-interval      Specify the interval of calculating
                     the lag statistics (in seconds).
                     Default: 60 seconds
`

func parseCommand() (*QMConfig, error) {

	var (
		brokers                  []string
		readInterval             *int
		statsdAddr, statsdPrefix *string
	)

	readInterval = flag.Int("read-interval", 60, "")
	statsdAddr = flag.String("statsd-addr", "127.0.0.1:8125", "")
	statsdPrefix = flag.String("statsd-prefix", "kqm", "")
	flag.Usage = func() {
		fmt.Println(description)
	}
	flag.Parse()

	brokers = flag.Args()
	if len(brokers) == 0 {
		return nil, fmt.Errorf("Please specify brokers")
	}

	cfg := &QMConfig{
		KafkaCfg: KafkaConfig{
			Brokers: brokers,
		},
		StatsdCfg: StatsdConfig{
			Addr:   *statsdAddr,
			Prefix: *statsdPrefix,
		},
		ReadInterval: time.Duration(*readInterval) * time.Second,
	}
	return cfg, nil
}

func main() {
	cfg, err := parseCommand()
	if err != nil {
		fmt.Printf("%s\n%s", err, description)
		os.Exit(1)
	}
	Start(cfg)
}
