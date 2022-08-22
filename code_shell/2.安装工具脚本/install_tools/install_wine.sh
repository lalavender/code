sudo dpkg --add-architecture i386
wget -nc https://dl.winehq.org/wine-builds/winehq.key
sudo apt-key add winehq.key
sudo touch /etc/apt/sources.list.d/winehq.list
sudo echo “deb https://dl.winehq.org/wine-builds/debian/ stretch main” >>/etc/apt/sources.list.d/winehq.list
sudo apt install dirmngr
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 76F1A20FF987672F
sudo apt update
sudo apt install --install-recommends winehq-stable
sudo apt install --install-recommends winehq-devel
wine --version


