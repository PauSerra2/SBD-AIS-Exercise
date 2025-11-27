# todo note commands

# 1. Firewall Rules

# Allow Docker Swarm Overlay Network to connect all the nodes 
New-NetFirewallRule -DisplayName "Docker Swarm Overlay UDP" -Direction Inbound -LocalPort 4789 -Protocol UDP -Action Allow

# Swarm control traffic (UDP) to find ever node part of the Swarm 
New-NetFirewallRule -DisplayName "Docker Swarm UDP" -Direction Inbound -LocalPort 7946 -Protocol UDP -Action Allow

# Swarm control traffic (TCP)  to coordinate the nodes 
New-NetFirewallRule -DisplayName "Docker Swarm TCP" -Direction Inbound -LocalPort 7946 -Protocol TCP -Action Allow

# 2. Manager node
# To let the workers know the manager IP to connect with him (it changes with every swarm)
docker swarm init --advertise-addr=172.20.10.4 --listen-addr=172.20.10.4:2377

# Retrieve the join token for workers nodes:
docker swarm join-token worker

# Join the worker nodes (it changes with every swarm)
docker swarm join --token SWMTKN-1-3nhvaxjm2tznawawa6ijgzj583dk8rddsat484el7mgvhovsxj-buvaogsbyyo78g5zz790kfcsv 172.17.117.76:2377

# 3. Deploy the Stack
# On the manager:
docker stack deploy -c docker-compose.yml ourstack

# 4. Cleanup (if necessary)
# Remove the deployed stack:
docker stack rm sbd

# To leave the swarm 
	# On the manager
	- docker swarm leave --force

	# On each worker
	- docker swarm leave