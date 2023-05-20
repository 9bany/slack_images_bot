#!/usr/bin/env bash
set -e
source /app/app.env
echo "start server"
exec "$@"