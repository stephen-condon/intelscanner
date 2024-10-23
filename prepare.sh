# update version file
printf 'Updated version file (y/N)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ] ;then 
    # run macos build
    go build
    # run windows build
    env GOOS=windows GOARCH=amd64 go build
    # move executables to dist
    cp ./intel-scanner ./dist/intel-scanner
    cp ./intel-scanner.exe ./dist/intel-scanner.exe

    rm ./intel-scanner
    rm ./intel-scanner.exe

    echo "Successfully built executables, moved to dist"
    exit 0
else
    echo "please update the version file following semver"
    exit 0
fi
