# nats-ping

A simple ping utility for NATS utilizing a variety of authentication schemes

## Usage

Connect to NGS using a chain file:
`nats-ping -s connect.ngs.global --chain ~/user.jwtnk`

Connect to NGS using Nkey and JWT:
`nats-ping -s connect.ngs.global --jwt user1.jwt -nk user1.nk`

Connect locally using a username/password
`nats-ping -user colin --pass secret`

Connect locally with no authentication
`nats-ping`