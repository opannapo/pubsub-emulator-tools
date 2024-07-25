# Experiment
# 
___
### PubSub Emulator Setup : 
(Opt 1) Setup PubSub Emulator (Docker)
```docker
docker pull google/cloud-sdk:emulators
```
```docker
docker run -d -p 8085:8085 google/cloud-sdk:emulators /bin/bash -c "gcloud beta emulators pubsub start --project=test-project-id --host-port='0.0.0.0:8085'"
```

(Opt 2) Setup PubSub Emulator (Docker Compose)
```docker
docker-compose-up
```
#
#
___
### Application Usage: 
___
#### 1. Export env variables
```bash
 source var
```

#### 2. Run tools
```bash
 make run-tools
```
Output :
```bash
Action List
[*] Emulator
        [1] setup-emulator-compose
[*] Application
        [2] topic-create
        [3] topic-list
        [4] topic-delete
        [5] subscription-create
        [6] subscription-list
        [7] subscription-delete
[*] Simulator
        [8] start-pub-http
        [9] start-sub-cli

Action number : 
```

#

#### Action [topic-create]
```bash
Action List
[*] Emulator
        [1] setup-emulator-compose
[*] Application
        [2] topic-create
        [3] topic-list
        [4] topic-delete
        [5] subscription-create
        [6] subscription-list
        [7] subscription-delete
[*] Simulator
        [8] start-pub-http
        [9] start-sub-cli

Action number : 2
```
```bash
[topic-create]  topic name  : test-topic
```

Output :
```bash
[topic-create]  topic name  : test-topic
2024/07/21 00:23:52 Topic Create
create topic test-topic
Topic projects/opannapo-project-id/topics/test-topic created.
Continue [Y/N] ?
```
Back to main menu, type **Y** / **y**

#

#### Action [topic-list]
```bash
Action List
[*] Emulator
        [1] setup-emulator-compose
[*] Application
        [2] topic-create
        [3] topic-list
        [4] topic-delete
        [5] subscription-create
        [6] subscription-list
        [7] subscription-delete
[*] Simulator
        [8] start-pub-http
        [9] start-sub-cli

Action number : 3
```

Output :
```bash
2024/07/21 00:52:04 Topic List
2024/07/21 00:52:04 projects/opannapo-project-id/topics/abc
2024/07/21 00:52:04 projects/opannapo-project-id/topics/abcd
2024/07/21 00:52:04 projects/opannapo-project-id/topics/def
2024/07/21 00:52:04 projects/opannapo-project-id/topics/test-topic
Continue [Y/N] ?
```
Back to main menu, type **Y** / **y**