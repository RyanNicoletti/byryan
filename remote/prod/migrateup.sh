#!/bin/bash
source /etc/byryanweb/production.env
migrate -path /opt/byryan/migrations -database "$PROD_DB_DSN" up