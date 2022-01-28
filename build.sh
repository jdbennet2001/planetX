echo Building web server
go build

echo Build command line app
cd scan
go build
mv scan ../catalog-scan
