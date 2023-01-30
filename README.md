# Logology

Logology is a log library for cloud driven Go programs which applies best practices which enforce better log details which lead to better metric and log collection. 

## Key features: 

- **Log Types:** compliance, log, metric, transaction rate, debug, trace

- **KPI metrics duration:** specific methods allow to "start" and "end" transactions generating duration metrics for performance or SLO evaluation

- **KPI transaction rate:** ability to define TPS or TPM sampling metrics for counting events

## Data structure features: 

- **Multitenancy:** each log and metric is bound to a tenant identifier related to the application´s customer (tenant). This allows to analyze details related to a specific tenant. 

- **Transaction driven:** a tranaction id is part of most log methods, which allows to co-relate logs related to a specific transaction. This is useful when analyzing and troubleshooting performance, user and sequence flows

- **Service driven:** logs and metrics are bound to a service identifier, enabling to analyze data for specific services (microservices)

- **Scope driven:** a scope is an internal identifier which can be used to identify a specific concept/functionality within a service/code

## What´s next:

- support for OpenTelemetry
- connectors to major log ingestion cloud services

