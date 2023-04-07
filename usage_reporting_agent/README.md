
# Prerequisites

1. Golang
2. Make

# Steps to implement a traceability agent using this stub

1. Locate the commented tag "CHANGE_HERE" for the package import paths in all files and fix them to reference your code path correctly.
2. Run "make dep" to resolve the dependencies. This should resolve all dependency packages and vendor them under ./vendor directory
3. Update Makefile to change the name of generated binary image from *apic_traceability_agent* to the desired name. Locate the comment "CHANGE_BINARY_NAME" and follow the instructions in the comment.
4. Update pkg/cmd/root.go to change the agent yaml file name and description of the agent. Locate *apic_traceability_agent* and *Sample Traceability Agent* and replace with the desired values.
5. Update your specific gateway configuration in pkg/config/config.go
    - Locate *gateway-section* in the sample YAML config file and replace it with the name of your gateway (e.g. MyGateway). Locate this same tag in pkg/cmd/root.go and sample YAML config file and replace it with the same name.
    - Define gateway specific config properties in *GatewayConfig* struct. Locate the struct variables *ConfigKey1* & struct *config_key_1* and replace them with your desired config properties.
    - Optionally add config validation for your config values. Locate the *ValidateCfg()* method in config.go and update the implementation to add validation specific to gateway specific config.
    - Optionally add ApplyResources configuration. Locate the *ApplyResources()* method in and update the implementation to add logic to copy ResourceConfiguration values to your gateway specific config.
    - Update the config binding with command line flags in init(). Locate *gateway-section.config_key_1* (by now 'gateway-section' should have been changed to the name of your gateway in a step above) and add replace all config property bindings with the correct values
    - Update the initialization of gateway specific config by parsing the binded properties. Locate *ConfigKey1* & *gateway-section.config_key_1* (again, 'gateway-section' should have been changed to the name of your gateway in a step above) and add/replace all config properties
6. Locate pkg/gateway/definition.go to define the structure of the log entry the traceability agent will receive. See pkg/gateway/definition.go for sample definition.
7. Implement the mechanism to read the log entry. Optionally you can wrap the existing beat(for e.g. filebeat) in beater.New() to read events and they setup output event processor to process the events
8. Locate pkg/gateway/eventmapper.go to map the log entry received by beat to event structure expected for Amplify Central Observer.
9. Locate pkg/gateway/eventprocessor.go to perform processing on event to be published. The processing can be performed either on the received event by beat input or before the event is published by transport. See pkg/gateway/eventprocessor.go for example of both type of processing.
10. Run "make build" to build the agent
11. Rename *apic_traceability_agent.yml* file to the agent name you previously specified in pkg/cmd/root.go and set up the agent config in the file.
12. Copy the YAML config file to the *bin* directory.
12. Execute the agent by running the binary file generated under *bin* directory.
13. To produce traffic update the ./logs/traffic.log file with a new entry. See ./logs/traffic.log for sample entries

Reference: [SDK Documentation - Building Traceability Agent](https://github.com/Axway/agent-sdk/blob/main/docs/traceability/index.md)
