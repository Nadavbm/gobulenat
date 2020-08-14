rock-my-dudale:
	docker-compose -f docker-compose.yml down && \
	docker-compose -f docker-compose.yml up --build -d && \
	sleep 5 && \
	echo "i think we are ready dude..."