#!/bin/bash
request_body=$(cat <<EOF
{
    "resource":"money"
}
EOF
)

curl -v -X POST \
     -d "$request_body" \
    'http://127.0.0.1:7878/kugo/api/v2/namespaces/chao/tasks' | python -m json.tool