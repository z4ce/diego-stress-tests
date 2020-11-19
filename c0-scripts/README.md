# Install tools
1. Setup environment to be tested
1. Login to opsmgr vm
2. Install tools
```
wget --content-disposition 'https://github.com/pivotal-cf/om/releases/download/6.4.1/om-linux-6.4.1'
wget --content-disposition 'https://packages.cloudfoundry.org/stable?release=linux64-binary&version=6.53.0&source=github-rel'
tar -xvf cf-cli_6.53.0_linux_x86-64.tgz
chmod a+x cf-*
mv cf-* /usr/local/bin
chmod a+x om-*
mv om-* /usr/local/bin

export GOPATH=/home/ubuntu/go
mkdir $GOPATH

export PATH="$PATH:/usr/lib/go-1.10/bin"
go get github.com/cloudfoundry/uptimer
export PATH="$PATH:/home/ubuntu/go/bin"
```
3. Copy the scripts in this folder to `scp -i key ./ ubuntu@opsmgr:/tmp`
4. Modify the scripts ending in .sh with your opsmgr password
5. Modify the uptimer definitions ending in .json with your cloudfoundry details (api/username/password)

# Create test load 
With a computer with `calc` installed (brew install calc)
1. Run `./start-apps <mem of your platform>`
for example
```
╰─ ./start-apps 3000
cf create-org stress-test
cf create-space -o stress-test stress-space
cf create-quota idowhatiwant -a -1 -i -1 -m 4000g -s 1000000 -r 90000000
cf set-quota stress-test idowhatiwant
cf push extra-heavy-1 -f ../manifests/manifest-heavy-extra.yml -i 53
cf push extra-heavy-2 -f ../manifests/manifest-heavy-extra.yml -i 53
cf push heavy-1 -f ../manifests/manifest-heavy.yml -i 53
cf push heavy-2 -f ../manifests/manifest-heavy.yml -i 53
cf push medium-1 -f ../manifests/manifest-medium.yml -i 105
cf push medium-2 -f ../manifests/manifest-medium.yml -i 105
cf push medium-extra-1 -f ../manifests/manifest-medium-extra.yml -i 105
cf push medium-extra-2 -f ../manifests/manifest-medium-extra.yml -i 105
cf push light-1 -f ../manifests/manifest-light.yml -i 105
cf push light-2 -f ../manifests/manifest-light.yml -i 105
cf push light-extra-1 -f ../manifests/manifest-light-extra.yml -i 105
cf push light-extra-2 -f ../manifests/manifest-light-extra.yml -i 105
```
2. Enter a machine with this repo and cf cli with access to your test environment (you can use opsmgr)
3. Enter the /cedar/assets/stress-app in this repository 
4. Run `go build`
5. Run the cf cli commands given by the start-apps command.

# CF Push time measurement
```
cd cedar/assets/stress-app
time cf push test-time -f ../manifests/manifest-medium.yml
```
# Run the tests

1. Execute the uptimer test, for example for a cert rotation

```
uptimer -configFile ./uptimer-rotate.json -resultFile ./uptimer-200-cert-rotation-results.json | tee uptimer-200-cert-rotation.log
```
Scripts exist for cert rotate, bbr, and apply-changes (upgrade). For apply-changes the upgrade procedure should be followed 
7. Observe availability statistics
8. Calculate total duration using the bosh begin and end times for task you run

# Setting up for Upgrade
Here some example commands you can use to accelerate the upgrade staging
```
OM_PASS='xx'
o() {
  om -k -t https://127.0.0.1 -u admin -p "${OM_PASS}" "$@"
}

pivnet login
pivnet download-product-files -p elastic-runtime -r 2.9.13 -g 'cf*.pivotal'
o upload-product --product cf-2.9.13-build.12.pivotal
pivnet download-product-files -p stemcells-ubuntu-xenial -r 621.87 -g '*vsphere*.tgz'
o upload-stemcell --stemcell bosh-stemcell-621.87-vsphere-esxi-ubuntu-xenial-go_agent.tgz
o stage-product --product-name cf --product-version 2.9.13
uptimer -configFile ./uptimer-apply-changes.json -resultFile ./uptimer-200-apply-changes-results.json | tee uptimer-200-apply-changes.log
```
