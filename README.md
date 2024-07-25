# Experiment
Internal tools to run & simulation Pub/Sub on localhost
___

### Preparation

1. Install Docker

```
https://docs.docker.com/desktop/install/linux-install/
https://docs.docker.com/desktop/install/mac-install/
```

2. Setup / Export environment variables (Linux/Mac)

```
export PUBSUB_EMULATOR_HOST=localhost:8085
export PUBSUB_PROJECT_ID=opannapo-project-id
```

___

### Application Usage :
Main Menu:
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

### Action setup-emulator-compose :
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

Action number : 1
```
```bash
[setup-emulator-compose] project ID  : project1
[setup-emulator-compose] port  : 8085 
Docker Compose created at  /Users/lion/.tmp-emulator/docker-compose.yml
WARN[0000] /Users/lion/.tmp-emulator/docker-compose.yml: `version` is obsolete 
[+] Running 5/7
 ⠏ pubsub-emulator [⣿⣿⣿⣿⣿⣀] 180.8MB / 463.7MB Pulling                                                                                                                                                                      38.0s 
   ✔ 5de87e84afee Pull complete                                                                                                                                                                                            10.0s 
   ✔ 34c68e5535f6 Pull complete                                                                                                                                                                                            10.0s 
   ✔ 819e44463e51 Pull complete                                                                                                                                                                                            10.0s 
   ✔ 5f0923bf6e65 Pull complete                                                                                                                                                                                            10.0s 
   ✔ 65b291c23283 Pull complete                                                                                                                                                                                            10.0s 
   ⠧ 864c6a15293c Downloading [=================>                                 ]  149.4MB/432.3MB                                                                                                                       25.8s 

```

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