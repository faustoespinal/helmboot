# Install script based on example by Chris Every (https://tech.ovoenergy.com/writing-a-helm-plugin/)
# cd to the plugin dir
cd $HELM_PLUGIN_DIR

# get the version
version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"

# find the OS and ARCH
unameOut="$(uname -s)"

case "${unameOut}" in
    Linux*)     os=linux;;
    Darwin*)    os=osx;;
    CYGWIN*)    os=windows;;
    MINGW*)     os=windows;;
    *)          os="UNKNOWN:${unameOut}"
esac

arch=`uname -m`

# set the url of the tar.gz
url="https://github.com/faustoespinal/helmboot/releases/download/v${version}/helmboot_${os}.tar.gz"

# set the filename
filename=`echo ${url} | sed -e "s/^.*\///g"`

# download the archive using curl or wget
if [ -n $(command -v curl) ]
then
    curl -sSL -O $url
elif [ -n $(command -v wget) ]
then
    wget -q $url
else
    echo "Need curl or wget"
    exit -1
fi

# extract the plugin binary into the bin dir
rm -rf bin && mkdir bin && tar xzvf $filename -C bin > /dev/null && rm -f $filename
