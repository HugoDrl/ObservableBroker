# Observable MQTT broker

## Goal

This is a MQTT broker that allows a user to retrieve data and metrics using an HTTP interface.  
This is most useful for supervising broker and clients.

## Usage

This tool is mainly meant to allow bridging between things and listeners using both protocols, as MQTT is used in Internet Of Things (IoT) and HTTP is used by cloud servers, mainly.  
This tool is mainly made because I like the idea of connecting and giving access to connectivity without having to do file configs or code, and try to make it attractive.

## Development

Here is an exhaustive list of what it does, and what it will do :  
This is mainly my ideas for future and are for me, these are not promises.

- [x] Allows MQTT clients to publish and subscribe
- [x] Allows HTTP clients to request for holded messages
- [ ] Allows HTTP clients to elaborate request using filters
- [x] Allows HTTP clients to request for currently connected clients
- [ ] Allows HTTP clients to request current clients connection time
- [ ] Allows HTTP clients to request for history connections
- [ ] Alerts for massive disconnection or reconnection
- [ ] Allows HTTP clients to request for unsubscribed holded messages
- [ ] Allows both clients to see data transmission and storage capacity
- [ ] Allows HTTP clients to "publish" to MQTT broker
- [ ] Keeps metrics of "health" and quality of topics and messages (format)
- [ ] Logs everything
- [ ] Allows authentication both on HTTP and MQTT side
- [ ] Encrypts both protocols
- [ ] Uses a real data storing feature
- [ ] Display a GUI for non-CLI people
- [ ] Display a TUI for others
- [ ] Allows HTTP clients to request for topic granularity (paquets size, pub/sub numbers, last message sent)
- [ ] Retransmission rates for QoS 1 and 2, percent of messages per QoS
- [ ] Topics with retained messages
- [ ] Computer usage metrics
- [ ] LWT list (maybe ?)
- [ ] Subscribed topics (including wildcards)
