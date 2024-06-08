# syntax=docker/dockerfile:1
FROM apache/airflow:2.2.4

WORKDIR /app

# Optionally copy your Airflow DAGs and configuration here
# COPY dags /opt/airflow/dags

EXPOSE 8080

ENTRYPOINT ["airflow", "standalone"]
