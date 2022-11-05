install:
	# install
	sudo apt-get update
	sudo apt-get install golang apache2 -y 
	sudo apt install mariadb-server mariadb-client -y
	# setup
	sudo ln -s ./Binary/ ./Server/
	sudo cp ./Server/ /var/www/html/Server