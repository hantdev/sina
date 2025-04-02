# MQTT Benchmarking Tool

A simple MQTT benchmarking tool for Mitras platform.

It connects Mitras clients as subscribers over a number of channels and
uses other Mitras clients to publish messages and create MQTT load.

Mitras clients used must be pre-provisioned first, and Mitras `provision` tool can be used for this purpose.

## Installation

```bash
cd tools/mqtt-bench
make
```
