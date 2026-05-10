dev:
	@trap 'kill 0' SIGINT; \
	(cd backend && go run main.go) & \
	(cd frontend && npm run dev -- --host) & \
	wait
