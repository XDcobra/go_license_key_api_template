# Stage 1: Template processing for env variables
FROM alpine:latest AS configbuilder
RUN apk add --no-cache gettext apache2-utils
WORKDIR /app
COPY prometheus.yml.template .
COPY web.yml.template .

# Pass build args (set via docker-compose)
ARG METRICS_USER
ARG METRICS_PASS
ARG PROM_USER
ARG PROM_PASS
ENV METRICS_USER=$METRICS_USER
ENV METRICS_PASS=$METRICS_PASS
ENV PROM_USER=$PROM_USER
ENV PROM_PASS=$PROM_PASS

# We need to hash the password, as it is saved as bcrypt within the web.yml file
RUN BCRYPT_HASH=$(htpasswd -nbB "$PROM_USER" "$PROM_PASS" | cut -d ':' -f 2 | sed 's/\$2y\$/\$2b\$/') && \
    env PROM_USER=$PROM_USER PROM_PASS_BCRYPT="$BCRYPT_HASH" \
    envsubst < web.yml.template > web.yml && \
    envsubst < prometheus.yml.template > prometheus.yml

# Stage 2: Prometheus runtime
FROM prom/prometheus:latest

# Copy the configuration files
COPY --from=configbuilder /app/prometheus.yml /etc/prometheus/prometheus.yml
COPY --from=configbuilder /app/web.yml        /etc/prometheus/web.yml

