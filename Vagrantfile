# -*- mode: ruby -*-
# vi: set ft=ruby :
#
# sudo docker -d -H tcp://0.0.0.0:4243 &
# or
# sudo docker -d & (unix socket)
# sudo docker images (example)
#
# after create a container
# sudo docker inspect 4c2d8aaf3158 | grep IPAddress | cut -d '"' -f 4
#
# mysql -u root -D mysql -e "GRANT ALL ON *.* to root@'%' IDENTIFIED BY ''; FLUSH PRIVILEGES;"
#
# scp -P 2222 ~/.ssh/id_rsa.pub vagrant@localhost:/home/vagrant/.ssh/
# vagrant ssh
# cd .ssh/
# cat id_rsa.pub >> authorized_keys
Vagrant.configure("2") do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"
  config.vm.network :forwarded_port, guest: 4243, host: 4243
  config.vm.synced_folder './', '/home/vagrant/go/src'
  config.vm.provision :shell, inline: <<-SCRIPT
    # install Go
    wget -O /tmp/golang.deb https://cupcake-ops.s3.amazonaws.com/go_1.1.2-godeb1_amd64.deb
    dpkg -i /tmp/golang.deb
    sudo -u vagrant mkdir /home/vagrant/go
    sudo -u vagrant mkdir -p /home/vagrant/go/{src,bin,pkg}
    sudo -u vagrant mkdir -p /home/vagrant/repos
    sudo -u vagrant sh -c "echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.bashrc"
    sudo -u vagrant sh -c "echo 'export PATH=$PATH:home/vagrant/go/bin' >> /home/vagrant/.bashrc"

    # install Docker
    wget -qO - https://get.docker.io/gpg | apt-key add -
    echo "deb http://get.docker.io/ubuntu docker main" > /etc/apt/sources.list.d/docker.list
    apt-get update
    apt-get install -y linux-image-generic-lts-raring linux-headers-generic-lts-raring lxc-docker git mercurial curl mysql-server mysql-client
    sed -i -E 's|	/usr/bin/docker -d|	/usr/bin/docker -d -H tcp://0.0.0.0:4243|' /etc/init/docker.conf
    sed -i -E 's|	bind-address|	#bind-address|' /etc/mysql/my.cnf
    reboot
  SCRIPT
end
