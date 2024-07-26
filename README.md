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
3. Install pubsub-emulator-tools Tolls
```bash
go install github.com/opannapo/pubsub-emulator-tools@latest
```
___

### Application Usage :
#### Main Menu:
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
--- 

#### Action setup-emulator-compose :
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
```
Docker Output :
```bash
Docker Compose created at  /Users/lion/.tmp-emulator/docker-compose.yml
WARN[0000] /Users/lion/.tmp-emulator/docker-compose.yml: `version` is obsolete 
[+] Running 5/7
 ⠏ pubsub-emulator [⣿⣿⣿⣿⣿⣀] 180.8MB / 463.7MB Pulling                            38.0s 
   ✔ 5de87e84afee Pull complete                                                   10.0s 
   ✔ 34c68e5535f6 Pull complete                                                   10.0s 
   ✔ 819e44463e51 Pull complete                                                   10.0s 
   ✔ 5f0923bf6e65 Pull complete                                                   10.0s 
   ✔ 65b291c23283 Pull complete                                                   10.0s 
   ⠧ 864c6a15293c Downloading [=================>            ]  149.4MB/432.3MB   25.8s
```
```bash
WARN[0000] /Users/lion/.tmp-emulator/docker-compose.yml: `version` is obsolete 
[+] Running 7/7
 ✔ pubsub-emulator Pulled                                                                                                                                                          87.8s 
   ✔ 5de87e84afee Pull complete                                                                                                                                                    10.0s 
   ✔ 34c68e5535f6 Pull complete                                                                                                                                                    10.0s 
   ✔ 819e44463e51 Pull complete                                                                                                                                                    10.0s 
   ✔ 5f0923bf6e65 Pull complete                                                                                                                                                    10.0s 
   ✔ 65b291c23283 Pull complete                                                                                                                                                    10.0s 
   ✔ 864c6a15293c Pull complete                                                                                                                                                    75.1s 
[+] Running 2/2
 ✔ Container pubsub-emulator        Started                                                                                                                                         1.2s 
 ! pubsub-emulator The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested                   0.0s 
Continue [Y/N] ? 

```
Check your docker
```bash
lion@M1 ~ % docker ps
CONTAINER ID   IMAGE                        COMMAND                  CREATED         STATUS         PORTS                    NAMES
3371b9dcc573   google/cloud-sdk:emulators   "/bin/bash -c 'gclou…"   4 minutes ago   Up 4 minutes   0.0.0.0:8085->8085/tcp   pubsub-emulator
lion@M1 ~ % 
```

---

#### Action topic-create :
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
---

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

---
#### Action [topic-delete]
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

Action number : 4
```
Output
```bash
2024/07/26 11:09:18 projects/opannapo-project-id/topics/sub1-topic1
2024/07/26 11:09:18 projects/opannapo-project-id/topics/test1
2024/07/26 11:09:18 projects/opannapo-project-id/topics/test2
[topic-delete] topic name  : test1
Topic test1 deleted.
Continue [Y/N] ? : 
```

---
#### Action [subscription-create]
```bash
  Action number : 5
```
Output
```bash
[subscription-create] subscription name  : sub2
[subscription-create]   topic name  : sub2-topic2
Created subscription: projects/opannapo-project-id/subscriptions/sub2
Continue [Y/N] ? : 
```
---

#### Action [subscription-list]
```bash
  Action number : 6
```
Output:
```bash
2024/07/26 11:11:39 projects/opannapo-project-id/subscriptions/sub1
2024/07/26 11:11:39 projects/opannapo-project-id/subscriptions/sub2
Continue [Y/N] ? : 
```

---
#### Action [subscription-delete]
```bash
  Action number : 7
```
Output:
```bash
2024/07/26 11:12:11 projects/opannapo-project-id/subscriptions/sub1
2024/07/26 11:12:11 projects/opannapo-project-id/subscriptions/sub2
[subscription-delete] subscription name : sub1
Subscription sub1 deleted.
Continue [Y/N] ? : 
```
---

#### Action [start-pub-http]
```bash
  Action number : 8
```
Output:
```bash
[start-pub-http] port  : 8080
[start-pub-http] subscription ID  : sub2
[start-pub-http] topic  : sub2-topic2
2024/07/26 11:16:42 Starting server on port 8080...
```
#### Simulator Testing :
##### 1. Postman
Url : ```POST localhost:8080/publish```
<br>Payload :
```json
{
    "first_name":"John",
    "last_name":"Doe"
}
```

##### 2. CUrl
```bash
curl --location 'localhost:8080/publish' \
--header 'Content-Type: application/json' \
--data '{
    "first_name":"John",
    "last_name":"Doe"
}'
```
---

#### Action [start-sub-cli]
```bash
  Action number : 9
```
Output:
```bash
[start-sub-cli] subscription ID  : sub2
listening for messages ...
```

```bash
[start-sub-cli] subscription ID  : sub2
listening for messages ...
Received message: {
    "first_name":"John1",
    "last_name":"Doe1"
}
Received message: {
    "first_name":"John2",
    "last_name":"Doe2"
}
```
---


### TODO :
| Item                                   |
|----------------------------------------|  
| Graceful Shutdown                      |
| Create Topic with configuration        |
| Create Subscription with configuration |
| Proper error handling                  |

