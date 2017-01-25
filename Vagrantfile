# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.provider "virtualbox" do |vb|
     vb.gui = false
     vb.cpus = 1
     vb.memory = 1024
  end

  config.vm.synced_folder '.', '/opt/goat'
  # config.vm.network 'public_network', bridge: 'enp0s31f6'
  # config.vm.network 'public_network', bridge: 'p1p1'
  # config.vm.network 'public_network', bridge: 'eth0'
  
  config.vm.define 'node01', autostart: true do |host|
    host.vm.hostname = 'node01'
    host.vm.network 'private_network', ip: '10.20.30.41'
    # host.vm.network "forwarded_port", guest: 80, host: 8080
    # host.vm.provision 'file', source: 'someconfig.cfg', destination: '/tmp/someconfig.cfg'
  end

  config.vm.define 'node02', autostart: true do |host|
    host.vm.hostname = 'node02'
    host.vm.network 'private_network', ip: '10.20.30.42'
    # host.vm.network "forwarded_port", guest: 80, host: 8080
    # host.vm.provision 'file', source: 'someconfig.cfg', destination: '/tmp/someconfig.cfg'
  end

  config.vm.define 'node03', autostart: true do |host|
    host.vm.hostname = 'node02'
    host.vm.network 'private_network', ip: '10.20.30.43'
    # host.vm.network "forwarded_port", guest: 80, host: 8080
    # host.vm.provision 'file', source: 'someconfig.cfg', destination: '/tmp/someconfig.cfg'
  end

  config.vm.provision "shell", inline: <<-SHELL
    BOOTSTRAPPED=/var/run/bootstrapped
    if [ ! -e $BOOTSTRAPPED ]; then
      sudo mv -v /etc/apt/apt.conf.d/70debconf /root/etc-apt-apt.conf.d-70debconf.bak
      sudo dpkg-reconfigure debconf -f noninteractive -p critical
      sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get -y dist-upgrade && sudo apt-get -y autoremove
      sudo apt-get install -y language-pack-en
      echo TODO
      touch $BOOTSTRAPPED
    else
      echo TODO
    fi
  SHELL
end
