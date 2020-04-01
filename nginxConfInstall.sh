#! /bin/bash

pre_config_file=nginx.conf
tmp=/tmp/go-talk
tmp_conf=$tmp/$pre_config_file
nginx_conf=/etx/nginx/conf.d/go-talk.conf

if [[ -z $1 ]]; then
  echo "go-talk: [error] hosting domiain not specified";
  exit 1;
elif [[ -z $pre_config_file ]]; then
  echo "go-talk: [error] nginx pre-config file not set";
  exit 1
elif [[ ! -f $pre_config_file ]]; then
  echo "go-talk: [error] nginx pre-config file not found";
  exit 1;
fi

# Set hosting domain, create nginx configuration
sed "s/GO_TALK_DOMAIN/$1" $pre_config_file > $tmp_conf;
# Copy nginx configuration over
cp "$tmp_conf"  $nginx_conf;
# Clean up tmp
rm -rf $tmp;
