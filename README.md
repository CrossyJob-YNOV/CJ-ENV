# CJENV

### CrossyJob Environment

## Setup

1. Clone this repo in `/data`

2. Execute the `install-packages.sh` script

3. Execute the `create-users.sh` script

4. Make a symbolic link from `/data/CJ-ENV/nginx.conf` to `/etc/nginx/sites-enabled/`

5. Apply nginx config
    ```shell
    # Test the config
    nginx -t
    # If it's ok, restart it
    systemctl restart nginx
    ```

6. Then, follow the installation steps of each service (harbor, jenkins, k8s)
