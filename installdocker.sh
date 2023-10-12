#!/bin/sh

sudo sh -c 'apt-get update
apt-get upgrade -y
apt-get dist-upgrade -y'

echo =========================================================================================================================================================================================
echo ==========START INSTALLING DOCKER====================START INSTALLING DOCKER====================START INSTALLING DOCKER==================================================================
echo =========================================================================================================================================================================================
sudo sh -c 'apt-get update
apt-get -y install \
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg2 \
	software-properties-common
curl -fsSL https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg |  apt-key add -
add-apt-repository \
	"deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
	$(lsb_release -cs) \
	edge"
apt-get update
apt-get install -y docker-ce'
echo =========================================================================================================================================================================================
echo ==========THE DOCKER ARE INSTALLED==============================THE DOCKER ARE INSTALLED==============================THE DOCKER ARE INSTALLED===========================================
echo =========================================================================================================================================================================================