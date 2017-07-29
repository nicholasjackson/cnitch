FROM grafana/grafana

COPY ./grafana.sh ./startup.sh
COPY ./dashboard.json ./dashboard.json

ENTRYPOINT ./startup.sh
