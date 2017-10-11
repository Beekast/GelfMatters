GELFMatters
===========

GelfMatters is a proxy aimed at listening for GELF messages, to forward them to Slack or Mattermost. Very usefull combined with Graylog Streams and the GELF output.


## Configuration
Place a YAML configuration file named *gelfmatters.conf* in the working directory.

    server:
      bind: 0.0.0.0
      port: 12301
    mattermost:
      url: https://mattermost.host.net/hooks/my_token
      payload: "{\"username\": \"graylog\", \"text\": \"field: {}\nanother: {}\"}"
    gelf:
      - _field1
      - _field2

- payload: is the text to send to Mattermost/Slack. Can be parametrized with '{}'
- gelf: is an ordered list of fields to extract from GELF messages, each '{}' in the payload will be replaced by the corresponding value

## Service
Exemple of a systemd unit file

    [Unit]
    Description=Gelf connector for Mattermost
    
    [Service]
    ExecStart=/usr/local/bin/gelfmatters
    WorkingDirectory=/etc/gelfmatters
    Restart=always
    RestartSec=5s
    KillMode=process
    
    [Install]
    WantedBy=multi-user.target

## Test
Can easily be tested with netcat

    cat test.in | nc 127.0.0.1 12301

## Limitations
- For now, works only with GELF TCP

