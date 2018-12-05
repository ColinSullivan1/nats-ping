# nats-ping

A simple ping utility for NATS supporting a variety of authentication schemes

## Usage

Connect to NGS using a chain file:
`nats-ping -s connect.ngs.global --chain ~/user.jwtnk`

Connect to NGS using Nkey and JWT:
`nats-ping -s connect.ngs.global --jwt user1.jwt -nk user1.nk`

Connect locally using a username/password
`nats-ping -user colin --pass secret`

Connect locally with no authentication
`nats-ping`

## Output

```text
MacBook-Pro:nats-ping colinsullivan$ ./nats-ping
5126-02-08 14:02:55.13: Connect time: 2.543559ms.
5126-02-08 14:02:55.13: Ping time:    182.664µs.
```
