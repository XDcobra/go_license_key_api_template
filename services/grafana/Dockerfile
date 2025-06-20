# Stage 1: Template processing for env variables
FROM alpine:latest AS configbuilder
RUN apk add --no-cache gettext
WORKDIR /app
COPY datasources.yml.template .

# Pass build args (set via docker-compose)
ARG PROM_USER
ARG PROM_PASS
ENV PROM_USER=$PROM_USER
ENV PROM_PASS=$PROM_PASS

RUN envsubst < datasources.yml.template > datasources.yml


# Stage 2: Prometheus runtime
FROM grafana/grafana:latest
COPY --from=configbuilder /app/datasources.yml /etc/grafana/provisioning/datasources/datasources.yml

