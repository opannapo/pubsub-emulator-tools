.PHONY: run-pub run-sub run-tools emulator-compose-up emulator-compose-down

#Application
run-tools:
	go run ./main.go
#Application


#Pubsub Emulator Setup
emulator-compose-up:
	cd emulator && docker compose up -d

emulator-compose-down:
	cd emulator && docker compose down
#End Pubsub Emulator Setup

