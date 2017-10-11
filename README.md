GELFMatters
===========

GelfMatters is a proxy aimed at listening for GELF messages, to forward them to Slack or Mattermost. Very usefull to combined with Graylog Streams and the GELF output.


## Configuration
Place the configuration file in the working directory.

    server:
      bind: 0.0.0.0  # Binding interface
      port: 12301    # Port to listen
    mattermost:
      # url to connect to mattermost, also work for slack
      url: https://mattermost.host.net/hooks/my_token
      # Payload to send at mattermost/slack
      # "{}" will be replaced by the value of the gelf fields defined below
      payload: "{\"username\": \"graylog\", \"text\": \"field: {}\nanother: {}\"}"
    gelf:  # fields to extract from the gelf
      - _field1
      - _field2

## Service
Exemple of systemd unit file

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

