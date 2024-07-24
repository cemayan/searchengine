#! /bin/bash

helm upgrade -i \
  redis-stack-server redis \
  --atomic \
  --repo https://charts.bitnami.com/bitnami \
  --version 17.1.4 \
  --values - <<EOF
global:
  redis:
    password: "weak"
image:
  repository: "redis/redis-stack-server"
  tag: "6.2.4-v2"
master:
  args:
    - -c
    - /opt/bitnami/scripts/merged-start-scripts/start-master.sh
  extraVolumes:
    - name: merged-start-scripts
      configMap:
        name: bitnami-redis-stack-server-merged
        defaultMode: 0755
  extraVolumeMounts:
    - name: merged-start-scripts
      mountPath: /opt/bitnami/scripts/merged-start-scripts
replica:
  args:
    - -c
    - /opt/bitnami/scripts/merged-start-scripts/start-replica.sh
  extraVolumes:
    - name: merged-start-scripts
      configMap:
        name: bitnami-redis-stack-server-merged
        defaultMode: 0755
  extraVolumeMounts:
    - name: merged-start-scripts
      mountPath: /opt/bitnami/scripts/merged-start-scripts
extraDeploy:
  - apiVersion: v1
    kind: ConfigMap
    metadata:
      name: bitnami-redis-stack-server-merged
    data:
      start-master.sh: |
        #!/usr/bin/dumb-init /bin/bash

        ### docker entrypoint script, for starting redis stack
        BASEDIR=/opt/redis-stack
        cd \${BASEDIR}

        CMD=\${BASEDIR}/bin/redis-server

        if [ -z "\${REDISEARCH_ARGS}" ]; then
        REDISEARCH_ARGS="MAXSEARCHRESULTS 10000 MAXAGGREGATERESULTS 10000"
        fi

        if [ -z "\${REDISGRAPH_ARGS}" ]; then
        REDISGRAPH_ARGS="MAX_QUEUED_QUERIES 25 TIMEOUT 1000 RESULTSET_SIZE 10000"
        fi

        [[ -f \$REDIS_PASSWORD_FILE ]] && export REDIS_PASSWORD="\$(< "\${REDIS_PASSWORD_FILE}")"
        if [[ -f /opt/bitnami/redis/mounted-etc/master.conf ]];then
            cp /opt/bitnami/redis/mounted-etc/master.conf /opt/bitnami/redis/etc/master.conf
        fi
        if [[ -f /opt/bitnami/redis/mounted-etc/redis.conf ]];then
            cp /opt/bitnami/redis/mounted-etc/redis.conf /opt/bitnami/redis/etc/redis.conf
        fi

        \${CMD} \
        --port "\${REDIS_PORT}" \
        --requirepass "\${REDIS_PASSWORD}" \
        --masterauth "\${REDIS_PASSWORD}" \
        --include "/opt/bitnami/redis/etc/redis.conf" \
        --include "/opt/bitnami/redis/etc/master.conf" \
        --loadmodule /opt/redis-stack/lib/redisearch.so \${REDISEARCH_ARGS} \
        --loadmodule /opt/redis-stack/lib/redisgraph.so \${REDISGRAPH_ARGS} \
        --loadmodule /opt/redis-stack/lib/redistimeseries.so \${REDISTIMESERIES_ARGS} \
        --loadmodule /opt/redis-stack/lib/rejson.so \${REDISJSON_ARGS} \
        --loadmodule /opt/redis-stack/lib/redisbloom.so \${REDISBLOOM_ARGS}
      start-replica.sh: |
        #!/usr/bin/dumb-init /bin/bash

        BASEDIR=/opt/redis-stack
        cd \${BASEDIR}
        CMD=\${BASEDIR}/bin/redis-server

        get_port() {
            hostname="\$1"
            type="\$2"

            port_var=\$(echo "\${hostname^^}_SERVICE_PORT_\$type" | sed "s/-/_/g")
            port=\${!port_var}

            if [ -z "\$port" ]; then
                case \$type in
                    "SENTINEL")
                        echo 26379
                        ;;
                    "REDIS")
                        echo 6379
                        ;;
                esac
            else
                echo \$port
            fi
        }

        get_full_hostname() {
            hostname="\$1"
            echo "\${hostname}.\${HEADLESS_SERVICE}"
        }

        REDISPORT=\$(get_port "\$HOSTNAME" "REDIS")

        [[ -f \$REDIS_PASSWORD_FILE ]] && export REDIS_PASSWORD="\$(< "\${REDIS_PASSWORD_FILE}")"
        [[ -f \$REDIS_MASTER_PASSWORD_FILE ]] && export REDIS_MASTER_PASSWORD="\$(< "\${REDIS_MASTER_PASSWORD_FILE}")"
        if [[ -f /opt/bitnami/redis/mounted-etc/replica.conf ]];then
            cp /opt/bitnami/redis/mounted-etc/replica.conf /opt/bitnami/redis/etc/replica.conf
        fi
        if [[ -f /opt/bitnami/redis/mounted-etc/redis.conf ]];then
            cp /opt/bitnami/redis/mounted-etc/redis.conf /opt/bitnami/redis/etc/redis.conf
        fi

        echo "" >> /opt/bitnami/redis/etc/replica.conf
        echo "replica-announce-port \$REDISPORT" >> /opt/bitnami/redis/etc/replica.conf
        echo "replica-announce-ip \$(get_full_hostname "\$HOSTNAME")" >> /opt/bitnami/redis/etc/replica.conf
        \${CMD} \
        --port "\${REDIS_PORT}" \
        --requirepass "\${REDIS_PASSWORD}" \
        --masterauth "\${REDIS_PASSWORD}" \
        --include "/opt/bitnami/redis/etc/redis.conf" \
        --include "/opt/bitnami/redis/etc/replica.conf" \
        --loadmodule /opt/redis-stack/lib/redisearch.so \${REDISEARCH_ARGS} \
        --loadmodule /opt/redis-stack/lib/redisgraph.so \${REDISGRAPH_ARGS} \
        --loadmodule /opt/redis-stack/lib/redistimeseries.so \${REDISTIMESERIES_ARGS} \
        --loadmodule /opt/redis-stack/lib/rejson.so \${REDISJSON_ARGS} \
        --loadmodule /opt/redis-stack/lib/redisbloom.so \${REDISBLOOM_ARGS}
EOF
