#!/usr/bin/bash

package=$1

if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}
platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64")
print() {
     echo "$@" | sed \
             -e "s/\(\(@\(red\|green\|yellow\|blue\|magenta\|cyan\|white\|reset\|b\|u\)\)\+\)[[]\{2\}\(.*\)[]]\{2\}/\1\4@reset/g" \
             -e "s/@red/$(tput setaf 1)/g" \
             -e "s/@green/$(tput setaf 2)/g" \
             -e "s/@yellow/$(tput setaf 3)/g" \
             -e "s/@blue/$(tput setaf 4)/g" \
             -e "s/@magenta/$(tput setaf 5)/g" \
             -e "s/@cyan/$(tput setaf 6)/g" \
             -e "s/@white/$(tput setaf 7)/g" \
             -e "s/@reset/$(tput sgr0)/g" \
             -e "s/@b/$(tput bold)/g" \
             -e "s/@u/$(tput sgr 0 1)/g"
}
for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name="Binary/$GOOS-$GOARCH"
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	
	print @b@yellow[["[+] build : $platform"]] 
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
	if [ $? -ne 0 ]; then
		print @b@red[["An error has occurred! Aborting the script execution..."]]
		exit 1
	fi
done
