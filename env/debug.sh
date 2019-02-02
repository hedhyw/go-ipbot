# One of panic fatal error warning info debug trace
export IPBOT_LOGLEVEL="debug"

# HTTP Proxy for connections
export IPBOT_HTTP_PROXY="socks5://HOST:PORT"

# Format is user1:user2:user3
export IPBOT_USERS="YOUR_USERNAME"

# Contact @botfather
export IPBOT_TG_TOKEN="SECRET"
export IPBOT_TG_DEBUG="true"
export IPBOT_TG_TIMEOUT="256"

# IFConfig settings
export IPBOT_IFCONFIG_URL="https://ifconfig.me"

if [ -f ./secret.sh ]; then
  source ./secret.sh
fi
if [ -f ./env/secret.sh ]; then
  source ./env/secret.sh
fi
