# go-virtual

## Introduction
A CAN ECU virtualization library written in Golang.

## Purpose
The aim of this project is to have an application that can virtualize CAN bus ECUs, including support for multiple ECUS, distributed across multiple networks.

This project aims to make setting up these sort of test environments simple and easy, with miminal overhead.
A developer should be able to virtualize a series of networks by setting up ECUs multiple ways:

1. Setup periodic transmission of CAN messages (typically defined through .dbc files) across multiple CAN buses

  1. Be able to easily specify default signal values or distribute a default value across the entire signal structure of the CAN message
  
  2. Allow for build up of signal queues, that will be depopulated at the transmit rate
  
2. Setup trigger based transmission of CAN messages or signals

  1. Trigger a CAN message or signal(s) to be transmitted if a selected signal or message is transmitted on the virtual bus
  
3. Allow for staggering of messages, while maintaining periodic transmission timing.

4. Pub/Sub architecture for allowing external applications to subscribe to certain messages/signals so the user can:

  1. Sub
  
    1. Develop their own external dashboard to monitor bus signals
    
    2. Feed signals to another application
    
  2. Pub
  
    1. Allow users to publish message/signals to the ECU virtualization tool. This will allow for users to send signals in near real time, so that they can develop simulations by updating values that should be sent by the simulated ECU(s).
    
    2. Dashboards can be developed to simulate things like joysticks, switches, etc., rather than relying on physical hardware
    
  3. Testing
  
    1. Utilize both pub/sub at the same time to develop test suites/unit tests to fulling contain their testing environment without relying on having a full test bench to do (manual) testing.
