#!/usr/bin/env bash

# exit on error
set -e

DIR=dst
BUCKET=s3://n.2p5.xyz
DIST_ID=$N_2P5_XYZ_CDN_DIST_ID
CACHE_ID=$(date +%s)
FORMATTED_JSON='{"Paths": {"Quantity": 1,"Items": ["/*"]},"CallerReference":"'${CACHE_ID}'"}'

cd $DIR
aws s3 sync . $BUCKET --acl public-read
aws cloudfront create-invalidation --distribution-id "$DIST_ID" --invalidation-batch "$FORMATTED_JSON"